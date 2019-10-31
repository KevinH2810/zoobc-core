package query

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/zoobc/zoobc-core/common/constant"
	"github.com/zoobc/zoobc-core/common/model"
)

type (
	NodeReceiptQueryInterface interface {
		InsertReceipt(receipt *model.Receipt) (str string, args []interface{})
		InsertReceipts(receipts []*model.Receipt) (str string, args []interface{})
		GetReceipts(paginate model.Pagination) string
		GetReceiptByRoot(root []byte) (str string, args []interface{})
		GetReceiptsWithUniqueRecipient(limit uint32, offset uint64, rmrLinked bool) string
		SelectReceipt(lowerHeight, upperHeight, limit uint32) (str string)
		ExtractModel(receipt *model.Receipt) []interface{}
		BuildModel(receipts []*model.Receipt, rows *sql.Rows) ([]*model.Receipt, error)
		Scan(receipt *model.Receipt, row *sql.Row) error
	}

	NodeReceiptQuery struct {
		Fields    []string
		TableName string
	}
)

// NewTransactionQuery returns TransactionQuery instance
func NewNodeReceiptQuery() *NodeReceiptQuery {
	return &NodeReceiptQuery{
		Fields: []string{
			"sender_public_key",
			"recipient_public_key",
			"datum_type",
			"datum_hash",
			"reference_block_height",
			"reference_block_hash",
			"rmr_linked",
			"recipient_signature",
			"rmr",
			"rmr_index",
		},
		TableName: "node_receipt",
	}
}

func (rq *NodeReceiptQuery) getTableName() string {
	return rq.TableName
}

// GetReceipts get a set of receipts that satisfies the params from DB
func (rq *NodeReceiptQuery) GetReceipts(paginate model.Pagination) string {

	query := fmt.Sprintf(
		"SELECT %s FROM %s ",
		strings.Join(rq.Fields, ", "),
		rq.getTableName(),
	)

	newLimit := paginate.GetLimit()
	if newLimit == 0 {
		newLimit = constant.ReceiptNodeMaximum
	}

	orderField := paginate.GetOrderField()
	if orderField == "" {
		orderField = "reference_block_height"
	}

	query += fmt.Sprintf(
		"ORDER BY %s %s LIMIT %d OFFSET %d",
		orderField,
		paginate.GetOrderBy(),
		newLimit,
		paginate.GetPage(),
	)
	return query
}

func (rq *NodeReceiptQuery) GetReceiptsWithUniqueRecipient(limit uint32, offset uint64, rmrLinked bool) string {
	var query string
	if limit == 0 {
		limit = 10
	}
	if rmrLinked {
		query = fmt.Sprintf("SELECT %s FROM %s AS rc WHERE rmr_linked IS NOT NULL AND "+
			"NOT EXISTS (SELECT datum_hash FROM published_receipt AS pr WHERE pr.datum_hash == rc.datum_hash) "+
			"GROUP BY recipient_public_key LIMIT %d, %d",
			strings.Join(rq.Fields, ", "), rq.getTableName(), offset, limit)
	} else {
		query = fmt.Sprintf("SELECT %s FROM %s AS rc WHERE "+
			"NOT EXISTS (SELECT datum_hash FROM published_receipt AS pr WHERE pr.datum_hash == rc.datum_hash) "+
			"GROUP BY recipient_public_key LIMIT %d, %d",
			strings.Join(rq.Fields, ", "), rq.getTableName(), offset, limit)
	}
	return query
}

// GetReceiptByRoot return sql query to fetch receipts by its merkle root, the datum_hash should not already exists in
// published_receipt table
func (rq *NodeReceiptQuery) GetReceiptByRoot(root []byte) (str string, args []interface{}) {
	query := fmt.Sprintf("SELECT %s FROM %s AS rc WHERE rc.rmr = ? AND "+
		"NOT EXISTS (SELECT datum_hash FROM published_receipt AS pr WHERE "+
		"pr.datum_hash = rc.datum_hash AND pr.recipient_public_key = rc.recipient_public_key) "+
		"GROUP BY recipient_public_key",
		strings.Join(rq.Fields, ", "), rq.getTableName())
	return query, []interface{}{
		root,
	}
}

