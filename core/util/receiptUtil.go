package util

import (
	"database/sql"

	"github.com/zoobc/zoobc-core/common/constant"
	"github.com/zoobc/zoobc-core/common/crypto"
	"github.com/zoobc/zoobc-core/common/kvdb"
	"github.com/zoobc/zoobc-core/common/model"
	"github.com/zoobc/zoobc-core/common/query"
	"github.com/zoobc/zoobc-core/common/util"
)

func GenerateBatchReceiptWithReminder(
	receivedDatumHash []byte,
	lastBlock *model.Block,
	senderPublicKey []byte,
	nodeSecretPhrase, receiptKey string,
	datumType uint32,
	signature crypto.SignatureInterface,
	queryExecutor query.ExecutorInterface,
	kvExecutor kvdb.KVExecutorInterface,
) (*model.BatchReceipt, error) {
	var (
		rmrLinked     []byte
		batchReceipt  *model.BatchReceipt
		err           error
		merkleQuery   = query.NewMerkleTreeQuery()
		nodePublicKey = util.GetPublicKeyFromSeed(nodeSecretPhrase)
		lastRmrQ      = merkleQuery.GetLastMerkleRoot()
		row           = queryExecutor.ExecuteSelectRow(lastRmrQ)
	)

	rmrLinked, err = merkleQuery.ScanRoot(row)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}
	// generate receipt
	batchReceipt, err = util.GenerateBatchReceipt(
		lastBlock,
		senderPublicKey,
		nodePublicKey,
		receivedDatumHash,
		rmrLinked,
		datumType,
	)
	if err != nil {
		return nil, err
	}
	batchReceipt.RecipientSignature = signature.SignByNode(
		util.GetUnsignedBatchReceiptBytes(batchReceipt),
		nodeSecretPhrase,
	)
	// store the generated batch receipt hash for reminder
	err = kvExecutor.Insert(receiptKey, receivedDatumHash, constant.KVdbExpiryReceiptReminder)
	if err != nil {
		return nil, err
	}
	return batchReceipt, nil
}