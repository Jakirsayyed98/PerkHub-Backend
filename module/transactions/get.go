package transactions

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func GetMiniAppTransaction(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)
	result, err := store.MiniAppTransactionStore.GetMiniAppTransaction(userId)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "Transaction Get Succesfully", "")
}
