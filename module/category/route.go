package category

import (
	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {

	app := api.Group("/admin")
	app.POST("/create-category", CreateCategory)
	app.GET("/get-category", GetAllCategory)
	app.POST("/update-category", UpdateCategory)
	app.POST("/delete-category", DeleteCategory)

}
