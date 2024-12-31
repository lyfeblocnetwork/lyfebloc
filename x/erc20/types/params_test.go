package types_test

import (
	"slices"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/lyfeblocnetwork/lyfebloc/x/erc20/types"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ParamsTestSuite struct {
	suite.Suite
}

func TestParamsTestSuite(t *testing.T) {
	suite.Run(t, new(ParamsTestSuite))
}

func (suite *ParamsTestSuite) TestParamsValidate() {
	testCases := []struct {
		name        string
		malleate    func() types.Params
		expError    bool
		errContains string
	}{
		{
			"default",
			types.DefaultParams,
			false,
			"",
		},
		{
			"valid",
			func() types.Params { return types.NewParams(true, []string{}, []string{}) },
			false,
			"",
		},
		{
			"valid address - dynamic precompile",
			func() types.Params { return types.NewParams(true, []string{}, []string{types.WLYFEBLOCContractMainnet}) },
			false,
			"",
		},
		{
			"valid address - native precompile",
			func() types.Params { return types.NewParams(true, []string{types.WLYFEBLOCContractMainnet}, []string{}) },
			false,
			"",
		},
		{
			"sorted address",
			// order of creation shouldn't matter since it should be sorted when defining new param
			func() types.Params {
				return types.NewParams(true, []string{types.WLYFEBLOCContractTestnet, types.WLYFEBLOCContractMainnet}, []string{})
			},
			false,
			"",
		},
		{
			"unsorted address",
			// order of creation shouldn't matter since it should be sorted when defining new param
			func() types.Params {
				return types.NewParams(true, []string{types.WLYFEBLOCContractMainnet, types.WLYFEBLOCContractTestnet}, []string{})
			},
			false,
			"",
		},
		{
			"empty",
			func() types.Params { return types.Params{} },
			false,
			"",
		},
		{
			"invalid address - native precompile",
			func() types.Params {
				return types.NewParams(true, []string{"qq"}, []string{})
			},
			true,
			"invalid precompile",
		},
		{
			"invalid address - dynamic precompile",
			func() types.Params {
				return types.NewParams(true, []string{}, []string{"0xqq"})
			},
			true,
			"invalid precompile",
		},
		{
			"repeated address in different params",
			func() types.Params {
				return types.NewParams(true, []string{types.WLYFEBLOCContractMainnet}, []string{types.WLYFEBLOCContractMainnet})
			},
			true,
			"duplicate precompile",
		},
		{
			"repeated address - native precompiles",
			func() types.Params {
				return types.NewParams(true, []string{types.WLYFEBLOCContractMainnet, types.WLYFEBLOCContractMainnet}, []string{})
			},
			true,
			"duplicate precompile",
		},
		{
			"repeated address - dynamic precompiles",
			func() types.Params {
				return types.NewParams(true, []string{}, []string{types.WLYFEBLOCContractMainnet, types.WLYFEBLOCContractMainnet})
			},
			true,
			"duplicate precompile",
		},
		{
			"repeated address - one EIP-55 other not",
			func() types.Params {
				return types.NewParams(true, []string{}, []string{"0xcc491f589b45d4a3c679016195b3fb87d7848210", "0xcc491f589B45d4a3C679016195B3FB87D7848210"})
			},
			true,
			"duplicate precompile",
		},
		{
			"unsorted addresses",
			func() types.Params {
				params := types.DefaultParams()
				params.NativePrecompiles = []string{types.WLYFEBLOCContractTestnet, types.WLYFEBLOCContractMainnet}
				return params
			},
			true,
			"precompiles need to be sorted",
		},
	}

	for _, tc := range testCases {
		p := tc.malleate()
		err := p.Validate()

		if tc.expError {
			suite.Require().Error(err, tc.name)
			suite.Require().ErrorContains(err, tc.errContains)
		} else {
			suite.Require().NoError(err, tc.name)
		}
	}
}

