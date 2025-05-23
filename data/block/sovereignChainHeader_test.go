package block

import (
	"math/big"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/TerraDharitri/drt-go-chain-core/core"
	"github.com/TerraDharitri/drt-go-chain-core/data"
)

func TestSovereignChainHeader_GetEpochStartHandler(t *testing.T) {
	t.Parallel()

	economics := Economics{
		NodePrice: big.NewInt(100),
	}
	epochStartData := EpochStartCrossChainData{
		Round: 4,
	}
	epochStartSov := EpochStartSovereign{
		Economics:                     economics,
		LastFinalizedCrossChainHeader: epochStartData,
	}
	sovHdr := &SovereignChainHeader{
		EpochStart: epochStartSov,
	}

	require.Equal(t, &epochStartSov, sovHdr.GetEpochStartHandler())
	require.Equal(t, &economics, sovHdr.GetEpochStartHandler().GetEconomicsHandler())
	require.Empty(t, sovHdr.GetEpochStartHandler().GetLastFinalizedHeaderHandlers())

	epochStartData.ShardID = 8
	err := sovHdr.GetEpochStartHandler().SetLastFinalizedHeaders([]data.EpochStartShardDataHandler{&epochStartData})
	require.Nil(t, err)
	require.Empty(t, sovHdr.GetEpochStartHandler().GetLastFinalizedHeaderHandlers())

	epochStartData.ShardID = core.MainChainShardId
	epochStartSov.LastFinalizedCrossChainHeader = epochStartData
	err = sovHdr.GetEpochStartHandler().SetLastFinalizedHeaders([]data.EpochStartShardDataHandler{&epochStartData})
	require.Nil(t, err)

	require.Equal(t, []data.EpochStartShardDataHandler{&epochStartData}, sovHdr.GetEpochStartHandler().GetLastFinalizedHeaderHandlers())
	require.Equal(t, &epochStartSov, sovHdr.GetEpochStartHandler())
}

func TestSovereignChainHeader_GetEconomicsHandler(t *testing.T) {
	t.Parallel()

	economics := Economics{
		NodePrice: big.NewInt(100),
	}
	sovHdr := &SovereignChainHeader{
		EpochStart: EpochStartSovereign{
			Economics: economics,
		},
	}
	require.Equal(t, &economics, sovHdr.GetEpochStartHandler().GetEconomicsHandler())

	economics2 := Economics{
		NodePrice: big.NewInt(200),
	}
	err := sovHdr.GetEpochStartHandler().SetEconomics(&economics2)
	require.Nil(t, err)
	require.Equal(t, &economics2, sovHdr.GetEpochStartHandler().GetEconomicsHandler())
}

func TestSovereignChainHeader_ShallowClone(t *testing.T) {
	t.Parallel()

	sovHdr := &SovereignChainHeader{
		DevFeesInEpoch: big.NewInt(100),
		OutGoingMiniBlockHeaders: []*OutGoingMiniBlockHeader{
			{
				Hash: []byte("h1"),
			},
			{
				Hash: []byte("h2"),
			},
		},
	}

	sovHdrClone := sovHdr.ShallowClone()
	require.Equal(t, sovHdr, sovHdrClone)
	require.False(t, sovHdr == sovHdrClone)

	sovHdr.Header = &Header{Nonce: 4}
	sovHdrClone = sovHdr.ShallowClone()
	require.Equal(t, sovHdr, sovHdrClone)
	require.False(t, sovHdr == sovHdrClone)
	require.False(t, sovHdr.Header == sovHdrClone.(*SovereignChainHeader).Header)
}

func TestSovereignChainHeader_ShallowCloneWithOutGoingMBs(t *testing.T) {
	t.Parallel()

	sovHdr := &SovereignChainHeader{
		DevFeesInEpoch: big.NewInt(100),
		OutGoingMiniBlockHeaders: []*OutGoingMiniBlockHeader{
			{
				Hash: []byte("h1"),
			},
			{
				Hash:                              []byte("h2"),
				LeaderSignatureOutGoingOperations: []byte("leaderSig"),
			},
		},
	}

	sovHdrHandlerClone := sovHdr.ShallowClone()

	sovHdr.OutGoingMiniBlockHeaders[0].Hash = []byte("h3")
	sovHdr.OutGoingMiniBlockHeaders[1].LeaderSignatureOutGoingOperations = nil

	require.Equal(t, &SovereignChainHeader{
		DevFeesInEpoch: big.NewInt(100),
		OutGoingMiniBlockHeaders: []*OutGoingMiniBlockHeader{
			{
				Hash: []byte("h1"),
			},
			{
				Hash:                              []byte("h2"),
				LeaderSignatureOutGoingOperations: []byte("leaderSig"),
			},
		},
	}, sovHdrHandlerClone.(*SovereignChainHeader))
}

