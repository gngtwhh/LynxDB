package main

import (
	"lynxdb"
	"lynxdb/internal/config"
)

type server struct {
	bind string
	db   *lynxdb.DB
}

func newserver(bind, path string) (*server, error) {
	// TODO: load config(or can be loaded in lynxdb.New())
	cfg := config.Config{}
	db, err := lynxdb.New(path, cfg)
	if err != nil {
		return nil, err
	}

	// TODO: register resp handler functions

	return &server{
		bind: bind,
		db:   db,
	}, nil
}
