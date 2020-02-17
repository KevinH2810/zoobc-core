package transaction

import (
	"database/sql"
	"encoding/binary"
	"reflect"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/query"
)

func TestApprovalEscrowTransaction_GetBodyBytes(t *testing.T) {
	type fields struct {
		ID                  int64
		Fee                 int64
		SenderAddress       string
		Height              uint32
		Body                *model.ApprovalEscrowTransactionBody
		Escrow              *model.Escrow
		AccountBalanceQuery query.AccountBalanceQueryInterface
		QueryExecutor       query.ExecutorInterface
		AccountLedgerQuery  query.AccountLedgerQueryInterface
		EscrowQuery         query.EscrowTransactionQueryInterface
		TransactionQuery    query.TransactionQueryInterface
		TypeActionSwitcher  TypeActionSwitcher
	}
	tests := []struct {
		name   string
		fields fields
		want   []byte
	}{
		{
			name: "wantSuccess",
			fields: fields{
				ID:            0,
				Fee:           0,
				SenderAddress: "",
				Height:        0,
				Body: &model.ApprovalEscrowTransactionBody{
					Approval:      1,
					TransactionID: 120978123123,
				},
			},
			want: []byte{1, 0, 0, 0, 115, 169, 219, 42, 28, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ApprovalEscrowTransaction{
				ID:                  tt.fields.ID,
				Fee:                 tt.fields.Fee,
				SenderAddress:       tt.fields.SenderAddress,
				Height:              tt.fields.Height,
				Body:                tt.fields.Body,
				Escrow:              tt.fields.Escrow,
				AccountBalanceQuery: tt.fields.AccountBalanceQuery,
				QueryExecutor:       tt.fields.QueryExecutor,
				AccountLedgerQuery:  tt.fields.AccountLedgerQuery,
				EscrowQuery:         tt.fields.EscrowQuery,
				TransactionQuery:    tt.fields.TransactionQuery,
				TypeActionSwitcher:  tt.fields.TypeActionSwitcher,
			}
			if got := tx.GetBodyBytes(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetBodyBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApprovalEscrowTransaction_ParseBodyBytes(t *testing.T) {
	type fields struct {
		ID                  int64
		Fee                 int64
		SenderAddress       string
		Height              uint32
		Body                *model.ApprovalEscrowTransactionBody
		Escrow              *model.Escrow
		AccountBalanceQuery query.AccountBalanceQueryInterface
		QueryExecutor       query.ExecutorInterface
		AccountLedgerQuery  query.AccountLedgerQueryInterface
		EscrowQuery         query.EscrowTransactionQueryInterface
		TransactionQuery    query.TransactionQueryInterface
		TypeActionSwitcher  TypeActionSwitcher
	}
	type args struct {
		bodyBytes []byte
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    model.TransactionBodyInterface
		wantErr bool
	}{
		{
			name: "wantSuccess",
			fields: fields{
				ID:            0,
				Fee:           0,
				SenderAddress: "",
				Height:        0,
				Body: &model.ApprovalEscrowTransactionBody{
					Approval:      1,
					TransactionID: 120978123123,
				},
			},
			args: args{bodyBytes: []byte{1, 0, 0, 0, 115, 169, 219, 42, 28, 0, 0, 0}},
			want: &model.ApprovalEscrowTransactionBody{
				Approval:      1,
				TransactionID: 120978123123,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ApprovalEscrowTransaction{
				ID:                  tt.fields.ID,
				Fee:                 tt.fields.Fee,
				SenderAddress:       tt.fields.SenderAddress,
				Height:              tt.fields.Height,
				Body:                tt.fields.Body,
				Escrow:              tt.fields.Escrow,
				AccountBalanceQuery: tt.fields.AccountBalanceQuery,
				QueryExecutor:       tt.fields.QueryExecutor,
				AccountLedgerQuery:  tt.fields.AccountLedgerQuery,
				EscrowQuery:         tt.fields.EscrowQuery,
				TransactionQuery:    tt.fields.TransactionQuery,
				TypeActionSwitcher:  tt.fields.TypeActionSwitcher,
			}
			got, err := tx.ParseBodyBytes(tt.args.bodyBytes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseBodyBytes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseBodyBytes() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type (
	mockQueryExecutorValidate struct {
		query.Executor
	}
	mockQueryExecutorValidateNotFound struct {
		query.Executor
	}
	mockAccountBalanceQueryValidateNotFound struct {
		query.AccountBalanceQuery
	}
	mockAccountBalanceQueryValidateFound struct {
		query.AccountBalanceQuery
	}
)

func (*mockQueryExecutorValidate) ExecuteSelectRow(qStr string, tx bool, args ...interface{}) (*sql.Row, error) {
	db, mock, _ := sqlmock.New()
	mockRow := mock.NewRows(query.NewEscrowTransactionQuery().Fields)
	mockRow.AddRow(
		120978123123,
		"ABC",
		"DEF",
		"GHI",
		1,
		10,
		100,
		0,
		1,
		true,
		"",
	)
	mock.ExpectQuery("").WillReturnRows(mockRow)
	mockedRow := db.QueryRow("")
	return mockedRow, nil
}
func (*mockQueryExecutorValidateNotFound) ExecuteSelectRow(qStr string, tx bool, args ...interface{}) (*sql.Row, error) {
	db, mock, _ := sqlmock.New()
	mockRow := mock.NewRows(query.NewEscrowTransactionQuery().Fields)
	mock.ExpectQuery("").WillReturnRows(mockRow)
	mockedRow := db.QueryRow("")
	return mockedRow, nil
}

func (*mockAccountBalanceQueryValidateNotFound) GetAccountBalanceByAccountAddress(sender string) (qStr string, args []interface{}) {
	return "", []interface{}{}
}
func (*mockAccountBalanceQueryValidateNotFound) Scan(accountBalance *model.AccountBalance, row *sql.Row) error {
	return sql.ErrNoRows
}
func (*mockAccountBalanceQueryValidateFound) GetAccountBalanceByAccountAddress(sender string) (qStr string, args []interface{}) {
	return "", []interface{}{}
}
func (*mockAccountBalanceQueryValidateFound) Scan(accountBalance *model.AccountBalance, row *sql.Row) error {
	accountBalance.AccountAddress = "GHI"
	accountBalance.Balance = 1000
	accountBalance.Latest = true

	return nil
}
func TestApprovalEscrowTransaction_Validate(t *testing.T) {
	type fields struct {
		ID                  int64
		Fee                 int64
		SenderAddress       string
		Height              uint32
		Body                *model.ApprovalEscrowTransactionBody
		Escrow              *model.Escrow
		AccountBalanceQuery query.AccountBalanceQueryInterface
		QueryExecutor       query.ExecutorInterface
		AccountLedgerQuery  query.AccountLedgerQueryInterface
		EscrowQuery         query.EscrowTransactionQueryInterface
		TransactionQuery    query.TransactionQueryInterface
		TypeActionSwitcher  TypeActionSwitcher
	}
	type args struct {
		dbTx bool
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "wantError:NotFound",
			fields: fields{
				ID:            0,
				Fee:           0,
				SenderAddress: "GHI",
				Height:        0,
				Body: &model.ApprovalEscrowTransactionBody{
					Approval:      1,
					TransactionID: 120978123123,
				},
				QueryExecutor: &mockQueryExecutorValidateNotFound{},
				EscrowQuery:   query.NewEscrowTransactionQuery(),
			},
			args:    args{dbTx: false},
			wantErr: true,
		},
		{
			name: "wantError:InvalidTransactionID",
			fields: fields{
				ID:            0,
				Fee:           0,
				SenderAddress: "GHI",
				Height:        0,
				Body: &model.ApprovalEscrowTransactionBody{
					Approval:      1,
					TransactionID: 0,
				},
				AccountBalanceQuery: nil,
				QueryExecutor:       &mockQueryExecutorValidate{},
				AccountLedgerQuery:  nil,
				EscrowQuery:         query.NewEscrowTransactionQuery(),
				TransactionQuery:    nil,
			},
			args:    args{dbTx: false},
			wantErr: true,
		},
		{
			name: "wantError:AccountBalanceNotFound",
			fields: fields{
				ID:            0,
				Fee:           0,
				SenderAddress: "GHI",
				Height:        0,
				Body: &model.ApprovalEscrowTransactionBody{
					Approval:      1,
					TransactionID: 120978123123,
				},
				AccountBalanceQuery: &mockAccountBalanceQueryValidateNotFound{},
				QueryExecutor:       &mockQueryExecutorValidate{},
				EscrowQuery:         query.NewEscrowTransactionQuery(),
			},
			args:    args{dbTx: false},
			wantErr: true,
		},
		{
			name: "wantSuccess",
			fields: fields{
				ID:            0,
				Fee:           0,
				SenderAddress: "GHI",
				Height:        0,
				Body: &model.ApprovalEscrowTransactionBody{
					Approval:      1,
					TransactionID: 120978123123,
				},
				AccountBalanceQuery: &mockAccountBalanceQueryValidateFound{},
				QueryExecutor:       &mockQueryExecutorValidate{},
				EscrowQuery:         query.NewEscrowTransactionQuery(),
			},
			args: args{dbTx: false},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ApprovalEscrowTransaction{
				ID:                  tt.fields.ID,
				Fee:                 tt.fields.Fee,
				SenderAddress:       tt.fields.SenderAddress,
				Height:              tt.fields.Height,
				Body:                tt.fields.Body,
				Escrow:              tt.fields.Escrow,
				AccountBalanceQuery: tt.fields.AccountBalanceQuery,
				QueryExecutor:       tt.fields.QueryExecutor,
				AccountLedgerQuery:  tt.fields.AccountLedgerQuery,
				EscrowQuery:         tt.fields.EscrowQuery,
				TransactionQuery:    tt.fields.TransactionQuery,
				TypeActionSwitcher:  tt.fields.TypeActionSwitcher,
			}
			if err := tx.Validate(tt.args.dbTx); (err != nil) != tt.wantErr {
				t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type (
	mockQueryExecutorUnconfirmed struct {
		query.Executor
	}
)

func (*mockQueryExecutorUnconfirmed) ExecuteTransaction(qStr string, args ...interface{}) error {
	return nil
}
func TestApprovalEscrowTransaction_ApplyUnconfirmed(t *testing.T) {
	type fields struct {
		ID                  int64
		Fee                 int64
		SenderAddress       string
		Height              uint32
		Body                *model.ApprovalEscrowTransactionBody
		Escrow              *model.Escrow
		AccountBalanceQuery query.AccountBalanceQueryInterface
		QueryExecutor       query.ExecutorInterface
		AccountLedgerQuery  query.AccountLedgerQueryInterface
		EscrowQuery         query.EscrowTransactionQueryInterface
		TransactionQuery    query.TransactionQueryInterface
		TypeActionSwitcher  TypeActionSwitcher
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "wantSuccess",
			fields: fields{
				ID:                  0,
				Fee:                 1,
				SenderAddress:       "",
				Height:              0,
				Body:                nil,
				Escrow:              nil,
				AccountBalanceQuery: query.NewAccountBalanceQuery(),
				QueryExecutor:       &mockQueryExecutorUnconfirmed{},
				AccountLedgerQuery:  nil,
				EscrowQuery:         nil,
				TransactionQuery:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ApprovalEscrowTransaction{
				ID:                  tt.fields.ID,
				Fee:                 tt.fields.Fee,
				SenderAddress:       tt.fields.SenderAddress,
				Height:              tt.fields.Height,
				Body:                tt.fields.Body,
				Escrow:              tt.fields.Escrow,
				AccountBalanceQuery: tt.fields.AccountBalanceQuery,
				QueryExecutor:       tt.fields.QueryExecutor,
				AccountLedgerQuery:  tt.fields.AccountLedgerQuery,
				EscrowQuery:         tt.fields.EscrowQuery,
				TransactionQuery:    tt.fields.TransactionQuery,
				TypeActionSwitcher:  tt.fields.TypeActionSwitcher,
			}
			if err := tx.ApplyUnconfirmed(); (err != nil) != tt.wantErr {
				t.Errorf("ApplyUnconfirmed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestApprovalEscrowTransaction_UndoApplyUnconfirmed(t *testing.T) {
	type fields struct {
		ID                  int64
		Fee                 int64
		SenderAddress       string
		Height              uint32
		Body                *model.ApprovalEscrowTransactionBody
		Escrow              *model.Escrow
		AccountBalanceQuery query.AccountBalanceQueryInterface
		QueryExecutor       query.ExecutorInterface
		AccountLedgerQuery  query.AccountLedgerQueryInterface
		EscrowQuery         query.EscrowTransactionQueryInterface
		TransactionQuery    query.TransactionQueryInterface
		TypeActionSwitcher  TypeActionSwitcher
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "wantSuccess",
			fields: fields{
				ID:                  0,
				Fee:                 1,
				SenderAddress:       "",
				Height:              0,
				Body:                nil,
				Escrow:              nil,
				AccountBalanceQuery: query.NewAccountBalanceQuery(),
				QueryExecutor:       &mockQueryExecutorUnconfirmed{},
				AccountLedgerQuery:  nil,
				EscrowQuery:         nil,
				TransactionQuery:    nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ApprovalEscrowTransaction{
				ID:                  tt.fields.ID,
				Fee:                 tt.fields.Fee,
				SenderAddress:       tt.fields.SenderAddress,
				Height:              tt.fields.Height,
				Body:                tt.fields.Body,
				Escrow:              tt.fields.Escrow,
				AccountBalanceQuery: tt.fields.AccountBalanceQuery,
				QueryExecutor:       tt.fields.QueryExecutor,
				AccountLedgerQuery:  tt.fields.AccountLedgerQuery,
				EscrowQuery:         tt.fields.EscrowQuery,
				TransactionQuery:    tt.fields.TransactionQuery,
				TypeActionSwitcher:  tt.fields.TypeActionSwitcher,
			}
			if err := tx.UndoApplyUnconfirmed(); (err != nil) != tt.wantErr {
				t.Errorf("UndoApplyUnconfirmed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

type (
	mockEscrowQueryApplyConfirmedOK struct {
		query.EscrowTransactionQuery
	}
	mockTransactionQueryApplyConfirmedOK struct {
		query.TransactionQuery
	}
	mockEscrowQueryExecutorApplyConfirmedOK struct {
		query.Executor
	}
)

// func (*mockTypeActionSwitcherApplyConfirmedOK) GetTransactionType(tx *model.Transaction) (TypeAction, error) {
// 	return
// }
func (*mockEscrowQueryExecutorApplyConfirmedOK) ExecuteSelectRow(string, bool, ...interface{}) (*sql.Row, error) {
	return &sql.Row{}, nil
}
func (*mockEscrowQueryExecutorApplyConfirmedOK) ExecuteTransactions(queries [][]interface{}) error {
	return nil
}
func (*mockEscrowQueryApplyConfirmedOK) GetLatestEscrowTransactionByID(int64) (qStr string, args []interface{}) {
	return "", []interface{}{}
}
func (*mockTransactionQueryApplyConfirmedOK) GetTransaction(int64) string {
	return ""
}
func (*mockTransactionQueryApplyConfirmedOK) Scan(tx *model.Transaction, row *sql.Row) error {
	tx.ID = -1273123123
	tx.BlockID = -123123123123
	tx.Version = 1
	tx.Height = 1
	tx.SenderAccountAddress = "BCZnSfqpP5tqFQlMTYkDeBVFWnbyVK7vLr5ORFpTjgtN"
	tx.RecipientAccountAddress = "BCZD_VxfO2S9aziIL3cn_cXW7uPDVPOrnXuP98GEAUC7"
	tx.TransactionType = binary.LittleEndian.Uint32([]byte{4, 0, 0, 0})
	tx.Fee = 1
	tx.Timestamp = 10000
	tx.TransactionHash = make([]byte, 200)
	tx.TransactionBodyLength = 88
	tx.TransactionBodyBytes = make([]byte, 88)
	tx.Signature = make([]byte, 68)
	tx.TransactionIndex = 1

	return nil
}
func (*mockEscrowQueryApplyConfirmedOK) Scan(escrow *model.Escrow, _ *sql.Row) error {
	escrow.ID = 1
	escrow.SenderAddress = "BCZnSfqpP5tqFQlMTYkDeBVFWnbyVK7vLr5ORFpTjgtN"
	escrow.RecipientAddress = "BCZD_VxfO2S9aziIL3cn_cXW7uPDVPOrnXuP98GEAUC7"
	escrow.ApproverAddress = "BCZKLvgUYZ1KKx-jtF9KoJskjVPvB9jpIjfzzI6zDW0J"
	escrow.Amount = 10
	escrow.Commission = 1
	escrow.Timeout = 120
	escrow.Status = 1
	escrow.BlockHeight = 0
	escrow.Latest = true
	return nil
}

func TestApprovalEscrowTransaction_ApplyConfirmed(t *testing.T) {
	type fields struct {
		ID                  int64
		Fee                 int64
		SenderAddress       string
		Height              uint32
		Body                *model.ApprovalEscrowTransactionBody
		Escrow              *model.Escrow
		AccountBalanceQuery query.AccountBalanceQueryInterface
		QueryExecutor       query.ExecutorInterface
		AccountLedgerQuery  query.AccountLedgerQueryInterface
		EscrowQuery         query.EscrowTransactionQueryInterface
		TransactionQuery    query.TransactionQueryInterface
		TypeActionSwitcher  TypeActionSwitcher
	}
	type args struct {
		blockTimestamp int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "wantSuccess",
			fields: fields{
				ID:            1234567890,
				Fee:           1,
				SenderAddress: "ABC",
				Height:        1,
				Body: &model.ApprovalEscrowTransactionBody{
					Approval:      1,
					TransactionID: 1234567890,
				},
				EscrowQuery:         &mockEscrowQueryApplyConfirmedOK{},
				QueryExecutor:       &mockEscrowQueryExecutorApplyConfirmedOK{},
				TransactionQuery:    &mockTransactionQueryApplyConfirmedOK{},
				AccountBalanceQuery: query.NewAccountBalanceQuery(),
				AccountLedgerQuery:  query.NewAccountLedgerQuery(),
				TypeActionSwitcher: &TypeSwitcher{
					Executor: &mockEscrowQueryExecutorApplyConfirmedOK{},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tx := &ApprovalEscrowTransaction{
				ID:                  tt.fields.ID,
				Fee:                 tt.fields.Fee,
				SenderAddress:       tt.fields.SenderAddress,
				Height:              tt.fields.Height,
				Body:                tt.fields.Body,
				Escrow:              tt.fields.Escrow,
				AccountBalanceQuery: tt.fields.AccountBalanceQuery,
				QueryExecutor:       tt.fields.QueryExecutor,
				AccountLedgerQuery:  tt.fields.AccountLedgerQuery,
				EscrowQuery:         tt.fields.EscrowQuery,
				TransactionQuery:    tt.fields.TransactionQuery,
				TypeActionSwitcher:  tt.fields.TypeActionSwitcher,
			}
			if err := tx.ApplyConfirmed(tt.args.blockTimestamp); (err != nil) != tt.wantErr {
				t.Errorf("ApplyConfirmed() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}