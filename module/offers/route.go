package mobile

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	admin := api.Group("/admin")
	admin.Use(middlewear.UserMiddleware())
	admin.GET("/offers-refresh", RefreshOffers)

}
