package data

import (
	"encoding/binary"
	"hash/crc32"
	"io"
)

// Error definitions
var (
	ErrKeySizeExceeded   = &Error{"key size exceeds maximum allowed"}
	ErrValueSizeExceeded = &Error{"value size exceeds maximum allowed"}
	ErrChecksumMismatch  = &Error{"checksum verification failed"}
)

// Error represents a data layer error
type Error struct {
	msg string
}

func (e *Error) Error() string {
	return e.msg
}

// Keydir is a hash table version implemented using the native map function of Go language.
type KeyDir map[string]*Entry

func (kd *KeyDir) LoadFromLogFile(lf LogFile) error {
	if *kd == nil {
		*kd = make(KeyDir)
	}

	if _, err := lf.File.Seek(0, io.SeekStart); err != nil {
		return err
	}

	reader := io.Reader(lf.File)

	// TODO: use variable-length integers to store the length fields
	// Read entries until EOF
	for {
		// Read checksum (4 bytes)
		var checksum uint32
		if err := binary.Read(reader, binary.LittleEndian, &checksum); err != nil {
			if err == io.EOF {
				break
			}
			return err
		}

		// Read key length (4 bytes)
		var keyLen uint32
		if err := binary.Read(reader, binary.LittleEndian, &keyLen); err != nil {
			return err
		}

		// Read value length (4 bytes)
		var valueLen uint32
		if err := binary.Read(reader, binary.LittleEndian, &valueLen); err != nil {
			return err
		}

		// Validate sizes against maximum allowed
		if uint64(keyLen) > lf.MaxKeySize {
			return ErrKeySizeExceeded
		}
		if uint64(valueLen) > lf.MaxValueSize {
			return ErrValueSizeExceeded
		}

		// Read key data
		key := make([]byte, keyLen)
		if _, err := io.ReadFull(reader, key); err != nil {
			return err
		}

		// Read value data
		value := make([]byte, valueLen)
		if _, err := io.ReadFull(reader, value); err != nil {
			return err
		}

		// Verify checksum
		calculatedChecksum := crc32.ChecksumIEEE(value)
		if calculatedChecksum != checksum {
			return ErrChecksumMismatch
		}

		entry := &Entry{
			CheckSum: checksum,
			Key:      key,
			Value:    value,
		}
		(*kd)[string(key)] = entry
	}
	return nil
}
