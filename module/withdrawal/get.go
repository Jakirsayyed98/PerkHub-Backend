package withdrawal

import (
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func WithdrawalTxnList(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)
	result, err := store.Withdrawal.WithdrawalTxnList(userId)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Withdrawal request list get successfully", "")
}
