package config

import "os"

// Config contains the parameters of an DB instance
type Config struct {
	DBVersion      uint32 `json:"db_version"`
	Path           string `json:"path"`
	MaxLogFileSize uint64 `json:"max_log_file_size"`

	MaxKeySize   uint64 `json:"max_key_size"`
	MaxValueSize uint64 `json:"max_value_size"`

	MaxConn  uint32      `json:"max_conn"`
	ReadOnly bool        `json:"read_only"`
	FileMode os.FileMode `json:"file_mode"`
}
