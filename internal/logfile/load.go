package logfile

import (
	"fmt"
	"lynxdb/internal/data"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

func LoadLogFiles(path string, maxKeySize, maxValueSize uint64, fileMode os.FileMode) (oldFiles map[int]data.LogFile, maxID int, err error) {
	fnames, err := filepath.Glob(fmt.Sprintf("%s/*.data", path))
	if err != nil {
		return nil, 0, err
	}
	sort.Strings(fnames)

	for _, name := range fnames {
		id, _ := strconv.Atoi(strings.TrimSuffix(name, ".data"))
		oldFiles[id], err = data.LoadLogFile(path, id, maxKeySize, maxValueSize, fileMode)
		if err != nil {
			return
		}
	}
	return oldFiles, max(0, len(fnames)-1), nil
}
