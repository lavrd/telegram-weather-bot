package rethinkdb

import "gopkg.in/gorethink/gorethink.v4"

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

		err = gorethink.TableCreate(userTableName).Exec(session)
		return err
	}
	// database doesn't exist

	if err := gorethink.DBCreate(databaseName).Exec(session); err != nil {
		return err
	}
	err = gorethink.TableCreate(userTableName).Exec(session)
	return err
}
