package user

import (
	"CloudGameServer/db"
	"CloudGameServer/service"
	bilicoin "CloudGameServer/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strconv"
)

//type CSRFImpl interface {
//
//}

func HandlerReg(context *gin.Context) {
	ui := &User{}
	if err := context.Bind(ui); err != nil {
		context.JSON(http.StatusInternalServerError, service.FailedResponse(err.Error(), bilicoin.InternalServerError))
		return
	}

	_, err := govalidator.ValidateStruct(ui)
	if err != nil {
		context.JSON(http.StatusInternalServerError, service.FailedResponse(err.Error(), bilicoin.StructValidateError))
		return
	}

	user := User{}
	user.Mobile = ui.Mobile
	if user.IsExist() {
		context.JSON(http.StatusForbidden, service.FailedResponse("already reg", bilicoin.UserExistError))
		return
	}

	if err = ui.CreateUser().Save(); err != nil {
		context.JSON(http.StatusInternalServerError, service.FailedResponse(err.Error(), bilicoin.DatabaseError))
		return
	}

	context.JSON(http.StatusOK, service.SucceedResponse("ok"))
}

func HandlerLogin(context *gin.Context) {
	ui := &User{}
	var err error
	if err = context.Bind(ui); err != nil {
		context.JSON(http.StatusInternalServerError, service.FailedResponse(err.Error(), bilicoin.InternalServerError))
		return
	}
	var token string
	if token, err = ui.Login(); err != nil {
		context.JSON(http.StatusOK, service.FailedResponse(err.Error(), bilicoin.UserTokenCreateError))
		return
	}
	context.JSON(http.StatusOK, service.SucceedResponse(token, bilicoin.UserLoginSucceed))
}

// HandlerInfo unsafe operation
func HandlerInfo(context *gin.Context) {
	bearer := context.GetHeader("Authorization")
	token := bilicoin.ParseToken(bearer)
	userInfo, err := GetInfoByToken(token)
	if err != nil {
		context.JSON(http.StatusOK, service.FailedResponse(err.Error(), bilicoin.InternalServerError))
	}
	context.JSON(http.StatusOK, service.SucceedResponse(userInfo, bilicoin.UserLoginSucceed))
}

func HandlerPay(context *gin.Context) {
	point, err := strconv.Atoi(context.Query("point"))
	if err != nil {
		context.JSON(http.StatusOK, service.FailedResponse(err.Error(), bilicoin.InternalServerError))
		return
	}
	bearer := context.GetHeader("Authorization")
	token := bilicoin.ParseToken(bearer)
	info, err := GetInfoByToken(token)

	ud := User{}
	if err := db.MDB().FindOne(bson.M{"mobile": info.Mobile}, &ud); err != nil {
		context.JSON(http.StatusOK, service.FailedResponse("not found", bilicoin.InternalServerError))
	}
	ud.Point += point * 60// demo，没有支付服务，直接加就算了

	err = ud.Update()
	if err != nil {
		context.JSON(http.StatusOK, service.FailedResponse("充值失败", bilicoin.InternalServerError))
		return
	}
	context.JSON(http.StatusOK, service.SucceedResponse("ok", bilicoin.UserChargeSucceed))
}
