package category

import (
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.CategoryStore.GetAllCategory()
	if err != nil {
		settings.StatusBadRequest(c, err, "")
	}
	settings.StatusOk(c, result, "Category fetched Successfully", "")
}
