package types

import "encoding/binary"

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
	// DelegatorKeyPrefix is the prefix to retrieve all Delegator
	DelegatorKeyPrefix = []byte{1, 0}
	// DelegatorKeyPrefixIndex2 is the prefix for a different key order for the DelegatorKeyPrefix
	DelegatorKeyPrefixIndex2 = []byte{1, 1}

	// DelegationEntriesKeyPrefix is the prefix to retrieve all DelegationEntries
	DelegationEntriesKeyPrefix = []byte{2, 0}

	// DelegationPoolDataKeyPrefix is the prefix to retrieve all DelegationPoolData
	DelegationPoolDataKeyPrefix = []byte{3, 0}
)

// DelegatorKey returns the store Key to retrieve a Delegator from the index fields
func DelegatorKey(poolId uint64, stakerAddress string, delegatorAddress string) []byte {
	return KeyPrefixBuilder{}.AInt(poolId).AString(stakerAddress).AString(delegatorAddress).Key
}

// DelegatorKeyIndex2 returns the store Key to retrieve a Delegator from the index fields
func DelegatorKeyIndex2(delegatorAddress string, poolId uint64, stakerAddress string) []byte {
	return KeyPrefixBuilder{}.AString(delegatorAddress).AInt(poolId).AString(stakerAddress).Key
}

// DelegationEntriesKey returns the store Key to retrieve a DelegationEntries from the index fields
func DelegationEntriesKey(poolId uint64, stakerAddress string, kIndex uint64) []byte {
	return KeyPrefixBuilder{}.AInt(poolId).AString(stakerAddress).AInt(kIndex).Key
}

// DelegationPoolDataKey returns the store Key to retrieve a DelegationPoolData from the index fields
func DelegationPoolDataKey(poolId uint64, stakerAddress string) []byte {
	return KeyPrefixBuilder{}.AInt(poolId).AString(stakerAddress).Key
}

type KeyPrefixBuilder struct {
	Key []byte
}

func (k KeyPrefixBuilder) AInt(n uint64) KeyPrefixBuilder {
	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, n)
	k.Key = append(k.Key, indexBytes...)
	k.Key = append(k.Key, []byte("/")...)
	return k
}

func (k KeyPrefixBuilder) AString(s string) KeyPrefixBuilder {
	k.Key = append(k.Key, []byte(s)...)
	k.Key = append(k.Key, []byte("/")...)
	return k
}
