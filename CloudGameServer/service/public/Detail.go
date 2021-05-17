package public

import (
	"CloudGameServer/db"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Detail struct {
	Id          primitive.ObjectID `bson:"_id" json:"_id"`
	Title       string             `json:"title"`
	Alias       string             `json:"alias"`
	Cover       string             `json:"cover"`
	Icon        string             `json:"icon"`
	Type        string             `json:"type"`
	AID         string             `json:"aid"`
	Update      bool               `json:"update"`
	New         bool               `json:"new"`
	Recommend   bool               `json:"recommend"`
	Popular     bool               `json:"popular"`
	Desc        string             `json:"desc"`
	Pack        string             `json:"pack"`
	Features    string             `json:"features"`
	Description string             `json:"description"`
	Snapshots   []string             `json:"snapshots"`
}

func GetDetail(c *gin.Context) {
	id := c.Param("id")
	var update Detail
	if err := db.MDB().FindOne(bson.M{"aid": id}, &update); err != nil {
		bilicoin.Warn("error find all")
	}
	c.JSON(200, bilicoin.ResponseWrapWithDate(200, "ok", update))
}

func GetDetails(c *gin.Context) {
	var update []Detail
	if err := db.MDB().FindAll(bson.M{}, &update); err != nil {
		bilicoin.Warn("error find all")
	}
	c.JSON(200, bilicoin.ResponseWrapWithDate(200, "ok", update))
}

func AddDetail(c *gin.Context) {
	var update Detail
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

func DelDetail(c *gin.Context) {
	hex, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
		return
	}
	err = db.MDB().DeleteOne(&Detail{Id: hex})
	if err != nil {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
		return
	}
	c.JSON(200, bilicoin.ResponseOKWrap("ok"))
}
