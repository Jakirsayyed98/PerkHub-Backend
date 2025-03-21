package transactions

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/")

	app.Use(middlewear.UserMiddleware())

	app.GET("/getTxnList", GetMiniAppTransaction)

}
