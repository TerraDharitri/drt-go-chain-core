package sovereign

import (
	"math/big"

	"github.com/TerraDharitri/drt-go-chain-core/core"
)

// Operation holds the data needed for an outgoing operation
type Operation struct {
	Address []byte
	Tokens  []DcdtToken
	Data    *EventData
}

// DcdtToken holds the token data
type DcdtToken struct {
	Identifier []byte
	Nonce      uint64
	Data       DcdtTokenData
}

// DcdtTokenData holds the properties for a token
type DcdtTokenData struct {
	TokenType  core.DCDTType
	Amount     *big.Int
	Frozen     bool
	Hash       []byte
	Name       []byte
	Attributes []byte
	Creator    []byte
	Royalties  *big.Int
	Uris       [][]byte
}
