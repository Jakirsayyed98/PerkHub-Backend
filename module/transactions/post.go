package transactions

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"
	"fmt"

	"github.com/gin-gonic/gin"
)

func AdminTransactionList(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewAdminAffiliateTransactionsRequest()

	if err := c.ShouldBindJSON(request); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.AdminStore.AffiliateTransactions(request)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}
	fmt.Println(result)
	settings.StatusOk(c, result, "Callback get Successfully", "")
}
