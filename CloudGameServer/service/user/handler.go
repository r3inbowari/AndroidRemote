package user

import (
	"CloudGameServer/service"
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
		context.JSON(http.StatusInternalServerError, service.FailedResponse(err.Error()))
		return
	}

	_, err := govalidator.ValidateStruct(ui)
	if err != nil {
		context.JSON(http.StatusInternalServerError, service.FailedResponse(err.Error()))
		return
	}

	user := User{}
	user.Mobile = ui.Mobile
	if user.IsExist() {
		context.JSON(http.StatusForbidden, service.FailedResponse("already reg"))
		return
	}

	if err = user.CreateUser().Save(); err != nil {
		context.JSON(http.StatusInternalServerError, service.FailedResponse(err.Error()))
		return
	}

	context.JSON(http.StatusOK, service.SucceedResponse("ok"))
}
