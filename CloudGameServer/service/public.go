package service

import (
	"CloudGameServer/db"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type Banner struct {

}

func GetBanner(c *gin.Context) {
	var banners []Banner


	if err := db.MDB().FindAll(bson.M{}, &banners); err != nil {
		bilicoin.Warn("error find all")
	}



}
