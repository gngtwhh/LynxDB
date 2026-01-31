package data

import (
	"fmt"
	"os"
	"path/filepath"
)

type LogFile struct {
	Fid          int
	File         *os.File
	Offset       int64
	MaxKeySize   uint64
	MaxValueSize uint64
}

func LoadLogFile(dirPath string, id int, maxKeySize, maxValueSize uint64, fileMode os.FileMode) (LogFile, error) {
	filePath := filepath.Join(dirPath, fmt.Sprintf("%d.data", id))
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, fileMode)
	if err != nil {
		return LogFile{}, err
	}
	return LogFile{
		Fid:          id,
		File:         file,
		Offset:       0,
		MaxKeySize:   maxKeySize,
		MaxValueSize: maxValueSize,
	}, nil
}
