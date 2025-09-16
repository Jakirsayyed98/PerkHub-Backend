package category

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {

	app := api.Group("/admin")
	app.Use(middlewear.UserMiddleware())
	app.POST("/create-category", CreateCategory)
	app.GET("/get-category", GetAllCategory)
	app.POST("/update-category", UpdateCategory)
	app.POST("/delete-category", DeleteCategory)
	app.POST("/active-deactive-category", ActiveDeactiveCategory)
	app.GET("/get-category/:id", GetCategoryByID)
}
