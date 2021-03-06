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
package service

import (
	"database/sql"
	"errors"
	"reflect"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/query"
)

func TestNewParticipationScoreService(t *testing.T) {
	type args struct {
		participationScoreQuery query.ParticipationScoreQueryInterface
		queryExecutor           query.ExecutorInterface
	}
	tests := []struct {
		name string
		args args
		want *ParticipationScoreService
	}{
		{
			name: "NewParticipationScore",
			args: args{
				participationScoreQuery: nil,
				queryExecutor:           nil,
			},
			want: &ParticipationScoreService{
				ParticipationScoreQuery: nil,
				QueryExecutor:           nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewParticipationScoreService(tt.args.participationScoreQuery, tt.args.queryExecutor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewParticipationScoreService() = %v, want %v", got, tt.want)
			}
		})
	}
}

type (
	// GetParticipationScore mocks

	// GetParticipationScoreByBlockHeightRange mocks
	mockGetParticipationScoreByBlockHeightRangeExecutorFail struct {
		query.Executor
	}
	mockGetParticipationScoreByBlockHeightRangeExecutorSuccess struct {
		query.Executor
	}
	mockGetParticipationScoreByBlockHeightRangeParticipationScoreQuerySuccess struct {
		query.ParticipationScoreQuery
	}
	mockGetParticipationScoreByBlockHeightRangeParticipationScoreQueryFail struct {
		query.ParticipationScoreQuery
	}
	// GetParticipationScoreByBlockHeightRange mocks
)

var (
	// GetParticipationScore mocks
	mockGetParticipationScoreResult = &model.ParticipationScore{
		Score: 1000,
	}
	// GetParticipationScore mocks
)

func (*mockGetParticipationScoreByBlockHeightRangeExecutorFail) ExecuteSelect(query string, tx bool, args ...interface{}) (*sql.Rows, error) {
	return nil, errors.New("mockError")
}

func (*mockGetParticipationScoreByBlockHeightRangeExecutorSuccess) ExecuteSelect(query string, tx bool, args ...interface{}) (*sql.Rows, error) {
	db, mock, _ := sqlmock.New()
	defer db.Close()
	mock.ExpectQuery(regexp.QuoteMeta("MOCKQUERY")).WillReturnRows(sqlmock.NewRows([]string{
		"dummyColumn"}).AddRow(
		[]byte{1}))
	rows, _ := db.Query("MOCKQUERY")
	return rows, nil
}

func (*mockGetParticipationScoreByBlockHeightRangeParticipationScoreQuerySuccess) BuildModel(
	participationScores []*model.ParticipationScore, rows *sql.Rows) ([]*model.ParticipationScore, error) {
	return []*model.ParticipationScore{
		mockGetParticipationScoreResult,
	}, nil
}

func (*mockGetParticipationScoreByBlockHeightRangeParticipationScoreQueryFail) BuildModel(
	participationScores []*model.ParticipationScore, rows *sql.Rows) ([]*model.ParticipationScore, error) {
	return nil, errors.New("mockError")
}

