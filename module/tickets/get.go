package tickets

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func GetTicketsByUserId(c *gin.Context) {
	// Handler logic
	// Handler logic
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)

	result, err := store.TicketStore.GetTicketsByUserId(userId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Tickets retrieved successfully", "")
}

func GetTicketMessagesByTicketID(c *gin.Context) {
	// Handler logic
	// Handler logic
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	ticketId := c.Param("id")

	result, err := store.TicketStore.GetTicketMessagesByTicketID(ticketId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Tickets retrieved successfully", "")
}

func GetTicketsByStatus(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	ticketStatus := c.Param("status")

	result, err := store.TicketStore.GetTicketsByStatus(ticketStatus)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Tickets retrieved successfully", "")
}

func AdminTicketClosed(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	ticketId := c.Param("id")

	result, err := store.TicketStore.GetTicketsByStatus(ticketId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Tickets retrieved successfully", "")
}
