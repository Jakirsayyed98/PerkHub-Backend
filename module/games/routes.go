package games

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/admin")
	app.GET("refresh-games", RefreshGames)
	app.GET("refresh-games-categories", RefreshGameCategories)
	api.Use(middlewear.UserMiddleware())
	{
		api.GET("getgames-categories", GetGameCategories)
		api.GET("getgames", GetGames)
		api.POST("getgames-bycategory", GameByCategory)
		api.POST("search-game", GameSearch)
		api.GET("get-popular-games", GetPopulargames)
		api.GET("get-trending-games", GetTrendingGames)

	}
}
