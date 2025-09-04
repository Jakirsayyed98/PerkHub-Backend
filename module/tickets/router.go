package tickets

import (
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(api *gin.RouterGroup) {
	app := api.Group("/tickets")
	app.GET("/by_user", GetTicketsByUserId)
	app.GET("/:id", GetTicketMessagesByTicketID)
	app.POST("/create", CreateTicket)
	app.POST("/send-msg", SendMsg)

	admin := api.Group("/admin")
	admin.GET("/tickets/:status", GetTicketsByStatus)
	admin.GET("/tickets-msg/:id", GetTicketMessagesByTicketID)
	admin.POST("/tickets-msg/reply", AdminSentTicketReply)
	admin.GET("/ticket-closed/:id", AdminTicketClosed)

}
