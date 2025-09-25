package notification

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(app *gin.RouterGroup) {
	admin := app.Group("/admin")
	admin.Use(middlewear.UserMiddleware())
	admin.POST("create-notification", CreateNotification)
	admin.POST("update-notification", UpdateNotification)
	admin.GET("get-notification-list", GetAdminNotificationList)
	admin.GET("send-notification/:id", SendNotification)
}
