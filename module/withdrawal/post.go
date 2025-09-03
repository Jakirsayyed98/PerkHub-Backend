package withdrawal

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func RequestWithdrawal(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	req := request.NewWithdrawalRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	userId := c.MustGet("user_id").(string)
	result, err := store.Withdrawal.RequestWithdrawal(req, userId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "withdrawal requested successfully", "")
}

func AddPaymentMethod(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	req := request.NewAddPaymentMethodRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	userId := c.MustGet("user_id").(string)
	result, err := store.Withdrawal.AddPaymentMethod(req, userId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "payment method added successfully", "")
}
