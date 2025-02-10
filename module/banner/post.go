package banner

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func CreateBannerCategory(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewBannerCategory()
	if err := c.ShouldBindJSON(&request); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.BannerStore.SaveBannerCategory(request)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusCreated(c, result, "Banner Category created Successfully", "")
}

func CreateBanner(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusInternalServerError(c, err, "")
		return
	}

	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewBanner()
	if err := request.Bind(c, awsInstance); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.BannerStore.SaveBanner(request)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusCreated(c, result, "Banner created Successfully", "")
}

func UpdateBanner(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewBanner()
	if err := request.Bind(c, awsInstance); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}
	result, err := store.BannerStore.UpdateBanner(request)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Banner Updated Successfully", "")
}

func DeleteBanner(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}
	id := c.Param("id")
	result, err := store.BannerStore.DeleteBanner(id)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Banner deleted Successfully", "")
}
