package upgrades

import (
	store "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"cosmossdk.io/x/upgrade/types"
)

type Upgrade struct {
	UpgradeName string

	CreateUpgradeHandler func(*module.Manager, module.Configurator) types.UpgradeHandler

	StoreUpgrades store.StoreUpgrades
}
