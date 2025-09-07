package offers

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/offers")
	app.Use(middlewear.UserMiddleware())
	app.POST("/search-by-store", SearchOffersByStoreName)
	app.GET("/homepage", HomePageOffers)

	admin := api.Group("/admin")
	admin.Use(middlewear.UserMiddleware())
	admin.GET("/offers-refresh", RefreshOffers)
	admin.GET("/get-offers-list/:type", GetAllActiveOffersList)
}
