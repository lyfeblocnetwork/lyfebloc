package v0_2_0_test

import (
	storetypes "cosmossdk.io/store/types"
	blobtypes "github.com/lyfeblocnetwork/lyfebloc/x/blob/types"
	bstypes "github.com/lyfeblocnetwork/lyfebloc/x/blobstream/types"
	datypes "github.com/lyfeblocnetwork/lyfebloc/x/da/types"

	"github.com/lyfeblocnetwork/lyfebloc/app/upgrades"
)

const UpgradeName string = "v0_2_0_test"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateUpgradeHandler,
	StoreUpgrades: storetypes.StoreUpgrades{
		Added:   []string{datypes.ModuleName},
		Deleted: []string{blobtypes.ModuleName, bstypes.ModuleName},
	},
}
