package bank_test

import (
	"testing"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lyfeblocnetwork/lyfebloc/precompiles/bank"
	"github.com/lyfeblocnetwork/lyfebloc/testutil/integration/lyfebloc/factory"
	"github.com/lyfeblocnetwork/lyfebloc/testutil/integration/lyfebloc/grpc"
	testkeyring "github.com/lyfeblocnetwork/lyfebloc/testutil/integration/lyfebloc/keyring"
	"github.com/lyfeblocnetwork/lyfebloc/testutil/integration/lyfebloc/network"
	integrationutils "github.com/lyfeblocnetwork/lyfebloc/testutil/integration/lyfebloc/utils"
	inflationtypes "github.com/lyfeblocnetwork/lyfebloc/x/inflation/v1/types"
	"github.com/stretchr/testify/suite"
)

var s *PrecompileTestSuite

// PrecompileTestSuite is the implementation of the TestSuite interface for ERC20 precompile
// unit tests.
type PrecompileTestSuite struct {
	suite.Suite

	bondDenom, tokenDenom string
	lyfeblocAddr, xmplAddr   common.Address

	// tokenDenom is the specific token denomination used in testing the ERC20 precompile.
	// This denomination is used to instantiate the precompile.
	network     *network.UnitTestNetwork
	factory     factory.TxFactory
	grpcHandler grpc.Handler
	keyring     testkeyring.Keyring

	precompile *bank.Precompile
}

func TestPrecompileTestSuite(t *testing.T) {
	s = new(PrecompileTestSuite)
	suite.Run(t, s)
}

func (s *PrecompileTestSuite) SetupTest() sdk.Context {
	s.tokenDenom = xmplDenom

	keyring := testkeyring.New(2)
	genesis := integrationutils.CreateGenesisWithTokenPairs(keyring)
	unitNetwork := network.NewUnitTestNetwork(
		network.WithPreFundedAccounts(keyring.GetAllAccAddrs()...),
		network.WithCustomGenesis(genesis),
		network.WithOtherDenoms([]string{s.tokenDenom}),
	)
	grpcHandler := grpc.NewIntegrationHandler(unitNetwork)
	txFactory := factory.New(unitNetwork, grpcHandler)

	ctx := unitNetwork.GetContext()
	sk := unitNetwork.App.StakingKeeper
	bondDenom, err := sk.BondDenom(ctx)
	s.Require().NoError(err, "failed to get bond denom")
	s.Require().NotEmpty(bondDenom, "bond denom cannot be empty")

	s.bondDenom = bondDenom
	s.factory = txFactory
	s.grpcHandler = grpcHandler
	s.keyring = keyring
	s.network = unitNetwork

	tokenPairID := s.network.App.Erc20Keeper.GetTokenPairID(s.network.GetContext(), s.bondDenom)
	tokenPair, found := s.network.App.Erc20Keeper.GetTokenPair(s.network.GetContext(), tokenPairID)
	s.Require().True(found)
	s.lyfeblocAddr = common.HexToAddress(tokenPair.Erc20Address)

	s.lyfeblocAddr = tokenPair.GetERC20Contract()

	// Mint and register a second coin for testing purposes
	err = s.network.App.BankKeeper.MintCoins(s.network.GetContext(), inflationtypes.ModuleName, sdk.Coins{{Denom: "xmpl", Amount: math.NewInt(1e18)}})
	s.Require().NoError(err)

	tokenPairID = s.network.App.Erc20Keeper.GetTokenPairID(s.network.GetContext(), s.tokenDenom)
	tokenPair, found = s.network.App.Erc20Keeper.GetTokenPair(s.network.GetContext(), tokenPairID)
	s.Require().True(found)
	s.xmplAddr = common.HexToAddress(tokenPair.Erc20Address)

	s.xmplAddr = tokenPair.GetERC20Contract()

	s.precompile = s.setupBankPrecompile()
	return ctx
}