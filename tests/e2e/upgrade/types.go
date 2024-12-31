// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)
package upgrade

// ProposalVersion defines the version of the proposal format.
type ProposalVersion int

const (
    // UpgradeProposalV50 represents the proposal format for version 5.0.
    UpgradeProposalV50 ProposalVersion = iota
    // LegacyProposalPreV50 represents the legacy proposal format before version 5.0.
    LegacyProposalPreV50
    // LegacyProposalPreV46 represents the legacy proposal format before version 4.6.
    LegacyProposalPreV46
    // DefaultProposal represents the default proposal format for a new blockchain.
    DefaultProposal
)