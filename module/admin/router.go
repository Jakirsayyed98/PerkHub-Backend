package admin

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/admin")

	// app.Use(middlewear.UserMiddleware())
	app.POST("/login", AdminLogin)
	app.POST("/register", RegisterAdmin)
	app.Use(middlewear.UserMiddleware())
	app.GET("/dashboard-data", middlewear.UserMiddleware(), GetAdminDashBoard)
	app.GET("/user-list", middlewear.UserMiddleware(), GetUserList)

}
