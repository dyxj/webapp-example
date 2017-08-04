package db

import (
	"gopkg.in/mgo.v2"
	"fmt"
)

// database variables
var iDB *mgo.Database
// "Read config file and set paths"
const urlDB = "mongodb://localhost:27017/webdev"
const defDB = "webdev"

// GetNewDB : Dials a new connection
func GetNewDB() *mgo.Database{
	s, err := mgo.Dial(urlDB)
	if err != nil {
		panic(err)
	}
	mDB := s.DB(defDB)
	return mDB
}

// Connect_iDB : Reassigns iDB
func Connect_iDB(){
	s, err := mgo.Dial(urlDB)
	if err != nil {
		panic(err)
	}
	iDB = s.DB(defDB)
}

// Close_iDB : Close iDB
func Close_iDB(){
	iDB.Session.Close()
	fmt.Println("Close iDB")
}

// CopyDB
func Copy_iDB() *mgo.Database {
	fmt.Println("Copy mongoDB")
	return iDB.Session.Copy().DB(defDB)
}

// CloseDB
func CloseDB(s *mgo.Session) {
	s.Close()
	fmt.Println("Close mongoDB")
}

// CopyDB
func CopyDB(m *mgo.Database) *mgo.Database {
	fmt.Println("Copy mongoDB")
	return m.Session.Copy().DB(m.Name)
}