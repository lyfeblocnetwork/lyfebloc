// Copyright Lyfeloop Inc.(Lyfebloc)
// SPDX-License-Identifier:ENCL-1.0(https://github.com/lyfeblocnetwork/lyfebloc/blob/main/LICENSE)
package coordinator

// Endpoint defines the identifiers for a chain's client, connection, and channel.
type Endpoint struct {
	ChainID      string
	ClientID     string
	ConnectionID string
	ChannelID    string
	PortID       string
}

// IBCConnection defines the connection between two chains.
type IBCConnection struct {
	EndpointA Endpoint
	EndpointB Endpoint
}