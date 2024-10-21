// Package utils provides utility functions for handling Ethereum transactions,
// converting JavaScript-like object strings to valid JSON, and parsing incoming messages
// from the Arbitrum sequencer feed.
//
// This package is designed to support the processing and manipulation of data related
// to Ethereum transactions and feed messages, with a focus on the Arbitrum Layer 2 solution.
// It includes functions for converting input formats, calculating Merkle roots, and parsing
// JSON data into defined Go structures. Additionally, the package defines useful constants
// related to the Arbitrum network.
package utils

// Constants related to the Arbitrum network.
const (
	// ArbiturmChainId represents the Chain ID for the Arbitrum network (Mainnet).
	ArbiturmChainId = 42161

	// ArbiturmGenesisBlockNumber is the block number of Etheruem where the Nitro starts.
	// On Arbitrum One, Retrieving L2 Block Number can be done by adding the Arbitrum One genesis block number (22207817) to the sequence number of the feed message.
	ArbiturmGenesisBlockNumber uint64 = 22207817
)
