// package data contains the definition of the database data model
package data

import "hash/crc32"

type Entry struct {
	CheckSum uint32
	Key      []byte // The length of the key is implicitly contained within the Key field
	Value    []byte // The length of the value is implicitly contained within the Value field
}

func NewEntry(key, value []byte) Entry {
	return Entry{
		CheckSum: crc32.ChecksumIEEE(value),
		Key:      key,
		Value:    value,
	}
}
