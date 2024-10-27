package transactions

import (
	"PerkHub/settings"
	"PerkHub/stores"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMiniAppTransaction(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)
	result, err := store.MiniAppTransactionStore.GetMiniAppTransaction(userId)

	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}
	c.JSON(http.StatusOK, result)
	// settings.StatusOk(c, result, "Transaction Get Successfully", "")

}
