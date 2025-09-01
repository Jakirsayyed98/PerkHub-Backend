package category

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func CreateCategory(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	data := request.NewCategory()

	// if err := data.Bind(c); err != nil {
	// 	settings.StatusBadRequest(c, err, "")
	// 	return
	// }
	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}
	if err := data.Bind(c, awsInstance); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.CategoryStore.SaveCategory(data)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
	}
	settings.StatusOk(c, result, "Category Saved Successfully", "")
}

func UpdateCategory(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	data := request.NewCategory()
	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}
	if err := data.Bind(c, awsInstance); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.CategoryStore.UpdateCategory(data)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
	}
	settings.StatusOk(c, result, "Category Updated Successfully", "")

}

func DeleteCategory(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	data := request.NewCategoryID()

	if err := c.ShouldBindJSON(data); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.CategoryStore.DeleteCategory(data.CategoryId)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
	}
	settings.StatusOk(c, result, "Category deleted Successfully", "")

}

func ActiveDeactiveCategory(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	data := request.NewCategoryActiveDeactive()

	if err := c.ShouldBindJSON(data); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.CategoryStore.ActiveDeactiveCategory(data)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
	}
	settings.StatusOk(c, result, "Category updated Successfully", "")

}
