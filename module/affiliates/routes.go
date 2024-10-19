package affiliates

import (
	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {
	app := api.Group("/admin")
	{
		app.GET("/cuelink-callback", CueLinkCallBack)
	}
}
