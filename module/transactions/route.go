package transactions

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/")

	app.Use(middlewear.UserMiddleware())

	app.GET("/getTxnList", GetMiniAppTransaction)
	// app.POST("/affiliate-transactions", AdminTransactionList)

	admin := api.Group("/admin")
	admin.Use(middlewear.UserMiddleware())
	{
		admin.POST("/affiliate-transactions/:status", AdminTransactionList)
	}

}
