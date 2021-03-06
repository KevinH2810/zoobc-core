// ZooBC Copyright (C) 2020 Quasisoft Limited - Hong Kong
// This file is part of ZooBC <https://github.com/zoobc/zoobc-core>
//
// ZooBC is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// ZooBC is distributed in the hope that it will be useful, but
// WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.
// See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with ZooBC.  If not, see <http://www.gnu.org/licenses/>.
//
// Additional Permission Under GNU GPL Version 3 section 7.
// As the special exception permitted under Section 7b, c and e,
// in respect with the Author’s copyright, please refer to this section:
//
// 1. You are free to convey this Program according to GNU GPL Version 3,
//     as long as you respect and comply with the Author’s copyright by
//     showing in its user interface an Appropriate Notice that the derivate
//     program and its source code are “powered by ZooBC”.
//     This is an acknowledgement for the copyright holder, ZooBC,
//     as the implementation of appreciation of the exclusive right of the
//     creator and to avoid any circumvention on the rights under trademark
//     law for use of some trade names, trademarks, or service marks.
//
// 2. Complying to the GNU GPL Version 3, you may distribute
//     the program without any permission from the Author.
//     However a prior notification to the authors will be appreciated.
//
// ZooBC is architected by Roberto Capodieci & Barton Johnston
//             contact us at roberto.capodieci[at]blockchainzoo.com
//             and barton.johnston[at]blockchainzoo.com
//
// Core developers that contributed to the current implementation of the
// software are:
//             Ahmad Ali Abdilah ahmad.abdilah[at]blockchainzoo.com
//             Allan Bintoro allan.bintoro[at]blockchainzoo.com
//             Andy Herman
//             Gede Sukra
//             Ketut Ariasa
//             Nawi Kartini nawi.kartini[at]blockchainzoo.com
//             Stefano Galassi stefano.galassi[at]blockchainzoo.com
//
// IMPORTANT: The above copyright notice and this permission notice
// shall be included in all copies or substantial portions of the Software.
package query

import (
	"database/sql"
	"fmt"
	"github.com/zoobc/zoobc-core/common/blocker"
	"strings"

	"github.com/zoobc/zoobc-core/common/model"
)

type (
	FeeVoteRevealVoteQueryInterface interface {
		GetFeeVoteRevealByAccountAddressAndRecentBlockHeight(accountAddress []byte, blockHeight uint32) (string, []interface{})
		GetFeeVoteRevealsInPeriod(
			lowerBlockHeight, upperBlockHeight uint32,
		) (string, []interface{})
		InsertRevealVote(revealVote *model.FeeVoteRevealVote) (string, []interface{})
		InsertRevealVotes(revealVotes []*model.FeeVoteRevealVote) (str string, args []interface{})
		Scan(vote *model.FeeVoteRevealVote, row *sql.Row) error
		BuildModel(votes []*model.FeeVoteRevealVote, rows *sql.Rows) ([]*model.FeeVoteRevealVote, error)
	}
	FeeVoteRevealVoteQuery struct {
		Fields    []string
		TableName string
	}
)

func NewFeeVoteRevealVoteQuery() *FeeVoteRevealVoteQuery {
	return &FeeVoteRevealVoteQuery{
		Fields: []string{
			"recent_block_hash",
			"recent_block_height",
			"fee_vote",
			"voter_address",
			"voter_signature",
			"block_height",
		},
		TableName: "fee_vote_reveal_vote",
	}
}

func (fvr *FeeVoteRevealVoteQuery) getTableName() string {
	return fvr.TableName
}

// GetFeeVoteRevealByAccountAddressAndRecentBlockHeight represents getting fee_vote_reveal by account address
func (fvr *FeeVoteRevealVoteQuery) GetFeeVoteRevealByAccountAddressAndRecentBlockHeight(
	accountAddress []byte,
	blockHeight uint32,
) (qry string, args []interface{}) {
	return fmt.Sprintf(
		"SELECT %s FROM %s WHERE voter_address = ? AND recent_block_height = ? ORDER BY block_height DESC LIMIT 1",
		strings.Join(fvr.Fields, ", "),
		fvr.getTableName(),
	), []interface{}{accountAddress, blockHeight}
}

// GetFeeVoteRevealsInPeriod select reveals within block-height range
// blockHeight limit are inclusive
func (fvr *FeeVoteRevealVoteQuery) GetFeeVoteRevealsInPeriod(
	lowerBlockHeight, upperBlockHeight uint32,
) (qry string, args []interface{}) {
	return fmt.Sprintf(
		"SELECT %s FROM %s WHERE block_height between ? AND ? ORDER BY fee_vote ASC",
		strings.Join(fvr.Fields, ", "),
		fvr.getTableName(),
	), []interface{}{lowerBlockHeight, upperBlockHeight}
}

// InsertRevealVote represents insert new record to fee_vote_reveal
func (fvr *FeeVoteRevealVoteQuery) InsertRevealVote(revealVote *model.FeeVoteRevealVote) (qry string, args []interface{}) {
	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES(%s)",
		fvr.getTableName(),
		strings.Join(fvr.Fields, ", "),
		fmt.Sprintf("?%s", strings.Repeat(", ?", len(fvr.Fields)-1)),
	), fvr.ExtractModel(revealVote)
}
func (fvr *FeeVoteRevealVoteQuery) InsertRevealVotes(revealVotes []*model.FeeVoteRevealVote) (str string, args []interface{}) {
	if len(revealVotes) > 0 {
		str = fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES ",
			fvr.getTableName(),
			strings.Join(fvr.Fields, ", "),
		)
		for k, revealVote := range revealVotes {
			str += fmt.Sprintf(
				"(?%s)",
				strings.Repeat(", ?", len(fvr.Fields)-1),
			)
			if k < len(revealVotes)-1 {
				str += ","
			}
			args = append(args, fvr.ExtractModel(revealVote)...)
		}
	}
	return str, args

}

