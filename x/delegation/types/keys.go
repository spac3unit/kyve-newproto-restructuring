package types

import (
	"github.com/KYVENetwork/chain/util"
)

const (
	// ModuleName defines the module name
	ModuleName = "delegation"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_delegation"
)

var (
	// DelegatorKeyPrefix is the prefix to retrieve all Delegator entries
	DelegatorKeyPrefix = []byte{1, 0}

	// DelegatorKeyPrefixIndex2 is the prefix for a different key order for the DelegatorKeyPrefix
	// TODO consider using memStore for second indexes, if they are only needed by the logic.
	DelegatorKeyPrefixIndex2 = []byte{1, 1}

	// DelegationEntriesKeyPrefix is the prefix to retrieve all DelegationEntries
	DelegationEntriesKeyPrefix = []byte{2, 0}

	// DelegationDataKeyPrefix ...
	DelegationDataKeyPrefix = []byte{3, 0}

	// DelegationSlashEntriesKeyPrefix ...
	DelegationSlashEntriesKeyPrefix = []byte{4, 0}

	// QueueKey ...
	QueueKey = []byte{5}

	// UndelegationQueueKeyPrefix ...
	UndelegationQueueKeyPrefix = []byte{6, 0}

	// UndelegationQueueKeyPrefixIndex2 ...
	UndelegationQueueKeyPrefixIndex2 = []byte{6, 1}

	// RedelegationCooldownPrefix ...
	RedelegationCooldownPrefix = []byte{7}
)

// DelegatorKey returns the store Key to retrieve a Delegator from the index fields
func DelegatorKey(stakerAddress string, delegatorAddress string) []byte {
	return util.GetByteKey(stakerAddress, delegatorAddress)
}

// DelegatorKeyIndex2 returns the store Key to retrieve a Delegator from the index fields
func DelegatorKeyIndex2(delegatorAddress string, stakerAddress string) []byte {
	return util.GetByteKey(delegatorAddress, stakerAddress)
}

// DelegationEntriesKey returns the store Key to retrieve a DelegationEntries from the index fields
func DelegationEntriesKey(stakerAddress string, kIndex uint64) []byte {
	return util.GetByteKey(stakerAddress, kIndex)
}

// DelegationDataKey returns the store Key to retrieve a DelegationPoolData from the index fields
func DelegationDataKey(stakerAddress string) []byte {
	return util.GetByteKey(stakerAddress)
}

// UndelegationQueueKey ...
func UndelegationQueueKey(kIndex uint64) []byte {
	return util.GetByteKey(kIndex)
}

// UndelegationQueueKeyIndex2 ...
func UndelegationQueueKeyIndex2(stakerAddress string, kIndex uint64) []byte {
	return util.GetByteKey(stakerAddress, kIndex)
}

func RedelegationCooldownKey(delegator string, block uint64) []byte {
	return util.GetByteKey(delegator, block)
}

func DelegationSlashEntriesKey(stakerAddress string, kIndex uint64) []byte {
	return util.GetByteKey(stakerAddress, kIndex)
}
