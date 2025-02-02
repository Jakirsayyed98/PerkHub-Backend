package miniapp

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	user := api.Group("/")
	user.Use(middlewear.UserMiddleware())
	user.POST("/searchMiniApps", SearchMiniApp)
	user.POST("/getMiniAppByCategory", GetMiniAppBycategory)
	user.POST("/genrateSubID", GenrateSubid)

	app := api.Group("/admin")
	app.Use(middlewear.UserMiddleware())
	{
		app.POST("/create-miniapp", CreateMiniApp)
		app.POST("/active-deactive-miniapp", UpdateActivateAndDeactive)
		app.POST("/delete-miniapp", DeleteMiniApp)
		app.GET("/AllMiniApps", GetAllMiniApps)
	}
}
