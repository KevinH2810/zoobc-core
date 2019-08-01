package util

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/query"
	"golang.org/x/crypto/sha3"
)

// GetTransactionBytes translate transaction model to its byte representation
// provide sign = true to translate transaction with its signature, sign = false
// for without signature (used for verify signature)
func GetTransactionBytes(transaction *model.Transaction, sign bool) ([]byte, error) {
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write(ConvertUint32ToBytes(transaction.TransactionType)[:2])
	buffer.Write(ConvertUint32ToBytes(transaction.Version)[:1])
	buffer.Write(ConvertUint64ToBytes(uint64(transaction.Timestamp)))
	buffer.Write(ConvertUint32ToBytes(transaction.SenderAccountType)[:2])
	buffer.Write([]byte(transaction.SenderAccountAddress))
	buffer.Write(ConvertUint32ToBytes(transaction.RecipientAccountType)[:2])
	if transaction.RecipientAccountAddress == "" {
		buffer.Write(make([]byte, 44)) // if no recipient pad with 44 (zoobc address length)
	} else {
		buffer.Write([]byte(transaction.RecipientAccountAddress))
	}
	buffer.Write(ConvertUint64ToBytes(uint64(transaction.Fee)))
	// transaction body length
	buffer.Write(ConvertUint32ToBytes(transaction.TransactionBodyLength))
	buffer.Write(transaction.TransactionBodyBytes)
	if sign {
		if transaction.Signature == nil {
			return nil, errors.New("TransactionSignatureNotExist")
		}
		buffer.Write(transaction.Signature)
	}
	return buffer.Bytes(), nil
}

func readTransactionBytes(buf *bytes.Buffer, nBytes int) ([]byte, error) {
	nextBytes := buf.Next(nBytes)
	if len(nextBytes) < nBytes {
		return nil, errors.New("EndOfBufferReached")
	}
	return nextBytes, nil
}

// ParseTransactionBytes build transaction from transaction bytes
func ParseTransactionBytes(transactionBytes []byte, sign bool) (*model.Transaction, error) {
	buffer := bytes.NewBuffer(transactionBytes)

	transactionTypeBytes, err := readTransactionBytes(buffer, 2)
	if err != nil {
		return nil, err
	}
	transactionType := ConvertBytesToUint32([]byte{transactionTypeBytes[0], transactionTypeBytes[1], 0, 0})
	transactionVersionByte, err := readTransactionBytes(buffer, 1)
	if err != nil {
		return nil, err
	}
	transactionVersion := ConvertBytesToUint32([]byte{transactionVersionByte[0], 0, 0, 0})
	timestampBytes, err := readTransactionBytes(buffer, 8)
	if err != nil {
		return nil, err
	}
	timestamp := ConvertBytesToUint64(timestampBytes)
	senderAccountType, err := readTransactionBytes(buffer, 2)
	if err != nil {
		return nil, err
	}
	senderAccountAddress := ReadAccountAddress(ConvertBytesToUint32([]byte{
		senderAccountType[0], senderAccountType[1], 0, 0,
	}), buffer)
	recipientAccountType, err := readTransactionBytes(buffer, 2)
	if err != nil {
		return nil, err
	}
	recipientAccountAddress := ReadAccountAddress(ConvertBytesToUint32([]byte{
		recipientAccountType[0], recipientAccountType[1], 0, 0,
	}), buffer)
	feeBytes, err := readTransactionBytes(buffer, 8)
	if err != nil {
		return nil, err
	}
	fee := ConvertBytesToUint64(feeBytes)
	transactionBodyLengthBytes, err := readTransactionBytes(buffer, 4)
	if err != nil {
		return nil, err
	}
	transactionBodyLength := ConvertBytesToUint32(transactionBodyLengthBytes)
	transactionBodyBytes, err := readTransactionBytes(buffer, int(transactionBodyLength))
	if err != nil {
		return nil, err
	}
	var sig []byte
	if sign {
		var err error
		//TODO: implement below logic to allow multiple signature algorithm to work
		// first 2 bytes of signature are the signature length
		// signatureLengthBytes, err := readTransactionBytes(buffer, 2)
		// if err != nil {
		// 	return nil, err
		// }
		// signatureLength := int(ConvertBytesToUint32(signatureLengthBytes))
		sig, err = readTransactionBytes(buffer, 64)
		if err != nil {
			return nil, errors.New("TrasnsactionSignatureNotExist")
		}
	}
	// compute and return tx hash and ID too
	transactionHash := sha3.Sum256(transactionBytes)
	fmt.Printf("%v", transactionHash)
	txID, _ := GetTransactionID(transactionHash[:])
	return &model.Transaction{
		ID:              txID,
		TransactionType: transactionType,
		Version:         transactionVersion,
		Timestamp:       int64(timestamp),
		SenderAccountType: ConvertBytesToUint32([]byte{
			senderAccountType[0], senderAccountType[1], 0, 0}),
		SenderAccountAddress: string(senderAccountAddress),
		RecipientAccountType: ConvertBytesToUint32([]byte{
			recipientAccountType[0], recipientAccountType[1], 0, 0,
		}),
		RecipientAccountAddress: string(recipientAccountAddress),
		Fee:                     int64(fee),
		TransactionBodyLength:   transactionBodyLength,
		TransactionBodyBytes:    transactionBodyBytes,
		TransactionHash:         transactionHash[:],
		Signature:               sig,
	}, nil
}

