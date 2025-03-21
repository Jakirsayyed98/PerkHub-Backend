package mobile

import (
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func GetHomePage(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.HomePageStore.GetHomePagedata()
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}
	settings.StatusOk(c, result, "data fetched successfully", "")
}
