// Copyright (c) 2023-2024 Nibi, Inc.
package ante

import (
	"cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/NibiruChain/nibiru/v2/x/evm"
)

// AnteDecoratorPreventEtheruemTxMsgs prevents invalid msg types from being executed
type AnteDecoratorPreventEtheruemTxMsgs struct{}

// AnteHandle rejects messages that requires ethereum-specific authentication.
// For example `MsgEthereumTx` requires fee to be deducted in the antehandler in
// order to perform the refund.
func (rmd AnteDecoratorPreventEtheruemTxMsgs) AnteHandle(
	ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler,
) (newCtx sdk.Context, err error) {
	for _, msg := range tx.GetMsgs() {
		if _, ok := msg.(*evm.MsgEthereumTx); ok {
			return ctx, errors.Wrapf(
				errortypes.ErrInvalidType,
				"MsgEthereumTx needs to be contained within a tx with 'ExtensionOptionsEthereumTx' option",
			)
		}
	}
	return next(ctx, tx, simulate)
}
