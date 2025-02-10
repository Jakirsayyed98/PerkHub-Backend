package banner

import (
	"PerkHub/settings"
	"PerkHub/stores"
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetBanners(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	id := c.Param("id")
	fmt.Println("Category_ID", id)
	result, err := store.BannerStore.GetBannersByCategoryID(id)

	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "fetched data successfully", "")
}

func GetBannerCategory(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.BannerStore.GetBannerCategory()
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Banner Category fetch Successfully", "")
}
