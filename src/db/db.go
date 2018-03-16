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

// ClearTable - Used to clear test data from the database
func (db *Session) clearTable() {
	db.DB("ingress").C("testusers").RemoveAll(nil)
	db.DB("ingress").C("testwarehouses").RemoveAll(nil)
}
