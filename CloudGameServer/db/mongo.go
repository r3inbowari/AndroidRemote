package db

import (
	bilicoin "CloudGameServer/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"reflect"
	"time"
)

type MongoExp struct {
	ctx    context.Context
	client *mongo.Client
	cancel context.CancelFunc
	opts   *options.ClientOptions
	dbn    string
}

var mdb MongoExp

func (db *MongoExp) setOpts(opts *options.ClientOptions) {
	db.opts = opts
}

func (db *MongoExp) connect() error {
	mdb.ctx, mdb.cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	if db.client, err = mongo.Connect(db.ctx, db.opts); err != nil {
		bilicoin.Info("mongodb connect failed")
	}
	return err
}

func (db *MongoExp) ping() error {
	var err error
	if err = mdb.client.Ping(db.ctx, readpref.Primary()); err != nil {
		bilicoin.Info("mongodb ping unavailable")
	} else {
		bilicoin.Info("mongodb connect succeed")
	}
	return err
}

func (db *MongoExp) disconnect() error {
	var err error
	if err = db.client.Disconnect(db.ctx); err != nil {
		bilicoin.Warn("mongodb disconnect failed")
		return err
	}
	db.cancel()
	return err
}

//func InitMongo() {
//	mdb = MongoExp{}
//
//	opts := options.Client().ApplyURI("mongodb://r3in.top:27017")
//	opts.SetAuth(options.Credential{
//		Username: "admin",
//		Password: "15946395951",
//	})
//	opts.SetMaxPoolSize(200)
//	opts.SetReadConcern(readconcern.Majority())
//
//	mdb.setOpts(opts)
//
//	var student Student
//	student.Value = 3.1415
//	student.Name = "xiao"
//
//	mdb.connect()
//	//time.Sleep(time.Second)
//	mdb.ping()
//	mdb.disconnect()
//	mdb.setDBName("testing")
//	// a :=mdb.get()
//
//	mdb.ping()
//	mdb.ping()
//	mdb.insertOne(student)
//	mdb.insertOne(student)
//	mdb.connect()
//	mdb.ping()
//	mdb.insertOne(student)
//	//mdb.insertOne(student)
//
//	// mdb.insertMany([]interface{}{student, student})
//	var users []Student
//	mdb.find(&users)
//	println("123")
//}

type Student struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (db *MongoExp) setDBName(dbn string) {
	db.dbn = dbn
}

// get collection
func (db *MongoExp) c(collectionName string) *mongo.Collection {
	return db.client.Database(db.dbn).Collection(collectionName)
}

func (db *MongoExp) bc(dbn string, collectionName string) *mongo.Collection {
	return db.client.Database(dbn).Collection(collectionName)
}

func (db *MongoExp) insertOne(idoc interface{}, opts ...*options.InsertOneOptions) error {
	colName := reflect.TypeOf(idoc).Name()
	var err error
	var doc interface{}
	if doc, err = bson.Marshal(idoc); err != nil {
		return err
	}
	_, err = db.c(colName).InsertOne(db.ctx, doc, opts...)
	return err
}

// only first type will predicted
func (db *MongoExp) insertMany(idocs []interface{}, opts ...*options.InsertManyOptions) error {
	var err error
	if len(idocs) == 0 {
		return errors.New("set len 0")
	}
	colName := reflect.TypeOf(idocs[0]).Name()
	var docs = make([]interface{}, len(idocs))
	for i, _ := range idocs {

		docs[i], err = bson.Marshal(idocs[i])
		if err != nil {
			return errors.New("marshal failed because some wrong caused by idocs")
		}
	}
	_, err = db.c(colName).InsertMany(db.ctx, docs, opts...)
	return err
}

func (db *MongoExp) find(result interface{}) error {

	//cur, err := db.c("Student").Find(db.ctx, bson.D{})
	//
	//if err != nil {
	//	return err
	//}
	//
	//
	//resultv := reflect.ValueOf(result)
	//
	//// ref for mgo.v2
	//// predict elem type
	//// result argument must be a slice address
	//if resultv.Kind() != reflect.Ptr || resultv.Elem().Kind() != reflect.Slice {
	//	panic("result argument must be a slice address")
	//}
	//
	//slicev := resultv.Elem()
	//slicev = slicev.Slice(0, slicev.Cap())
	//elemt := slicev.Type().Elem()
	//
	//i := 0
	//
	//for {
	//	if slicev.Len() == i {
	//		elemp := reflect.New(elemt)
	//		if !cur.Next(db.ctx) {
	//
	//		}
	//	}
	//}

	//cur, err := db.c("Student").Find(db.ctx, bson.D{})
	//
	//if err != nil {
	//	return err
	//}
	//
	//// target := make([]interface{}, 10, 100)
	//
	//var target []interface{}
	////vk := targets.([]interface{})
	////println(vk)
	//
	//ind := 0
	//
	//for cur.Next(db.ctx) {
	//	var result bson.D
	//	err := cur.Decode(&result)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	// do something with result....
	//	var s Student
	//	bsonBytes, _ := bson.Marshal(result)
	//	// baoli
	//
	//
	//	bson.Unmarshal(bsonBytes, &s)
	//
	//	target = append(target, s)
	//	//targets[ind] = s
	//	// println(s.Name)
	//
	//	//var st Student
	//	// bson.Unmarshal(result, st)
	//	ind++
	//}
	//
	//
	//if err := cur.Err(); err != nil {
	//	log.Fatal(err)
	//}
	//
	//targets = &target
	//return err

	//targets = 123
	return nil
}

func Hex(o interface{}) string {
	return o.(primitive.ObjectID).Hex()
}