func (suite *ParamsTestSuite) TestIsNativePrecompile() {
	testCases := []struct {
		name     string
		malleate func() types.Params
		addr     common.Address
		expRes   bool
	}{
		{
			"default",
			types.DefaultParams,
			common.HexToAddress(types.WLYFEBLOCContractMainnet),
			true,
		},
		{
			"not native precompile",
			func() types.Params { return types.NewParams(true, nil, nil) },
			common.HexToAddress(types.WLYFEBLOCContractMainnet),
			false,
		},
		{
			"EIP-55 address - is native precompile",
			func() types.Params {
				return types.NewParams(true, []string{"0xcc491f589B45d4a3C679016195B3FB87D7848210"}, nil)
			},
			common.HexToAddress(types.WLYFEBLOCContractTestnet),
			true,
		},
		{
			"NOT EIP-55 address - is native precompile",
			func() types.Params {
				return types.NewParams(true, []string{"0xcc491f589b45d4a3c679016195b3fb87d7848210"}, nil)
			},
			common.HexToAddress(types.WLYFEBLOCContractTestnet),
			true,
		},
	}

	for _, tc := range testCases {
		p := tc.malleate()
		suite.Require().Equal(tc.expRes, p.IsNativePrecompile(tc.addr), tc.name)
	}
}

func (suite *ParamsTestSuite) TestIsDynamicPrecompile() {
	testCases := []struct {
		name     string
		malleate func() types.Params
		addr     common.Address
		expRes   bool
	}{
		{
			"default - not dynamic precompile",
			types.DefaultParams,
			common.HexToAddress(types.WLYFEBLOCContractMainnet),
			false,
		},
		{
			"no dynamic precompiles",
			func() types.Params { return types.NewParams(true, nil, nil) },
			common.HexToAddress(types.WLYFEBLOCContractMainnet),
			false,
		},
		{
			"EIP-55 address - is dynamic precompile",
			func() types.Params {
				return types.NewParams(true, nil, []string{"0xcc491f589B45d4a3C679016195B3FB87D7848210"})
			},
			common.HexToAddress(types.WLYFEBLOCContractTestnet),
			true,
		},
		{
			"NOT EIP-55 address - is dynamic precompile",
			func() types.Params {
				return types.NewParams(true, nil, []string{"0xcc491f589b45d4a3c679016195b3fb87d7848210"})
			},
			common.HexToAddress(types.WLYFEBLOCContractTestnet),
			true,
		},
	}

	for _, tc := range testCases {
		p := tc.malleate()
		suite.Require().Equal(tc.expRes, p.IsDynamicPrecompile(tc.addr), tc.name)
	}
}

func (suite *ParamsTestSuite) TestParamsValidatePriv() {
	suite.Require().Error(types.ValidateBool(1))
	suite.Require().NoError(types.ValidateBool(true))
}

func TestValidatePrecompiles(t *testing.T) {
	testCases := []struct {
		name        string
		precompiles []string
		expError    bool
		errContains string
	}{
		{
			"invalid precompile address",
			[]string{"0xct491f589b45d4a3c679016195b3fb87d7848210", "0xcc491f589B45d4a3C679016195B3FB87D7848210"},
			true,
			"invalid precompile",
		},
		{
			"same address but one EIP-55 and other don't",
			[]string{"0xcc491f589b45d4a3c679016195b3fb87d7848210", "0xcc491f589B45d4a3C679016195B3FB87D7848210"},
			false,
			"",
		},
	}
	for _, tc := range testCases {

		slices.Sort(tc.precompiles)
		addrs, err := types.ValidatePrecompiles(tc.precompiles)

		if tc.expError {
			require.Error(t, err, tc.name)
			require.ErrorContains(t, err, tc.errContains)
		} else {
			require.NoError(t, err, tc.name)
			require.Equal(t, len(tc.precompiles), len(addrs), tc.name)
		}
	}
}