func TestParticipationScoreService_GetParticipationScoreByBlockHeightRange(t *testing.T) {
	type fields struct {
		ParticipationScoreQuery query.ParticipationScoreQueryInterface
		QueryExecutor           query.ExecutorInterface
	}
	type args struct {
		fromBlockHeight uint32
		toBlockHeight   uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.ParticipationScore
		wantErr bool
	}{
		{
			name: "wantError:ExecuteSelectError",
			fields: fields{
				QueryExecutor:           &mockGetParticipationScoreByBlockHeightRangeExecutorFail{},
				ParticipationScoreQuery: &mockGetParticipationScoreByBlockHeightRangeParticipationScoreQuerySuccess{},
			},
			args: args{
				fromBlockHeight: 1,
				toBlockHeight:   2,
			},
			wantErr: true,
		},
		{
			name: "wantError:BuildModelError",
			fields: fields{
				QueryExecutor:           &mockGetParticipationScoreByBlockHeightRangeExecutorSuccess{},
				ParticipationScoreQuery: &mockGetParticipationScoreByBlockHeightRangeParticipationScoreQueryFail{},
			},
			args: args{
				fromBlockHeight: 1,
				toBlockHeight:   2,
			},
			wantErr: true,
		},
		{
			name: "wantSuccess",
			fields: fields{
				QueryExecutor:           &mockGetParticipationScoreByBlockHeightRangeExecutorSuccess{},
				ParticipationScoreQuery: &mockGetParticipationScoreByBlockHeightRangeParticipationScoreQuerySuccess{},
			},
			args: args{
				fromBlockHeight: 1,
				toBlockHeight:   2,
			},
			want: []*model.ParticipationScore{
				mockGetParticipationScoreResult,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pss := &ParticipationScoreService{
				ParticipationScoreQuery: tt.fields.ParticipationScoreQuery,
				QueryExecutor:           tt.fields.QueryExecutor,
			}
			got, err := pss.GetParticipationScoreByBlockHeightRange(tt.args.fromBlockHeight, tt.args.toBlockHeight)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParticipationScoreService.GetParticipationScoreByBlockHeightRange() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParticipationScoreService.GetParticipationScoreByBlockHeightRange() = %v, want %v", got, tt.want)
			}
		})
	}
}

var (
	mockParticipationQueryGetLatestSuccess = &model.ParticipationScore{
		NodeID: 100,
		Score:  1000,
		Latest: true,
		Height: 1001,
	}
)

type (
	mockExecutorGetLatestParticipationScoreByNodeIDSuccess struct {
		query.Executor
	}
	mockExecutorGetLatestParticipationScoreByNodeIDFail struct {
		query.Executor
	}
)

func (*mockExecutorGetLatestParticipationScoreByNodeIDFail) ExecuteSelectRow(
	query string, tx bool, args ...interface{},
) (*sql.Row, error) {
	return nil, errors.New("mockedError")
}

func (*mockExecutorGetLatestParticipationScoreByNodeIDSuccess) ExecuteSelectRow(
	qStr string, tx bool, args ...interface{},
) (*sql.Row, error) {
	dbMocked, mock, _ := sqlmock.New()
	mockedRows := mock.NewRows(query.NewParticipationScoreQuery().Fields)
	mockedRows.AddRow(
		mockParticipationQueryGetLatestSuccess.NodeID,
		mockParticipationQueryGetLatestSuccess.Score,
		mockParticipationQueryGetLatestSuccess.Latest,
		mockParticipationQueryGetLatestSuccess.Height,
	)

	mock.ExpectQuery(regexp.QuoteMeta(qStr)).WillReturnRows(mockedRows)
	return dbMocked.QueryRow(qStr), nil
}

func TestParticipationScoreService_GetLatestParticipationScoreByNodeID(t *testing.T) {
	type fields struct {
		ParticipationScoreQuery query.ParticipationScoreQueryInterface
		QueryExecutor           query.ExecutorInterface
	}
	type args struct {
		nodeID int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.ParticipationScore
		wantErr bool
	}{
		{
			name: "getLatestParticipationScoreByNodeID - success",
			fields: fields{
				ParticipationScoreQuery: query.NewParticipationScoreQuery(),
				QueryExecutor:           &mockExecutorGetLatestParticipationScoreByNodeIDSuccess{},
			},
			args: args{
				nodeID: mockParticipationQueryGetLatestSuccess.NodeID,
			},
			want:    mockParticipationQueryGetLatestSuccess,
			wantErr: false,
		},
		{
			name: "getLatestParticipationScoreByNodeID - error executor",
			fields: fields{
				ParticipationScoreQuery: query.NewParticipationScoreQuery(),
				QueryExecutor:           &mockExecutorGetLatestParticipationScoreByNodeIDFail{},
			},
			args: args{
				nodeID: mockParticipationQueryGetLatestSuccess.NodeID,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pss := &ParticipationScoreService{
				ParticipationScoreQuery: tt.fields.ParticipationScoreQuery,
				QueryExecutor:           tt.fields.QueryExecutor,
			}
			got, err := pss.GetLatestParticipationScoreByNodeID(tt.args.nodeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetLatestParticipationScoreByNodeID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetLatestParticipationScoreByNodeID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
