package types

import (
	cosmostypes "github.com/cosmos/cosmos-sdk/types"
)
const (
	// ModuleName defines the module name
	ModuleName = "promos"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_promos"
)

var (
	ParamsKey = []byte("p_promos")
	HeaderPrefix = []byte{0x01}
)


func HeaderKey(height uint64) []byte {
	return cosmostypes.Uint64ToBigEndian(height)
}

func KeyPrefix(p string) []byte {
	return []byte(p)
}
