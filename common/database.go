package common

import (
	"net/http"

	"github.com/gorilla/context"
	mgo "gopkg.in/mgo.v2"
)

const databaseContextName = "database"
const databaseName = "pipenator"
const databaseUser = "pipenator"
const databasePass = "pipenator123"
const databaseUrl = "ds263670.mlab.com:63670/"

func CreateDB() (*mgo.Session, error) {
	db, err := mgo.Dial("mongodb://" + databaseUser + ":" + databasePass + databaseUrl + databaseName)

	if err != nil {
		return nil, err
	}

	defer db.Close()

	return db, nil
}

func GetDBFromContext(r *http.Request) *mgo.Database {
	return context.Get(r, databaseContextName).(*mgo.Session).DB(databaseName)
}
