package user

import (
	"CloudGameServer/service"
	bilicoin "CloudGameServer/utils"
	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"net/http"
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

func HandlerInfo(context *gin.Context) {
	bearer := context.GetHeader("Authorization")
	token := bilicoin.ParseToken(bearer)

	// var err error
	userInfo, err := GetInfoByToken(token)
	if err != nil {
		context.JSON(http.StatusOK, service.FailedResponse(err.Error(), bilicoin.InternalServerError))
	}
	context.JSON(http.StatusOK, service.SucceedResponse(userInfo, bilicoin.UserLoginSucceed))
}
