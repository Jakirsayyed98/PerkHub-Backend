package tickets

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	app := api.Group("/tickets")
	app.GET("/by_user", GetTicketsByUserId)
	app.GET("/:id", GetTicketMessagesByTicketID)
	app.POST("/create", CreateTicket)
}
