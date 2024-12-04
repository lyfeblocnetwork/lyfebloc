package v0_2_6_test

import (
	storetypes "cosmossdk.io/store/types"

	"github.com/lyfeblocnetwork/lyfebloc/app/upgrades"
)

const UpgradeName string = "v0_2_6_test"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added:   []string{},
		Deleted: []string{},
	},
}
