package db

import "github.com/globalsign/mgo"

// Session - struct containing connection to database
type Session struct {
	*mgo.Session
}

//InitDB - establish connection to the DB
func InitDB(dataSourceName string) (*Session, error) {
	db, err := mgo.Dial(dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Session{db}, nil
}
