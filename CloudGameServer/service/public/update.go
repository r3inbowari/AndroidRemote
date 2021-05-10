package public

import (
	"CloudGameServer/db"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Update struct {
	Id        primitive.ObjectID `bson:"_id" json:"_id"`
	Title     string             `json:"title"`
	Alias     string             `json:"alias"`
	Cover     string             `json:"cover"`
	Icon      string             `json:"icon"`
	Type      string             `json:"type"`
	AID       string             `json:"aid"`
	Update    bool               `json:"update"`
	New       bool               `json:"new"`
	Recommend bool               `json:"recommend"`
	Popular   bool               `json:"popular"`
	Desc      string             `json:"desc"`
}

func GetUpdate(c *gin.Context) {
	var update []Update
	if err := db.MDB().FindAll(bson.M{}, &update); err != nil {
		bilicoin.Warn("error find all")
	}
	c.JSON(200, bilicoin.ResponseWrapWithDate(200, "ok", update))
}

func AddUpdate(c *gin.Context) {
	var update Update
	update.Id = primitive.NewObjectID()
	err := c.Bind(&update)
	if err != nil {
		return
	}
	if err != db.MDB().InsertOne(&update) {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
	} else {
		c.JSON(200, bilicoin.ResponseOKWrap("ok"))
	}
}

func DelUpdate(c *gin.Context) {
	hex, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
		return
	}
	err = db.MDB().DeleteOne(&Update{Id: hex})
	if err != nil {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
		return
	}
	c.JSON(200, bilicoin.ResponseOKWrap("ok"))
}
