package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // EnergyKeyPrefix is the prefix to retrieve all Energy
	EnergyKeyPrefix = "Energy/value/"
)

// EnergyKey returns the store key to retrieve a Energy from the index fields
func EnergyKey(
index string,
) []byte {
	var key []byte
    
    indexBytes := []byte(index)
    key = append(key, indexBytes...)
    key = append(key, []byte("/")...)
    
	return key
}