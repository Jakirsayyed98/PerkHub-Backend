package tickets

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func CreateTicket(c *gin.Context) {
	// Handler logic
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	// Validate and bind the request
	req := request.NewCreateTicketRequest()
	if err := c.ShouldBindJSON(req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)

	result, err := store.TicketStore.CreateTicket(req, userId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Ticket created successfully", "")
}
