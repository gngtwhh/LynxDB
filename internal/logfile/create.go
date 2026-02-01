package logfile

import (
	"fmt"
	"lynxdb/internal/data"
	"os"
	"path/filepath"
)

func CreateLogFile(path string, id int, maxKeySize, maxValueSize uint64, fileMode os.FileMode) (lf data.LogFile, err error) {
	fname := filepath.Join(path, fmt.Sprintf("%d.data", id))
	f, err := os.OpenFile(fname, os.O_CREATE|os.O_RDWR, fileMode)
	if err != nil {
		return data.LogFile{}, err
	}
	lf = data.LogFile{
		Fid:          id,
		File:         f,
		MaxKeySize:   maxKeySize,
		MaxValueSize: maxValueSize,
	}
	return lf, nil
}
