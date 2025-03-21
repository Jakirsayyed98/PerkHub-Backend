package banner

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/admin")
	app.Use(middlewear.UserMiddleware())
	app.POST("/create-banner-category", CreateBannerCategory)
	app.GET("/banner-category", GetBannerCategory)
	app.POST("/create-banner", CreateBanner)
	app.POST("/update-banner", UpdateBanner)
	app.POST("/delete-banner/:id", DeleteBanner)
	app.GET("/get-banners/:id", GetBanners)

}
