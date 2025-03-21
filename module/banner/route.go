package banner

import "github.com/gin-gonic/gin"

func Routes(api *gin.RouterGroup) {

	app := api.Group("/admin")
	app.POST("/create-banner", CreateBanner)
	app.POST("/update-banner", UpdateBanner)
	app.POST("/delete-banner/:id", DeleteBanner)
	app.GET("/get-banners", GetBanners)

}
