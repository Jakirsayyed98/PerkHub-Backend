package offers

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func SearchOffersByStoreName(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	req := request.NewSearchOffers()
	if err := c.ShouldBindJSON(req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.OffersStore.SearchOffersByStoreName(req.BrandName)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Offer get Successfully", "")
}

func HomePageOffers(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.OffersStore.OffersForHomePage()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Offer get Successfully", "")
}
