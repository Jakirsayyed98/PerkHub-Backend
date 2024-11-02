package withdrawal

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {

	app := api.Group("/withdrawal")
	app.Use(middlewear.UserMiddleware())
	app.POST("/request", RequestWithdrawal)
	app.GET("/txnList", WithdrawalTxnList)

}
