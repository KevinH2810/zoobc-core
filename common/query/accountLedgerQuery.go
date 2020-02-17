package query

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zoobc/zoobc-core/common/model"
)

type (

	// AccountLedgerQuery schema of AccountLedger
	AccountLedgerQuery struct {
		Fields    []string
		TableName string
	}
	// AccountLedgerQueryInterface includes interface methods for AccountLedgerQuery
	AccountLedgerQueryInterface interface {
		ExtractModel(accountLedger *model.AccountLedger) []interface{}
		BuildModel(accountLedgers []*model.AccountLedger, rows *sql.Rows) ([]*model.AccountLedger, error)
		InsertAccountLedger(accountLedger *model.AccountLedger) (qStr string, args []interface{})
	}
)

// NewAccountLedgerQuery func that return AccountLedger schema with value
func NewAccountLedgerQuery() *AccountLedgerQuery {
	return &AccountLedgerQuery{
		Fields: []string{
			"account_address",
			"balance_change",
			"block_height",
			"transaction_id",
			"event_type",
			"timestamp",
		},
		TableName: "account_ledger",
	}
}

func (q *AccountLedgerQuery) getTableName() interface{} {
	return q.TableName
}

// InsertAccountLedger represents insert query for AccountLedger
func (q *AccountLedgerQuery) InsertAccountLedger(accountLedger *model.AccountLedger) (qStr string, args []interface{}) {
	return fmt.Sprintf(
			"INSERT INTO %s (%s) VALUES(%s)",
			q.getTableName(),
			strings.Join(q.Fields, ", "),
			fmt.Sprintf("? %s", strings.Repeat(", ?", len(q.Fields)-1)),
		),
		q.ExtractModel(accountLedger)
}

// ExtractModel will extract accountLedger model to []interface
func (*AccountLedgerQuery) ExtractModel(accountLedger *model.AccountLedger) []interface{} {
	return []interface{}{
		accountLedger.GetAccountAddress(),
		accountLedger.GetBalanceChange(),
		accountLedger.GetBlockHeight(),
		accountLedger.GetTransactionID(),
		accountLedger.GetEventType(),
		accountLedger.GetTimestamp(),
	}
}

// BuildModel will create or build models that extracted from rows
func (*AccountLedgerQuery) BuildModel(accountLedgers []*model.AccountLedger, rows *sql.Rows) ([]*model.AccountLedger, error) {
	for rows.Next() {
		var (
			accountLedger model.AccountLedger
			err           error
		)
		err = rows.Scan(
			&accountLedger.AccountAddress,
			&accountLedger.BalanceChange,
			&accountLedger.BlockHeight,
			&accountLedger.TransactionID,
			&accountLedger.EventType,
			&accountLedger.Timestamp,
		)
		if err != nil {
			return nil, err
		}
		accountLedgers = append(accountLedgers, &accountLedger)
	}
	return accountLedgers, nil
}

// Rollback represents delete query in block_height n
func (q *AccountLedgerQuery) Rollback(height uint32) (multiQueries [][]interface{}) {
	return [][]interface{}{
		{
			fmt.Sprintf("DELETE FROM %s WHERE block_height > ?", q.getTableName()),
			height,
		},
	}
}