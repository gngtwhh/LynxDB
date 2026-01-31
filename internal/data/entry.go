// package data contains the definition of the database data model
package data

import "hash/crc32"

type Entry struct {
	CheckSum uint32
	Key      []byte
	Value    []byte
}

func NewEntry(key, value []byte) Entry {
	return Entry{
		CheckSum: crc32.ChecksumIEEE(value),
		Key:      key,
		Value:    value,
	}
}
