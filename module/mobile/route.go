package mobile

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/")
	app.Use(middlewear.UserMiddleware())
	app.GET("/getHomePage", GetHomePage)
	app.POST("/updateNotificationToken", UpdateNotificationToken)

}
