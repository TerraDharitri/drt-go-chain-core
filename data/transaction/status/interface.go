package status

import (
	"github.com/TerraDharitri/drt-go-chain-core/data/block"
	"github.com/TerraDharitri/drt-go-chain-core/data/transaction"
)

// StatusComputerHandler computes a transaction status
type StatusComputerHandler interface {
	ComputeStatusWhenInStorageKnowingMiniblock(miniblockType block.Type, tx *transaction.ApiTransactionResult) (transaction.TxStatus, error)
	ComputeStatusWhenInStorageNotKnowingMiniblock(destinationShard uint32, tx *transaction.ApiTransactionResult) (transaction.TxStatus, error)
	SetStatusIfIsRewardReverted(
		tx *transaction.ApiTransactionResult,
		miniblockType block.Type,
		headerNonce uint64,
		headerHash []byte,
	) (bool, error)
}
