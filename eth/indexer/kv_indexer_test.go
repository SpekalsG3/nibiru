package indexer_test

import (
	"math/big"
	"testing"

	tmlog "cosmossdk.io/log"
	"cosmossdk.io/simapp/params"
	abci "github.com/cometbft/cometbft/abci/types"
	tmtypes "github.com/cometbft/cometbft/types"
	dbm "github.com/cosmos/cosmos-db"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ethereum/go-ethereum/common"
	gethcore "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/require"

	"github.com/NibiruChain/nibiru/v2/app"
	"github.com/NibiruChain/nibiru/v2/eth"
	"github.com/NibiruChain/nibiru/v2/eth/crypto/ethsecp256k1"
	evmenc "github.com/NibiruChain/nibiru/v2/eth/encoding"
	"github.com/NibiruChain/nibiru/v2/eth/indexer"
	"github.com/NibiruChain/nibiru/v2/x/evm"
	evmtest "github.com/NibiruChain/nibiru/v2/x/evm/evmtest"
)

func TestKVIndexer(t *testing.T) {
	priv, err := ethsecp256k1.GenerateKey()
	require.NoError(t, err)
	from := common.BytesToAddress(priv.PubKey().Address().Bytes())
	signer := evmtest.NewSigner(priv)
	ethSigner := gethcore.LatestSignerForChainID(nil)

	to := common.BigToAddress(big.NewInt(1))
	ethTxParams := evm.EvmTxArgs{
		Nonce:    0,
		To:       &to,
		Amount:   big.NewInt(1000),
		GasLimit: 21000,
	}
	tx := evm.NewTx(&ethTxParams)
	tx.From = from.Hex()
	require.NoError(t, tx.Sign(ethSigner, signer))
	txHash := tx.AsTransaction().Hash()

	// encCfg := MakeEncodingConfig()
	encCfg := app.MakeEncodingConfig()
	eth.RegisterInterfaces(encCfg.InterfaceRegistry)
	evm.RegisterInterfaces(encCfg.InterfaceRegistry)
	clientCtx := client.Context{}.
		WithTxConfig(encCfg.TxConfig).
		WithCodec(encCfg.Codec)

	// build cosmos-sdk wrapper tx
	tmTx, err := tx.BuildTx(clientCtx.TxConfig.NewTxBuilder(), eth.EthBaseDenom)
	require.NoError(t, err)
	txBz, err := clientCtx.TxConfig.TxEncoder()(tmTx)
	require.NoError(t, err)

	// build an invalid wrapper tx
	builder := clientCtx.TxConfig.NewTxBuilder()
	require.NoError(t, builder.SetMsgs(tx))
	tmTx2 := builder.GetTx()
	txBz2, err := clientCtx.TxConfig.TxEncoder()(tmTx2)
	require.NoError(t, err)

	testCases := []struct {
		name        string
		block       *tmtypes.Block
		blockResult []*abci.ExecTxResult
		expSuccess  bool
	}{
		{
			"success, format 1",
			&tmtypes.Block{Header: tmtypes.Header{Height: 1}, Data: tmtypes.Data{Txs: []tmtypes.Tx{txBz}}},
			[]*abci.ExecTxResult{
				{
					Code: 0,
					Events: []abci.Event{
						{Type: evm.EventTypeEthereumTx, Attributes: []abci.EventAttribute{
							{Key: "ethereumTxHash", Value: txHash.Hex()},
							{Key: "txIndex", Value: "0"},
							{Key: "amount", Value: "1000"},
							{Key: "txGasUsed", Value: "21000"},
							{Key: "txHash", Value: ""},
							{Key: "recipient", Value: "0x775b87ef5D82ca211811C1a02CE0fE0CA3a455d7"},
						}},
					},
				},
			},
			true,
		},
		{
			"success, format 2",
			&tmtypes.Block{Header: tmtypes.Header{Height: 1}, Data: tmtypes.Data{Txs: []tmtypes.Tx{txBz}}},
			[]*abci.ExecTxResult{
				{
					Code: 0,
					Events: []abci.Event{
						{Type: evm.EventTypeEthereumTx, Attributes: []abci.EventAttribute{
							{Key: "ethereumTxHash", Value: txHash.Hex()},
							{Key: "txIndex", Value: "0"},
						}},
						{Type: evm.EventTypeEthereumTx, Attributes: []abci.EventAttribute{
							{Key: "amount", Value: "1000"},
							{Key: "txGasUsed", Value: "21000"},
							{Key: "txHash", Value: "14A84ED06282645EFBF080E0B7ED80D8D8D6A36337668A12B5F229F81CDD3F57"},
							{Key: "recipient", Value: "0x775b87ef5D82ca211811C1a02CE0fE0CA3a455d7"},
						}},
					},
				},
			},
			true,
		},
		{
			"success, exceed block gas limit",
			&tmtypes.Block{Header: tmtypes.Header{Height: 1}, Data: tmtypes.Data{Txs: []tmtypes.Tx{txBz}}},
			[]*abci.ExecTxResult{
				{
					Code:   11,
					Log:    "out of gas in location: block gas meter; gasWanted: 21000",
					Events: []abci.Event{},
				},
			},
			true,
		},
		{
			"fail, failed eth tx",
			&tmtypes.Block{Header: tmtypes.Header{Height: 1}, Data: tmtypes.Data{Txs: []tmtypes.Tx{txBz}}},
			[]*abci.ExecTxResult{
				{
					Code:   15,
					Log:    "nonce mismatch",
					Events: []abci.Event{},
				},
			},
			false,
		},
		{
			"fail, invalid events",
			&tmtypes.Block{Header: tmtypes.Header{Height: 1}, Data: tmtypes.Data{Txs: []tmtypes.Tx{txBz}}},
			[]*abci.ExecTxResult{
				{
					Code:   0,
					Events: []abci.Event{},
				},
			},
			false,
		},
		{
			"fail, not eth tx",
			&tmtypes.Block{Header: tmtypes.Header{Height: 1}, Data: tmtypes.Data{Txs: []tmtypes.Tx{txBz2}}},
			[]*abci.ExecTxResult{
				{
					Code:   0,
					Events: []abci.Event{},
				},
			},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			db := dbm.NewMemDB()
			idxer := indexer.NewKVIndexer(db, tmlog.NewNopLogger(), clientCtx)

			err = idxer.IndexBlock(tc.block, tc.blockResult)
			require.NoError(t, err)
			if !tc.expSuccess {
				first, err := idxer.FirstIndexedBlock()
				require.NoError(t, err)
				require.Equal(t, int64(-1), first)

				last, err := idxer.LastIndexedBlock()
				require.NoError(t, err)
				require.Equal(t, int64(-1), last)
			} else {
				first, err := idxer.FirstIndexedBlock()
				require.NoError(t, err)
				require.Equal(t, tc.block.Header.Height, first)

				last, err := idxer.LastIndexedBlock()
				require.NoError(t, err)
				require.Equal(t, tc.block.Header.Height, last)

				res1, err := idxer.GetByTxHash(txHash)
				require.NoError(t, err)
				require.NotNil(t, res1)
				res2, err := idxer.GetByBlockAndIndex(1, 0)
				require.NoError(t, err)
				require.Equal(t, res1, res2)
			}
		})
	}
}

// MakeEncodingConfig creates the EncodingConfig
func MakeEncodingConfig() params.EncodingConfig {
	return evmenc.MakeConfig(app.ModuleBasics)
}
