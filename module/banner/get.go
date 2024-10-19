package banner

import (
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func GetBanners(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.BannerStore.GetAllBanners()

	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "fetched data successfully", "")
}
