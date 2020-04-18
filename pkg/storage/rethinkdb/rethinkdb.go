package rethinkdb

import (
	"gopkg.in/gorethink/gorethink.v4"
)

const (
	databaseName  = "telegram-weather-bot"
	userTableName = "user"
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
		Database: databaseName,
	})
	if err != nil {
		return nil, err
	}

	return &RethinkDB{
		session: session,
	}, nil
}
