package types

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
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}
