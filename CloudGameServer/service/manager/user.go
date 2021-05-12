package manager

import (
	"CloudGameServer/service/rtmsg"
	bilicoin "CloudGameServer/utils"
	"github.com/gin-gonic/gin"
)

func GetSessions(c *gin.Context) {

	var ret []rtmsg.UserSession

	rtmsg.WSSessionMap.Range(func(key, value interface{}) bool {
		ret = append(ret, value.(rtmsg.UserSession))
		return true
	})

	c.JSON(200, bilicoin.ResponseWrapWithDate(200, "", ret))
}

func KillSession(c *gin.Context) {
	stub := c.Param("stub")

	ret, ok := rtmsg.WSSessionMap.Load(stub)

	if ok {
		p := ret.(rtmsg.UserSession)
		err := p.Ws.Close()
		if err != nil {
			c.JSON(501, bilicoin.ResponseOKWrap("operation failed"))
			return
		}
	}
	c.JSON(200, bilicoin.ResponseOKWrap("ok"))
}
