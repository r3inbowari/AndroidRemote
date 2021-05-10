package public

import (
	"CloudGameServer/db"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Banner struct {
	Id    primitive.ObjectID `bson:"_id" json:"_id"`
	Title string             `json:"title"`
	Alias string             `json:"alias"`
	Cover string             `json:"cover"`
	Icon  string             `json:"icon"`
	AID   string             `json:"aid"`
}

func GetBanner(c *gin.Context) {
	var banners []Banner
	if err := db.MDB().FindAll(bson.M{}, &banners); err != nil {
		bilicoin.Warn("error find all")
	}
	c.JSON(200, bilicoin.ResponseWrapWithDate(200, "ok", banners))
}

func AddBanner(c *gin.Context) {
	var banner Banner
	banner.Id = primitive.NewObjectID()
	err := c.Bind(&banner)
	if err != nil {
		return
	}
	if err != db.MDB().InsertOne(&banner) {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
	} else {
		c.JSON(200, bilicoin.ResponseOKWrap("ok"))
	}
}

func DelBanner(c *gin.Context) {
	hex, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
		return
	}
	err = db.MDB().DeleteOne(&Banner{Id: hex})
	if err != nil {
		c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
		return
	}
	c.JSON(200, bilicoin.ResponseOKWrap("ok"))
}
