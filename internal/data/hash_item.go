package data

// Record refers to the hash table entries stored in memory that represent key-value pair mappings.
// It records the file and index where the latest record of this key-value pair is located.
type Record struct {
	FileID int   `json:"fileid"`
	Offset int64 `json:"offset"`
	Size   int64 `json:"size"`
}
