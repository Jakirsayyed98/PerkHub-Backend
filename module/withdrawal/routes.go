package withdrawal

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {

	app := api.Group("/withdrawal")
	app.Use(middlewear.UserMiddleware())
	app.POST("/addPaymentMethod", AddPaymentMethod)
	app.GET("/getPaymentMethod", GetPaymentMethodsByUserID)
	app.POST("/request", RequestWithdrawal)
	app.GET("/txnList", WithdrawalTxnList)

}
