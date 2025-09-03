package category

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func GetAllCategory(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.CategoryStore.GetAllCategory()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "Category fetched Successfully", "")
}

func GetCategoryByID(c *gin.Context) {
	id := c.Param("id")

	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.CategoryStore.GetCategoryByID(id)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "Category fetched Successfully", "")
}
