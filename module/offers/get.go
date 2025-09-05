package mobile

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func RefreshOffers(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.OffersStore.GetOffersRefresh()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "data fetched successfully", "")
}
