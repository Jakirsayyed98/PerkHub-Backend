package tickets

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"
	"fmt"

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

func SendMsg(c *gin.Context) {
	// Handler logic
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	// Validate and bind the request
	req := request.NewSendTicketMsg()
	if err := c.ShouldBindJSON(req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)

	result, err := store.TicketStore.SendTicketMessage(req, userId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Ticket message sent successfully.", "")
}

func AdminSentTicketReply(c *gin.Context) {
	// Handler logic
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	// Validate and bind the request
	req := request.NewAdminReplyTicketMsg()
	if err := c.ShouldBindJSON(req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)

	result, err := store.TicketStore.AdminSentTicketReply(req, userId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Ticket message sent successfully.", "")
}

func ChangeTicketStatus(c *gin.Context) {
	// Handler logic
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	// Validate and bind the request
	req := request.NewChangeTicketStatusRequest()
	if err := c.ShouldBindJSON(req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	fmt.Println(req)
	err = store.TicketStore.ChangeTicketStatus(req)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, nil, "Ticket status changed successfully.", "")
}
