package util

import "C"
import (
	"encoding/binary"
)

type KeyPrefixBuilder struct {
	Key []byte
}

func (k *KeyPrefixBuilder) AInt(n uint64) {
	indexBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(indexBytes, n)
	k.Key = append(k.Key, indexBytes...)
	k.Key = append(k.Key, []byte("/")...)
}

func (k *KeyPrefixBuilder) AString(s string) {
	k.Key = append(k.Key, []byte(s)...)
	k.Key = append(k.Key, []byte("/")...)
}

func GetByteKey(keys ...interface{}) []byte {
	builder := KeyPrefixBuilder{}
	for _, key := range keys {
		switch v := key.(type) {
		default:
			// TODO Maybe dangerous
			panic("Unsupported byte type")
		case uint64:
			builder.AInt(v)
		case string:
			builder.AString(v)
		}
	}
	return builder.Key
}
