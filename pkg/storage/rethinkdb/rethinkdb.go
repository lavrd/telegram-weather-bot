package rethinkdb

import (
	"twb/pkg/storage"

	"gopkg.in/gorethink/gorethink.v4"
)

const (
	databaseName  = "telegram-weather-bot"
	userTableName = "user"
)

const (
	fieldTelegramId = "telegramId"
	fieldLang       = "lang"
	fieldUnits      = "units"
	fieldLocation   = "location"
	fieldLat        = "lat"
	fieldLon        = "lon"
)

const ConflictOpt = "update"

type RethinkDB struct {
	session   *gorethink.Session
	userTable *gorethink.Term
}

func (r *RethinkDB) Upsert(user *storage.User) error {
	args := map[string]interface{}{
		fieldTelegramId: user.TelegramID,
		fieldLang:       user.Lang,
		fieldUnits:      user.Units,
		fieldLocation:   user.Location,
		fieldLat:        user.Lat,
		fieldLon:        user.Lon,
	}
	opts := gorethink.InsertOpts{Conflict: ConflictOpt}
	err := r.userTable.Insert(args, opts).Exec(r.session)
	return err
}

func (r *RethinkDB) GetUser(telegramID int64) (*storage.User, error) {
	filter := gorethink.Row.Field(fieldTelegramId).Eq(telegramID)
	cur, err := r.userTable.Filter(filter).Run(r.session)
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

func initialize(session *gorethink.Session) error {
	cur, err := gorethink.DBList().Contains(databaseName).Run(session)
	if err != nil {
		return err
	}
	var ok bool
	if err := cur.One(&ok); err != nil {
		return err
	}
	defer cur.Close()

	// database exists
	if ok {
		cur, err = gorethink.TableList().Contains(userTableName).Run(session)
		if err != nil {
			return err
		}
		var ok bool
		if err := cur.One(&ok); err != nil {
			return err
		}
		defer cur.Close()

		// table exists
		if ok {
			return nil
		}
		// table doesn't exist

		opts := gorethink.TableCreateOpts{PrimaryKey: fieldTelegramId}
		err = gorethink.TableCreate(userTableName, opts).Exec(session)
		return err
	}
	// database doesn't exist

	if err := gorethink.DBCreate(databaseName).Exec(session); err != nil {
		return err
	}
	err = gorethink.TableCreate(userTableName).Exec(session)
	return err
}
