package database

import (
	"fmt"
	"sync"

	"gopkg.in/mgo.v2"
)

//Database ... Database model
type Database struct {
	session *mgo.Session
	name    string
}

var singleton *Database
var err error
var once sync.Once

//DBWTimeout ... Milliseconds to wait for db server before timing out
const DBWTimeout = 5000 //5 Seconds

//CreateDatabase ... Create single instance of the database
func CreateDatabase(uri string, dbName string) (*Database, error) {
	if uri == "" {
		return nil, fmt.Errorf("No DB connection string provided")
	}
	once.Do(func() {
		//tlsConfig := &tls.Config{}

		dialInfo, err := mgo.ParseURL(uri)
		if err != nil {
			fmt.Println("Database ParseURL failure", err.Error())
		}
		// dialInfo.DialServer = func(addr *mgo.ServerAddr) (net.Conn, error) {
		// 	conn, err := tls.Dial("tcp", addr.String(), tlsConfig)
		// 	return conn, err
		// }
		session, err1 := mgo.DialWithInfo(dialInfo)
		if err1 != nil {
			fmt.Println("Database Connect Access failure", err1.Error())
		}
		fmt.Println("Database Access success")
		singleton = &Database{session: session, name: dbName}
	})

	return singleton, nil
}

//GetSession ... Return the mongo db and session
func (database *Database) GetSession() (*mgo.Database, *mgo.Session) {
	fmt.Println("\nGetSession - Creating database session")
	session := database.session.Copy()
	// Error check on every access
	session.SetSafe(&mgo.Safe{WTimeout: DBWTimeout})

	db := session.DB(database.name)
	return db, session
}
