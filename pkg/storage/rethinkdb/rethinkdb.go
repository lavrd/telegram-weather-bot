package rethinkdb

import (
	"telegram-weather-bot/pkg/storage"

	"gopkg.in/gorethink/gorethink.v4"
)

type RethinkDB struct {
	session *gorethink.Session
}

func (r *RethinkDB) Close() error {
	return r.session.Close()
}

func New(dsn string) (*RethinkDB, error) {
	session, err := gorethink.Connect(gorethink.ConnectOpts{
		Address:  dsn,
		Database: storage.DBName,
	})
	if err != nil {
		return nil, err
	}

	return &RethinkDB{
		session: session,
	}, nil
}
