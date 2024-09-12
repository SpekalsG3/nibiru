// Copyright (c) 2023-2024 Nibi, Inc.
package evm

import (
	"context"

	"cosmossdk.io/core/address"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

// AccountKeeper defines the expected account keeper interface
type AccountKeeper interface {
	NewAccountWithAddress(ctx context.Context, addr sdk.AccAddress) sdk.AccountI

	// GetModuleAccount gets the module account from the auth account store, if the
	// account does not exist in the AccountKeeper, then it is created. This
	// differs from the "GetModuleAddress" function, which performs a pure
	// computation.
	GetModuleAccount(ctx context.Context, moduleName string) sdk.ModuleAccountI

	// GetModuleAddress returns an address based on the module name, however it
	// does not modify state at all. To create initialize the module account,
	// instead use "GetModuleAccount".
	GetModuleAddress(moduleName string) sdk.AccAddress
	GetAllAccounts(ctx context.Context) (accounts []sdk.AccountI)
	IterateAccounts(ctx context.Context, cb func(account sdk.AccountI) bool)
	GetSequence(context.Context, sdk.AccAddress) (uint64, error)
	GetAccount(ctx context.Context, addr sdk.AccAddress) sdk.AccountI
	SetAccount(ctx context.Context, account sdk.AccountI)
	RemoveAccount(ctx context.Context, account sdk.AccountI)
	GetParams(ctx context.Context) (params authtypes.Params)
	SetModuleAccount(ctx context.Context, macc sdk.ModuleAccountI)
	AddressCodec() address.Codec
}

// BankKeeper defines the expected interface needed to retrieve account balances.
type BankKeeper interface {
	authtypes.BankKeeper
	GetBalance(ctx context.Context, addr sdk.AccAddress, denom string) sdk.Coin
	SendCoinsFromModuleToAccount(ctx context.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	MintCoins(ctx context.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx context.Context, moduleName string, amt sdk.Coins) error

	GetDenomMetaData(ctx context.Context, denom string) (metadata bank.Metadata, isFound bool)
	SetDenomMetaData(ctx context.Context, denomMetaData bank.Metadata)
}

// StakingKeeper returns the historical headers kept in store.
type StakingKeeper interface {
	GetHistoricalInfo(ctx context.Context, height int64) (stakingtypes.HistoricalInfo, error)
	GetValidatorByConsAddr(ctx context.Context, consAddr sdk.ConsAddress) (validator stakingtypes.Validator, found error)
}
