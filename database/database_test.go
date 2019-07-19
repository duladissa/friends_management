package database

import "testing"

func Test_database_connection(t *testing.T) {
	uri := "mongodb://localhost:27017"
	dbName := "testing"
	database, _ := CreateDatabase(uri, dbName)

	if database == nil {
		t.Error("Database Connection Failed")
	}

	mgoDB, _ := database.getDatabase(dbName)
	if mgoDB == nil {
		t.Error("MGO Database Connection Failed")
	}

	if mgoDB.Name != dbName {
		t.Error("Error connecting to the mongo db instance")
	}
}
