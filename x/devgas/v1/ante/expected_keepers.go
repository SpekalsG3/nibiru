package ante

// Interfaces needed for the Nibiru Chain ante handler

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"

	devgastypes "github.com/NibiruChain/nibiru/v2/x/devgas/v1/types"
)

type BankKeeper interface {
	SendCoinsFromAccountToModule(
		ctx context.Context, senderAddr sdk.AccAddress,
		recipientModule string, amt sdk.Coins,
	) error
	SendCoinsFromModuleToAccount(
		ctx context.Context, senderModule string,
		recipientAddr sdk.AccAddress, amt sdk.Coins,
	) error
}

type IDevGasKeeper interface {
	GetParams(ctx context.Context) devgastypes.ModuleParams
	GetFeeShare(ctx context.Context, contract sdk.Address) (devgastypes.FeeShare, bool)
}