// SelectReceipt select list of receipt by some filter
func (rq *NodeReceiptQuery) SelectReceipt(
	lowerHeight, upperHeight, limit uint32,
) (str string) {
	query := fmt.Sprintf("SELECT %s FROM %s AS nr WHERE EXISTS "+
		"(SELECT rmr_linked FROM published_receipt AS pr WHERE nr.rmr = pr.rmr_linked AND "+
		"block_height >= %d AND block_height <= %d ) LIMIT %d",
		strings.Join(rq.Fields, ", "), rq.getTableName(), lowerHeight, upperHeight, limit)

	return query
}

// InsertReceipts inserts a new receipts into DB
func (rq *NodeReceiptQuery) InsertReceipt(receipt *model.Receipt) (str string, args []interface{}) {

	return fmt.Sprintf(
		"INSERT INTO %s (%s) VALUES(%s)",
		rq.getTableName(),
		strings.Join(rq.Fields, ", "),
		fmt.Sprintf("? %s", strings.Repeat(", ? ", len(rq.Fields)-1)),
	), rq.ExtractModel(receipt)
}

// InsertReceipts build query for bulk store receipts
func (rq *NodeReceiptQuery) InsertReceipts(receipts []*model.Receipt) (qStr string, args []interface{}) {

	var (
		query  string
		values []interface{}
	)

	query = fmt.Sprintf(
		"INSERT INTO %s (%s) ",
		rq.getTableName(),
		strings.Join(rq.Fields, ", "),
	)

	for k, receipt := range receipts {
		query += fmt.Sprintf("VALUES(?%s)", strings.Repeat(",? ", len(rq.Fields)-1))
		if k < len(receipts)-1 {
			query += ", "
		}
		values = append(values, rq.ExtractModel(receipt)...)
	}
	return query, values
}

// ExtractModel extract the model struct fields to the order of NodeReceiptQuery.Fields
func (*NodeReceiptQuery) ExtractModel(receipt *model.Receipt) []interface{} {
	return []interface{}{
		&receipt.BatchReceipt.SenderPublicKey,
		&receipt.BatchReceipt.RecipientPublicKey,
		&receipt.BatchReceipt.DatumType,
		&receipt.BatchReceipt.DatumHash,
		&receipt.BatchReceipt.ReferenceBlockHeight,
		&receipt.BatchReceipt.ReferenceBlockHash,
		&receipt.BatchReceipt.RMRLinked,
		&receipt.BatchReceipt.RecipientSignature,
		&receipt.RMR,
		&receipt.RMRIndex,
	}
}

func (*NodeReceiptQuery) BuildModel(receipts []*model.Receipt, rows *sql.Rows) ([]*model.Receipt, error) {

	for rows.Next() {
		var (
			receipt      model.Receipt
			batchReceipt model.BatchReceipt
			err          error
		)

		err = rows.Scan(
			&batchReceipt.SenderPublicKey,
			&batchReceipt.RecipientPublicKey,
			&batchReceipt.DatumType,
			&batchReceipt.DatumHash,
			&batchReceipt.ReferenceBlockHeight,
			&batchReceipt.ReferenceBlockHash,
			&batchReceipt.RMRLinked,
			&batchReceipt.RecipientSignature,
			&receipt.RMR,
			&receipt.RMRIndex,
		)
		if err != nil {
			return nil, err
		}
		receipt.BatchReceipt = &batchReceipt
		receipts = append(receipts, &receipt)
	}

	return receipts, nil
}

func (*NodeReceiptQuery) Scan(receipt *model.Receipt, row *sql.Row) error {

	err := row.Scan(
		&receipt.BatchReceipt.SenderPublicKey,
		&receipt.BatchReceipt.RecipientPublicKey,
		&receipt.BatchReceipt.DatumType,
		&receipt.BatchReceipt.DatumHash,
		&receipt.BatchReceipt.ReferenceBlockHeight,
		&receipt.BatchReceipt.ReferenceBlockHash,
		&receipt.BatchReceipt.RMRLinked,
		&receipt.BatchReceipt.RecipientSignature,
		&receipt.RMR,
		&receipt.RMRIndex,
	)
	return err

}