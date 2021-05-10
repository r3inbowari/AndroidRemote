package public

import (
	"CloudGameServer/db"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Hot struct {
	Id    primitive.ObjectID `bson:"_id" json:"_id"`
	Title string             `json:"title"`
	Alias string             `json:"alias"`
	Cover string             `json:"cover"`
	Icon  string             `json:"icon"`
	Type  string             `json:"type"`
	Tags  []string           `json:"tags"`
	AID   string             `json:"aid"`
}

func GetHot(c *gin.Context) {
	var hot []Hot
	if err := db.MDB().FindAll(bson.M{}, &hot); err != nil {
		bilicoin.Warn("error find all")
	}
	c.JSON(200, bilicoin.ResponseWrapWithDate(200, "ok", hot))
}

func AddHot(c *gin.Context) {
	var hot Hot
	hot.Id = primitive.NewObjectID()
	err := c.Bind(&hot)
	if err != nil {
		return
	}
	if err != db.MDB().InsertOne(&hot) {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
	} else {
		c.JSON(200, bilicoin.ResponseOKWrap("ok"))
	}
}

func DelHot(c *gin.Context) {
	hex, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
		return
	}
	err = db.MDB().DeleteOne(&Hot{Id: hex})
	if err != nil {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
		return
	}
	c.JSON(200, bilicoin.ResponseOKWrap("ok"))
}