// GetTransactionID calculate and returns a transaction ID given a transaction model
func GetTransactionID(transactionHash []byte) (int64, error) {
	if len(transactionHash) == 0 {
		return -1, errors.New("InvalidTransactionHash")
	}
	ID := int64(ConvertBytesToUint64(transactionHash))
	return ID, nil
}

// ReadAccountAddress support different way to read the sender or recipient address depending on
// their types.
func ReadAccountAddress(accountType uint32, buf *bytes.Buffer) []byte {
	switch accountType {
	case 0:
		return buf.Next(44) // zoobc account address length
	default:
		return buf.Next(44) // default to zoobc account address
	}
}

func ValidateTransaction(
	tx *model.Transaction,
	queryExecutor query.ExecutorInterface,
	accountBalanceQuery query.AccountBalanceQueryInterface,
	verifySignature bool,
) error {
	// don't validate genesis transactions
	if tx.Height == 0 {
		return nil
	}
	if tx.Fee <= 0 {
		return errors.New("TxFeeZero")
	}
	if tx.SenderAccountAddress == "" {
		return errors.New("TxSenderEmpty")
	}
	if err := ValidateAccountAddress(tx.SenderAccountType, tx.SenderAccountAddress); err != nil {
		return err
	}
	if tx.RecipientAccountAddress != "" {
		if err := ValidateAccountAddress(tx.RecipientAccountType, tx.RecipientAccountAddress); err != nil {
			return err
		}
	}

	// validate sender account
	senderAccountID := CreateAccountIDFromAddress(tx.SenderAccountType, tx.SenderAccountAddress)
	sqlQ, arg := accountBalanceQuery.GetAccountBalanceByAccountID(senderAccountID)
	rows, err := queryExecutor.ExecuteSelect(sqlQ, arg)
	if err != nil {
		return err
	}
	defer rows.Close()
	res := accountBalanceQuery.BuildModel([]*model.AccountBalance{}, rows)
	if len(res) == 0 {
		return errors.New("TxSenderNotFound")
	}
	senderAccountBalance := res[0]
	if senderAccountBalance.SpendableBalance < tx.Fee {
		return errors.New("TxAccountBalanceNotEnough")
	}

	// formally validate transaction body
	if len(tx.TransactionBodyBytes) != int(tx.TransactionBodyLength) {
		return errors.New("TxInvalidBodyFormat")
	}

	//FIXME: comemented out for now because gives circular dependency (both this and crypto packages import common/util)..
	// transactionBytes, err := GetTransactionBytes(tx, true)
	// if err != nil {
	// 	return err
	// }
	// if verifySignature {
	// 	if !crypto.NewSignature().VerifySignature(transactionBytes, tx.Signature, tx.SenderAccountType, tx.SenderAccountAddress) {
	// 		return errors.New("TxInvalidSignature")
	// 	}
	// }

	return nil
}