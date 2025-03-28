package transaction

import "github.com/TerraDharitri/drt-go-chain-core/data/block"

// StatusComputerHandler computes a transaction status
type StatusComputerHandler interface {
	ComputeStatusWhenInStorageKnowingMiniblock(miniblockType block.Type, tx *ApiTransactionResult) (TxStatus, error)
	ComputeStatusWhenInStorageNotKnowingMiniblock(destinationShard uint32, tx *ApiTransactionResult) (TxStatus, error)
	SetStatusIfIsRewardReverted(
		tx *ApiTransactionResult,
		miniblockType block.Type,
		headerNonce uint64,
		headerHash []byte,
	) (bool, error)
}
