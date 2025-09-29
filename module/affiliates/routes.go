package affiliates

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	callback := api.Group("/")
	callback.GET("cuelink-callback", CueLinkCallBack)

	app := api.Group("/admin")
	app.Use(middlewear.UserMiddleware()) // Add any middleware if needed
	{
		app.POST("/create-affiliate", CreateAffiliate)
		app.POST("/update-affiliate", UpdateAffiliate)
		app.GET("/all-affiliates", ListAffiliates)
		app.POST("/active-deactive-affiliate", UpdateAffiliateFlag)
		app.POST("/delete-affiliate", DeleteAffiliate)
		app.POST("/affiliate", GetAffiliateByID)
	}
}