// ImportSnapshot takes payload from downloaded snapshot and insert them into database
func (fvr *FeeVoteRevealVoteQuery) ImportSnapshot(payload interface{}) ([][]interface{}, error) {
	var (
		queries [][]interface{}
	)
	reveals, ok := payload.([]*model.FeeVoteRevealVote)
	if !ok {
		return nil, blocker.NewBlocker(blocker.DBErr, "ImportSnapshotCannotCastTo"+fvr.TableName)
	}
	if len(reveals) > 0 {
		recordsPerPeriod, rounds, remaining := CalculateBulkSize(len(fvr.Fields), len(reveals))
		for i := 0; i < rounds; i++ {
			qry, args := fvr.InsertRevealVotes(reveals[i*recordsPerPeriod : (i*recordsPerPeriod)+recordsPerPeriod])
			queries = append(queries, append([]interface{}{qry}, args...))
		}
		if remaining > 0 {
			qry, args := fvr.InsertRevealVotes(reveals[len(reveals)-remaining:])
			queries = append(queries, append([]interface{}{qry}, args...))
		}
	}
	return queries, nil
}

// RecalibrateVersionedTable recalibrate table to clean up multiple latest rows due to import function
func (fvr *FeeVoteRevealVoteQuery) RecalibrateVersionedTable() []string {
	return []string{}
}

// ExtractModel extracting model.FeeVoteRevealVote values
func (*FeeVoteRevealVoteQuery) ExtractModel(revealVote *model.FeeVoteRevealVote) []interface{} {
	return []interface{}{
		revealVote.VoteInfo.GetRecentBlockHash(),
		revealVote.VoteInfo.GetRecentBlockHeight(),
		revealVote.VoteInfo.GetFeeVote(),
		revealVote.GetVoterAddress(),
		revealVote.GetVoterSignature(),
		revealVote.GetBlockHeight(),
	}
}

// Scan build model.FeeVoteRevealVote from sql.Row
func (fvr *FeeVoteRevealVoteQuery) Scan(vote *model.FeeVoteRevealVote, row *sql.Row) error {
	var (
		voterAddress   []byte
		blockHeight    uint32
		voterSignature []byte
		feeVoteInfo    model.FeeVoteInfo
	)
	err := row.Scan(
		&feeVoteInfo.RecentBlockHash,
		&feeVoteInfo.RecentBlockHeight,
		&feeVoteInfo.FeeVote,
		&voterAddress,
		&voterSignature,
		&blockHeight,
	)
	vote.VoterAddress = voterAddress
	vote.BlockHeight = blockHeight
	vote.VoterSignature = voterSignature
	vote.VoteInfo = &feeVoteInfo
	return err
}

func (fvr *FeeVoteRevealVoteQuery) BuildModel(
	votes []*model.FeeVoteRevealVote, rows *sql.Rows,
) ([]*model.FeeVoteRevealVote, error) {
	for rows.Next() {
		var (
			revealVote = model.FeeVoteRevealVote{
				VoteInfo:       &model.FeeVoteInfo{},
				VoterSignature: nil,
				VoterAddress:   nil,
				BlockHeight:    0,
			}
			err error
		)
		err = rows.Scan(
			&revealVote.VoteInfo.RecentBlockHash,
			&revealVote.VoteInfo.RecentBlockHeight,
			&revealVote.VoteInfo.FeeVote,
			&revealVote.VoterAddress,
			&revealVote.VoterSignature,
			&revealVote.BlockHeight,
		)
		if err != nil {
			return nil, err
		}
		votes = append(votes, &revealVote)
	}
	return votes, nil
}

// Rollback delete records `WHERE block_height > "block_height"`
func (fvr *FeeVoteRevealVoteQuery) Rollback(height uint32) (multiQueries [][]interface{}) {
	return [][]interface{}{
		{
			fmt.Sprintf("DELETE FROM %s WHERE block_height > ?", fvr.getTableName()),
			height,
		},
	}
}

// SelectDataForSnapshot select only the block at snapshot block_height
func (fvr *FeeVoteRevealVoteQuery) SelectDataForSnapshot(fromHeight, toHeight uint32) string {
	return fmt.Sprintf(`SELECT %s FROM %s WHERE block_height >= %d AND block_height <= %d AND block_height != 0`,
		strings.Join(fvr.Fields, ", "), fvr.getTableName(), fromHeight, toHeight)
}

// TrimDataBeforeSnapshot delete entries to assure there are no duplicates before applying a snapshot
func (fvr *FeeVoteRevealVoteQuery) TrimDataBeforeSnapshot(fromHeight, toHeight uint32) string {
	return fmt.Sprintf(`DELETE FROM %s WHERE block_height >= %d AND block_height <= %d AND block_height != 0`,
		fvr.getTableName(), fromHeight, toHeight)
}
