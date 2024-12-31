// This file contains unit tests for the e2e package.
package upgrade

import (
    "testing"
    "github.com/stretchr/testify/require"
)

func TestCheckUpgradeProposalVersion(t *testing.T) {
    testCases := []struct {
        Name string
        Ver  string
        Exp  ProposalVersion
    }{
        {
            Name: "default proposal - version with whitespace - v1.0.0",
            Ver:  "\tv1.0.0 ",
            Exp:  DefaultProposal,
        },
        {
            Name: "default proposal - version without v - 1.0.0",
            Ver:  "1.0.0",
            Exp:  DefaultProposal,
        },
    }
    for _, tc := range testCases {
        t.Run(tc.Name, func(t *testing.T) {
            proposalVersion := CheckUpgradeProposalVersion(tc.Ver)
            require.Equal(t, tc.Exp, proposalVersion, "expected: %v, got: %v", tc.Exp, proposalVersion)
        })
    }
}

// TestRetrieveUpgradesList tests if the list of available upgrades in the codebase
// can be correctly retrieved
func TestRetrieveUpgradesList(t *testing.T) {
    upgradeList, err := RetrieveUpgradesList("../../../app/upgrades")
    require.NoError(t, err, "expected no error while retrieving upgrade list")
    require.Empty(t, upgradeList, "expected upgrade list to be empty for a new blockchain")
}