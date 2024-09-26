package sudo

import (
	"context"

	"github.com/NibiruChain/nibiru/x/sudo/keeper"
	"github.com/NibiruChain/nibiru/x/sudo/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the module's state from a provided genesis state JSON.
func InitGenesis(ctx context.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := genState.Validate(); err != nil {
		panic(err)
	}
	k.Sudoers.Set(sdk.UnwrapSDKContext(ctx), genState.Sudoers)
}

// ExportGenesis returns the module's exported genesis state.
// This fn assumes InitGenesis has already been called.
func ExportGenesis(ctx context.Context, k keeper.Keeper) *types.GenesisState {
	pbSudoers, err := k.Sudoers.Get(sdk.UnwrapSDKContext(ctx))
	if err != nil {
		panic(err)
	}

	return &types.GenesisState{
		Sudoers: pbSudoers,
	}
}

// DefaultGenesis: A blank genesis state. The DefaultGenesis is invalid because
// it does not specify a "Sudoers.Root".
func DefaultGenesis() *types.GenesisState {
	return &types.GenesisState{
		Sudoers: types.Sudoers{
			Root:      "",
			Contracts: []string{},
		},
	}
}
