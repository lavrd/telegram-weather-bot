package rethinkdb

import (
	"telegram-weather-bot/pkg/storage"

	"gopkg.in/gorethink/gorethink.v4"
)

const (
	databaseName  = "telegram-weather-bot"
	userTableName = "user"
)

const (
	FieldTelegramId = "telegramId"
)

type RethinkDB struct {
	session   *gorethink.Session
	userTable *gorethink.Term
}

func (r *RethinkDB) GetUser(telegramID int64) (*storage.User, error) {
	cur, err := r.userTable.Filter(gorethink.Row.Field(FieldTelegramId).Eq(telegramID)).Run(r.session)
	if err != nil {
		return nil, err
	}
	defer cur.Close()

	if cur.IsNil() {
		return nil, storage.ErrUserNotFound
	}

	user := &storage.User{}
	err = cur.One(user)
	return user, err
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

	if err := initialize(session); err != nil {
		return nil, err
	}

	userTable := gorethink.DB(databaseName).Table(userTableName)

	return &RethinkDB{
		session:   session,
		userTable: &userTable,
	}, nil
}
