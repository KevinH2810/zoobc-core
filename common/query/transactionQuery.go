package query

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zoobc/zoobc-core/common/chaintype"
	"github.com/zoobc/zoobc-core/common/model"
)

type (
	TransactionQueryInterface interface {
		InsertTransaction(tx *model.Transaction) (str string, args []interface{})
		GetTransaction(id int64) string
		GetTransactionsByIds(txIds []int64) (str string, args []interface{})
		GetTransactionsByBlockID(blockID int64) (str string, args []interface{})
		ExtractModel(tx *model.Transaction) []interface{}
		BuildModel(txs []*model.Transaction, rows *sql.Rows) ([]*model.Transaction, error)
		DeleteTransactions(id int64) string
		Scan(tx *model.Transaction, row *sql.Row) error
	}

	TransactionQuery struct {
		Fields    []string
		TableName string
		ChainType chaintype.ChainType
	}
)

// NewTransactionQuery returns TransactionQuery instance
func NewTransactionQuery(chaintype chaintype.ChainType) *TransactionQuery {
	return &TransactionQuery{
		Fields: []string{
			"id",
			"block_id",
			"block_height",
			"sender_account_address",
			"recipient_account_address",
			"transaction_type",
			"fee",
			"timestamp",
			"transaction_hash",
			"transaction_body_length",
			"transaction_body_bytes",
			"signature",
			"version",
			"transaction_index",
			"multisig_child",
		},
		TableName: "\"transaction\"",
		ChainType: chaintype,
	}
}

func (tq *TransactionQuery) getTableName() string {
	return tq.TableName
}

// GetTransaction get a single transaction from DB
func (tq *TransactionQuery) GetTransaction(id int64) string {
	query := fmt.Sprintf("SELECT %s from %s", strings.Join(tq.Fields, ", "), tq.getTableName())
	var queryParam = []string{"multisig_child = false"}
	if id != 0 {
		queryParam = append(queryParam, fmt.Sprintf("id = %d", id))
	}
	if len(queryParam) > 0 {
		query = query + " WHERE " + strings.Join(queryParam, " AND ")
	}
	return query
}

// InsertTransaction inserts a new transaction into DB
func (tq *TransactionQuery) InsertTransaction(tx *model.Transaction) (str string, args []interface{}) {
	var value = fmt.Sprintf("?%s", strings.Repeat(", ?", len(tq.Fields)-1))
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES(%s)",
		tq.getTableName(), strings.Join(tq.Fields, ", "), value)
	return query, tq.ExtractModel(tx)
}

func (tq *TransactionQuery) GetTransactionsByBlockID(blockID int64) (str string, args []interface{}) {
	query := fmt.Sprintf("SELECT %s FROM %s WHERE block_id = ? AND multisig_child = false "+
		"ORDER BY transaction_index ASC", strings.Join(tq.Fields, ", "), tq.getTableName())
	return query, []interface{}{blockID}
}

func (tq *TransactionQuery) GetTransactionsByIds(txIds []int64) (str string, args []interface{}) {
	var txIdsStr []string

	for _, txID := range txIds {
		txIdsStr = append(txIdsStr, fmt.Sprintf("%d", txID))
	}
	query := fmt.Sprintf("SELECT %s FROM %s WHERE multisig_child = false AND id in (%s)",
		strings.Join(tq.Fields, ", "), tq.getTableName(), strings.Join(txIdsStr, ","))
	return query, []interface{}{}
}

// DeleteTransactions. delete some transactions according to timestamp
func (tq *TransactionQuery) DeleteTransactions(id int64) string {
	return fmt.Sprintf("DELETE FROM %v WHERE height >= (SELECT height FROM %v WHERE ID = %v)", tq.getTableName(), tq.getTableName(), id)
}

// ExtractModel extract the model struct fields to the order of TransactionQuery.Fields
func (*TransactionQuery) ExtractModel(tx *model.Transaction) []interface{} {
	return []interface{}{
		&tx.ID,
		&tx.BlockID,
		&tx.Height,
		&tx.SenderAccountAddress,
		&tx.RecipientAccountAddress,
		&tx.TransactionType,
		&tx.Fee,
		&tx.Timestamp,
		&tx.TransactionHash,
		&tx.TransactionBodyLength,
		&tx.TransactionBodyBytes,
		&tx.Signature,
		&tx.Version,
		&tx.TransactionIndex,
		&tx.MultisigChild,
	}
}

func (*TransactionQuery) BuildModel(txs []*model.Transaction, rows *sql.Rows) ([]*model.Transaction, error) {
	for rows.Next() {
		var (
			tx  model.Transaction
			err error
		)
		err = rows.Scan(
			&tx.ID,
			&tx.BlockID,
			&tx.Height,
			&tx.SenderAccountAddress,
			&tx.RecipientAccountAddress,
			&tx.TransactionType,
			&tx.Fee,
			&tx.Timestamp,
			&tx.TransactionHash,
			&tx.TransactionBodyLength,
			&tx.TransactionBodyBytes,
			&tx.Signature,
			&tx.Version,
			&tx.TransactionIndex,
			&tx.MultisigChild,
		)
		if err != nil {
			return nil, err
		}
		txs = append(txs, &tx)
	}
	return txs, nil
}

func (*TransactionQuery) Scan(tx *model.Transaction, row *sql.Row) error {
	err := row.Scan(
		&tx.ID,
		&tx.BlockID,
		&tx.Height,
		&tx.SenderAccountAddress,
		&tx.RecipientAccountAddress,
		&tx.TransactionType,
		&tx.Fee,
		&tx.Timestamp,
		&tx.TransactionHash,
		&tx.TransactionBodyLength,
		&tx.TransactionBodyBytes,
		&tx.Signature,
		&tx.Version,
		&tx.TransactionIndex,
		&tx.MultisigChild,
	)
	return err
}

// Rollback delete records `WHERE height > "height"
func (tq *TransactionQuery) Rollback(height uint32) (multiQueries [][]interface{}) {
	return [][]interface{}{
		{
			fmt.Sprintf("DELETE FROM %s WHERE block_height > ?", tq.getTableName()),
			height,
		},
	}
}
