package miniapp

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	user := api.Group("/")
	user.Use(middlewear.UserMiddleware())
	user.POST("/searchMiniApps", SearchMiniApp)
	user.POST("/getStoresByCategory", GetStoresByCategory)
	user.POST("/genrateSubID", GenrateSubid)

	app := api.Group("/admin")
	app.Use(middlewear.UserMiddleware())
	{
		app.POST("/create-miniapp", CreateMiniApp)
		app.POST("/get-store-by-id", GetStoreByID)
		app.POST("/active-deactive-miniapp", UpdateActivateAndDeactive)
		app.GET("/delete-miniapp/:id", DeleteMiniApp)
		app.GET("/AllMiniApps", GetAllMiniApps)
		app.GET("/miniAppCategory-refresh", GetStoresCategoryRefresh)
		app.GET("/miniApp-refresh", GetStoresRefresh)

	}
}
