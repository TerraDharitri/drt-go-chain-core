package api

import "github.com/TerraDharitri/drt-go-chain-core/core"

// AccountQueryOptions holds options for account queries
type AccountQueryOptions struct {
	OnFinalBlock   bool
	OnStartOfEpoch core.OptionalUint32
	BlockNonce     core.OptionalUint64
	BlockHash      []byte
	BlockRootHash  []byte
	HintEpoch      core.OptionalUint32
	WithKeys       bool
}

// BlockQueryOptions holds options for block queries
type BlockQueryOptions struct {
	WithTransactions bool
	WithLogs         bool
	ForHyperblock    bool
}
