package affiliates

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func CueLinkCallBack(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewCueLinkCallBackRequest()

	if err := request.Bind(c); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AffiliatesStore.CueLinkCallBack(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "Callback get Successfully", "")
}

func CreateAffiliate(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewCreateAffiliateRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AffiliatesStore.CreateAffiliate(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Affiliate created successfully", "")
}

func UpdateAffiliate(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewCreateAffiliateRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AffiliatesStore.UpdateAffiliate(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Affiliate updated successfully", "")
}

func DeleteAffiliate(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewCreateAffiliateRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AffiliatesStore.DeleteAffiliate(request.Id)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Affiliate deleted successfully", "")
}

func UpdateAffiliateFlag(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewCreateAffiliateRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AffiliatesStore.UpdateAffiliateFlag(request.Id, request.Status)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Affiliate flag updated successfully", "")
}

func GetAffiliateByID(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewCreateAffiliateRequest()
	if err := c.ShouldBindJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AffiliatesStore.GetAffiliateByID(request.Id)
	if err != nil {
		utils.RespondInternalError(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Affiliate fetched successfully", "")
}
