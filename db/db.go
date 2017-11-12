package db

import (
	"gopkg.in/mgo.v2"
)

var (
	mgoSession *mgo.Session
)

func Init() *mgo.Database {
	var err error

	mgoSession, err = mgo.Dial("mongodb://localhost:27017")

	if err != nil {
		panic(err)
	}

	return mgoSession.DB("smsMiddleware")

}

func CloseSession() {
	mgoSession.Close()
}
