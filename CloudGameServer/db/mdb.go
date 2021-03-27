package db

import (
	bilicoin "CloudGameServer/utils"
	"gopkg.in/mgo.v2"
	"os"
)

var db *mgo.Database

func InitMongo() {
	session := InitMdbSession()
	db = session.DB(bilicoin.GetConfig().MdbName)
}

func MDB() *mgo.Database {
	return db
}

func InitMdbSession() *mgo.Session {
	session, err := mgo.Dial(bilicoin.GetConfig().MdbUrl)
	if err != nil {
		bilicoin.Error("mgo dial failed")
		os.Exit(0)
	}
	session.SetMode(mgo.Eventual, true)
	db := session.DB("admin")
	err = db.Login("admin", "15946395951")
	if err != nil {
		bilicoin.Error("mgo login failed")
		os.Exit(0)
	}
	session.SetPoolLimit(10)
	return session
}