func TestSovereignChainHeader_GetOutGoingMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	sovHdr := &SovereignChainHeader{
		OutGoingMiniBlockHeaders: []*OutGoingMiniBlockHeader{
			{
				Hash: []byte("h1"),
			},
			{
				Hash: []byte("h2"),
			},
		},
	}

	require.Equal(t,
		[]data.OutGoingMiniBlockHeaderHandler{sovHdr.OutGoingMiniBlockHeaders[0], sovHdr.OutGoingMiniBlockHeaders[1]},
		sovHdr.GetOutGoingMiniBlockHeaderHandlers(),
	)
}

func TestSovereignChainHeader_GetOutGoingMiniBlockHeaderHandler(t *testing.T) {
	t.Parallel()

	sovHdr := &SovereignChainHeader{
		OutGoingMiniBlockHeaders: []*OutGoingMiniBlockHeader{
			{
				Type: OutGoingMbTx,
				Hash: []byte("h1"),
			},
			{
				Type: OutGoingMbChangeValidatorSet,
				Hash: []byte("h2"),
			},
		},
	}

	mb := sovHdr.GetOutGoingMiniBlockHeaderHandler(int32(OutGoingMbTx))
	require.Equal(t, sovHdr.OutGoingMiniBlockHeaders[0], mb)
	mb = sovHdr.GetOutGoingMiniBlockHeaderHandler(int32(OutGoingMbChangeValidatorSet))
	require.Equal(t, sovHdr.OutGoingMiniBlockHeaders[1], mb)
	require.Nil(t, sovHdr.GetOutGoingMiniBlockHeaderHandler(-1))
}

func TestSovereignChainHeader_SetOutGoingMiniBlockHeaderHandlers(t *testing.T) {
	t.Parallel()

	sovHdr := &SovereignChainHeader{}
	mbHeader1 := &OutGoingMiniBlockHeader{Type: OutGoingMbTx, Hash: []byte("h1")}
	mbHeader2 := &OutGoingMiniBlockHeader{Type: OutGoingMbChangeValidatorSet, Hash: []byte("h2")}

	err := sovHdr.SetOutGoingMiniBlockHeaderHandlers([]data.OutGoingMiniBlockHeaderHandler{mbHeader1, mbHeader2})
	require.Nil(t, err)
	require.Equal(t,
		[]data.OutGoingMiniBlockHeaderHandler{mbHeader1, mbHeader2},
		sovHdr.GetOutGoingMiniBlockHeaderHandlers(),
	)
}

func TestSovereignChainHeader_SetOutGoingMiniBlockHeaderHandler(t *testing.T) {
	t.Parallel()

	sovHdr := &SovereignChainHeader{}
	mbHeader1 := &OutGoingMiniBlockHeader{Type: OutGoingMbTx, Hash: []byte("h1")}
	mbHeader2 := &OutGoingMiniBlockHeader{Type: OutGoingMbTx, Hash: []byte("h2")}
	mbHeader3 := &OutGoingMiniBlockHeader{Type: OutGoingMbChangeValidatorSet, Hash: []byte("h3")}

	err := sovHdr.SetOutGoingMiniBlockHeaderHandler(mbHeader1)
	require.Nil(t, err)
	require.Equal(t,
		[]data.OutGoingMiniBlockHeaderHandler{mbHeader1},
		sovHdr.GetOutGoingMiniBlockHeaderHandlers(),
	)

	err = sovHdr.SetOutGoingMiniBlockHeaderHandler(mbHeader2)
	require.Nil(t, err)
	require.Equal(t,
		[]data.OutGoingMiniBlockHeaderHandler{mbHeader2},
		sovHdr.GetOutGoingMiniBlockHeaderHandlers(),
	)

	err = sovHdr.SetOutGoingMiniBlockHeaderHandler(mbHeader3)
	require.Nil(t, err)
	require.Equal(t,
		[]data.OutGoingMiniBlockHeaderHandler{mbHeader2, mbHeader3},
		sovHdr.GetOutGoingMiniBlockHeaderHandlers(),
	)
}
