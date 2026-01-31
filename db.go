package lynxdb

import (
	"errors"
	"lynxdb/internal/config"
	"lynxdb/internal/data"
	"lynxdb/internal/logfile"
)

type DB struct {
	config *config.Config
	keydir data.KeyDir

	logFiles   map[int]data.LogFile
	activeFile data.LogFile
}

// New create a new database instance
func New(dbPath string, cfg config.Config) (*DB, error) {
	// TODO: config load
	db := &DB{
		config: &cfg,
		// keydir:   make(data.KeyDir),
		// logFiles: make(map[int]data.LogFile),
	}
	if err := db.setup(); err != nil {
		return nil, err
	}
	return db, nil
}

// setup load all log files and init hash map to initialize the DB
func (db *DB) setup() error {
	oldFiles, maxID, err := logfile.LoadLogFiles(
		db.config.Path, db.config.MaxKeySize, db.config.MaxValueSize, db.config.FileMode,
	)
	if err != nil {
		return err
	}
	db.logFiles = make(map[int]data.LogFile)
	for _, f := range oldFiles {
		db.logFiles[f.Fid] = f
	}
	if err := db.loadKeyDir(maxID); err != nil {
		return err
	}

	if err := db.createActiveFile(maxID); err != nil {
		return err
	}

	return nil
}

func (db *DB) loadKeyDir(maxID int) error {
	return errors.New("implement me")
}
func (db *DB) createActiveFile(maxID int) error {
	return errors.New("implement me")
}
