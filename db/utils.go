package db

import (
	m "github.com/lavrs/telegram-weather-bot/model"
	"github.com/lavrs/telegram-weather-bot/utils/errors"
	r "gopkg.in/gorethink/gorethink.v3"
)

// create rethink db
func createTelegramDB() {
	_, err := r.DBCreate(db).RunWrite(session)
	errors.Check(err)
}

// create rethink table
func createUsersTable() {
	_, err := r.TableCreate(table).RunWrite(session)
	errors.Check(err)
}

// decode rethink quert result
func decodeOneBoolQueryResult(c *r.Cursor) (bool, error) {
	var res bool
	if err := c.One(&res); err != nil {
		return false, err
	}
	return res, nil
}

// get user from db
func getUser(telegramID int64) *m.DB {
	res, err := r.Table(table).Filter(
		r.Row.Field("telegramID").Eq(telegramID)).Run(session)
	errors.Check(err)
	defer res.Close()

	if res.IsNil() {
		return nil
	}

	var user m.DB
	err = res.One(&user)
	errors.Check(err)

	return &user
}

// get user ID from db
func getUserID(telegramID int64) *string {
	res, err := r.Table(table).Filter(
		r.Row.Field("telegramID").Eq(telegramID)).Field("id").Run(session)
	errors.Check(err)
	defer res.Close()

	if res.IsNil() {
		return nil
	}

	var ID string
	err = res.One(&ID)
	errors.Check(err)

	return &ID
}

// check db, table exists
func isTableAndDB() {
	query, err := r.DBList().Contains(db).Run(session)
	errors.Check(err)

	isDB, err := decodeOneBoolQueryResult(query)
	errors.Check(err)

	// db exists
	if isDB {
		query, err = r.TableList().Contains(table).Run(session)
		errors.Check(err)

		table, err := decodeOneBoolQueryResult(query)
		errors.Check(err)

		// table not exists
		if !table {
			createUsersTable()
			return
		}

		// table exists
		return
	}

	// db, table not exists
	createTelegramDB()
	createUsersTable()
	return
}
