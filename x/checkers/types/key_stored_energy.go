package types

import "encoding/binary"

var _ binary.ByteOrder

const (
    // StoredEnergyKeyPrefix is the prefix to retrieve all StoredEnergy
	StoredEnergyKeyPrefix = "StoredEnergy/value/"
)

// StoredEnergyKey returns the store key to retrieve a StoredEnergy from the index fields
func StoredEnergyKey(
index string,
) []byte {
	var key []byte
    
    indexBytes := []byte(index)
    key = append(key, indexBytes...)
    key = append(key, []byte("/")...)
    
	return key
}