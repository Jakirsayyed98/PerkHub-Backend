package transactions

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func AdminTransactionList(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	status := c.Param("status")
	request := request.NewAdminAffiliateTransactionsRequest()

	if err := c.ShouldBindJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AdminStore.AffiliateTransactions(request, status)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "Affiliate Transactions Fetched Successfully", "")
}
