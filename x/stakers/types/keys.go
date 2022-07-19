package types

import "encoding/binary"

const (
	// ModuleName defines the module name
	ModuleName = "stakers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey is the message route for slashing
	RouterKey = ModuleName

	// QuerierRoute defines the module's query routing key
	QuerierRoute = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_stakers"
)

var (
	// StakerKeyPrefix is the prefix to retrieve all Staker
	StakerKeyPrefix = []byte{1}
)

const (
	MaxStakers        = 50
	DefaultCommission = "0.9"
)

// StakerKey returns the store Key to retrieve a Staker from the index fields
func StakerKey(staker string, poolId uint64) []byte {
	return KeyPrefixBuilder{}.AString(staker).AInt(poolId).Key
}

// TODO maybe outsource to util
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
