package db

import (
	bilicoin "CloudGameServer/utils"
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readconcern"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"reflect"
	"time"
)

// 没有
// opts aggr

type MongoExp struct {
	Ctx    context.Context
	Client *mongo.Client
	Cancel context.CancelFunc
	Opts   *options.ClientOptions
	dbn    string
}

var mdb_instance MongoExp

func (db *MongoExp) SetOpts(opts *options.ClientOptions) {
	db.Opts = opts
}

func (db *MongoExp) Connect() error {
	db.Ctx, db.Cancel = context.WithTimeout(context.Background(), 10*time.Second)
	var err error
	if db.Client, err = mongo.Connect(context.TODO(), db.Opts); err != nil {
		bilicoin.Info("mongodb connect failed")
	}
	return err
}

func (db *MongoExp) Ping() error {
	var err error
	if err = db.Client.Ping(context.TODO(), readpref.Primary()); err != nil {
		bilicoin.Info("mongodb ping unavailable")
	} else {
		bilicoin.Info("mongodb connect succeed")
	}
	return err
}

func (db *MongoExp) disconnect() error {
	var err error
	if err = db.Client.Disconnect(context.TODO()); err != nil {
		bilicoin.Warn("mongodb disconnect failed")
		return err
	}
	db.Cancel()
	return err
}

func InitMongo() {
	mdb_instance = MongoExp{}

	opts := options.Client().ApplyURI(bilicoin.GetConfig().MdbUrl)
	opts.SetAuth(options.Credential{
		Username: bilicoin.GetConfig().MdbUsername,
		Password: bilicoin.GetConfig().MdbPassword,
	})
	opts.SetMaxPoolSize(200)
	opts.SetReadConcern(readconcern.Majority())

	mdb_instance.SetOpts(opts)

	mdb_instance.Connect()
	mdb_instance.SetDBName(bilicoin.GetConfig().MdbName)
}

func MDB() *MongoExp {
	return &mdb_instance
}

type Student struct {
	Name  string  `json:"name"`
	Value float64 `json:"value"`
}

func (db *MongoExp) SetDBName(dbn string) {
	db.dbn = dbn
}

// get collection
func (db *MongoExp) C(collectionName string) *mongo.Collection {
	return db.Client.Database(db.dbn).Collection(collectionName)
}

func (db *MongoExp) BC(dbn string, collectionName string) *mongo.Collection {
	return db.Client.Database(dbn).Collection(collectionName)
}

func (db *MongoExp) InsertOne(idoc interface{}, opts ...*options.InsertOneOptions) error {
	colName := reflect.TypeOf(idoc).Elem().Name()
	var err error
	var doc interface{}
	if doc, err = bson.Marshal(idoc); err != nil {
		return err
	}
	_, err = db.C(colName).InsertOne(context.TODO(), doc, opts...)
	return err
}

// only first type will predicted
func (db *MongoExp) InsertMany(idocs []interface{}, opts ...*options.InsertManyOptions) error {
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
	_, err = db.C(colName).InsertMany(context.TODO(), docs, opts...)
	return err
}

func (db *MongoExp) FindAll(condition interface{}, result interface{}) error {
	// type exp panic !!
	cn := reflect.TypeOf(result).Elem().Elem().Name()
	cursor, err := db.C(cn).Find(context.TODO(), condition)
	if err != nil {
		return err
	}
	if err = cursor.All(db.Ctx, result); err != nil {
		return err
	}
	return nil
}

func (db *MongoExp) FindOne(condition interface{}, result interface{}) error {
	name := reflect.TypeOf(result).Elem().Name()
	res := db.C(name).FindOne(context.TODO(), condition)
	return res.Decode(result)
}

func Hex(o interface{}) string {
	return o.(primitive.ObjectID).Hex()
}

func (db *MongoExp) UpdateOne(filter primitive.ObjectID, content interface{}) error {
	name := reflect.TypeOf(content).Elem().Name()
	bs, _ := bson.Marshal(&content)
	_, err := db.C(name).ReplaceOne(context.TODO(),
		bson.M{"_id": filter},
		bs)
	return err
}

func (db *MongoExp) DeleteOne(content interface{}) error {
	name := reflect.TypeOf(content).Elem().Name()
	id := reflect.ValueOf(content).Elem().Field(0).Interface()
	_, err := db.C(name).DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}
	return err
}
