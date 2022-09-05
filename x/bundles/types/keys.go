package types

import "github.com/KYVENetwork/chain/util"

const (
	// ModuleName defines the module name
	ModuleName = "bundles"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_bundles"
)

var (
	// BundleKeyPrefix ...
	BundleKeyPrefix = []byte{1}
	// FinalizedBundlePrefix ...
	FinalizedBundlePrefix = []byte{2}
)

// BundleProposalKey ...
func BundleProposalKey(poolId uint64) []byte {
	return util.GetByteKey(poolId)
}

// FinalizedBundleKey ...
func FinalizedBundleKey(poolId uint64, id uint64) []byte {
	return util.GetByteKey(poolId, id)
}
