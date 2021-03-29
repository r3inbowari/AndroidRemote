package test

import (
	"CloudGameServer/db"
	user2 "CloudGameServer/service/user"
	bilicoin "CloudGameServer/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func dbt() *db.MongoExp {
	md := db.MongoExp{}
	opts := options.Client().ApplyURI(bilicoin.GetConfig().MdbUrl)
	opts.SetAuth(options.Credential{
		Username: bilicoin.GetConfig().MdbUsername,
		Password: bilicoin.GetConfig().MdbPassword,
	})
	opts.SetMaxPoolSize(200)
	opts.SetReadConcern(readconcern.Majority())
	md.SetOpts(opts)
	md.SetDBName(bilicoin.GetConfig().MdbName)
	_ = md.Connect()
	err := md.Ping()
	println(err != nil)
	return &md
}

func TestFindAll(t *testing.T) {
	md := dbt()
	var user []user2.User
	md.FindAll("User", &user, bson.M{})
	println(len(user))
}

func TestAdd(t *testing.T) {
	db.InitMongo()

	var user = user2.User{
		Username: "aaa111111111",
		Mobile:   "aaa",
		Avatar:   "aaa",
		Password: "aaa",
		Status:   0,
	}

	err := user.Save()
	println(err)
}

func TestUpdate(t *testing.T) {
	db.InitMongo()
	id, _ := primitive.ObjectIDFromHex("605f5e6c4cf3b7023774538b")
	var user = user2.User{
		Id:     id,
		Avatar: "hello1",
	}
	_ = user.Update()
}
