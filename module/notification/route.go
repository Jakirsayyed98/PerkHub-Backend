package notification

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/")
	app.Use(middlewear.UserMiddleware())
	app.GET("get-notification-history", GetUserNotificationHistory)

	admin := api.Group("/admin")
	admin.Use(middlewear.UserMiddleware())
	admin.POST("create-notification", CreateNotification)
	admin.POST("update-notification", UpdateNotification)
	admin.GET("get-notification-list", GetAdminNotificationList)
	admin.GET("send-notification/:id", SendNotification)
	admin.GET("get-notification-by-id/:id", GetNotificationById)
}
