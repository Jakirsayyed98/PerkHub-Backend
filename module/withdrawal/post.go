package withdrawal

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func RequestWithdrawal(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	req := request.NewWithdrawalRequest()
	if err := c.ShouldBindJSON(&req); err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}
	userId := c.MustGet("user_id").(string)
	result, err := store.Withdrawal.RequestWithdrawal(req, userId)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}
	settings.StatusOk(c, result, "withdrawal requested successfully", "")
}
