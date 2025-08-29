package affiliates

import (
	"PerkHub/stores"

	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func ListAffiliates(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AffiliatesStore.ListAffiliates()
	if err != nil {
		utils.RespondInternalError(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Affiliates list fetched successfully", "")
}
