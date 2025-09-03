package banner

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func CreateBanner(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewBanner()
	if err := request.Bind(c, awsInstance); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.BannerStore.SaveBanner(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondCreated(c, result, "Banner created Successfully", "")
}

func UpdateBanner(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewBanner()
	if err := request.Bind(c, awsInstance); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	result, err := store.BannerStore.UpdateBanner(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Banner Updated Successfully", "")
}

func DeleteBanner(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	id := c.Param("id")
	result, err := store.BannerStore.DeleteBanner(id)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Banner deleted Successfully", "")
}
