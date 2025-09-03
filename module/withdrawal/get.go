package withdrawal

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func WithdrawalTxnList(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)
	result, err := store.Withdrawal.WithdrawalTxnList(userId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Withdrawal request list get successfully", "")
}

func GetPaymentMethodsByUserID(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)
	result, err := store.Withdrawal.GetPaymentMethodsByUserID(userId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Withdrawal request list get successfully", "")
}
