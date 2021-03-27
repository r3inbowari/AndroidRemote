package user

import (
	"CloudGameServer/db"
	"go.mongodb.org/mongo-driver/bson"
)

func Exist(name string) bool {
	db.MDB().C("User").Find(bson.M{ "username": "xiaoming"})
	return false
}

