package banner

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func GetBannerByID(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	id := c.Param("id")
	result, err := store.BannerStore.GetBannerbyId(id)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "fetched data successfully", "")
}

func GetBannerByCategoryID(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	id := c.Param("id")
	result, err := store.BannerStore.GetBannersByCategoryID(id)

	if err != nil && err.Error() != "no data found" {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "fetched data successfully", "")
}
