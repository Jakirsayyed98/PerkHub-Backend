package miniapp

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func CreateMiniApp(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewMiniAppRequest()

	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	if err := request.Bind(c, awsInstance); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.MiniAppStore.CreateMiniApp(&request)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "MiniApp Work Successfully", "")

}

func UpdateActivateAndDeactive(c *gin.Context) {
	stores, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewActiveDeactiveminiAppReq()

	if err := c.ShouldBindBodyWithJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := stores.MiniAppStore.ActivateSomekey(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Updated Successfully", "")
}

func GetAllMiniApps(c *gin.Context) {
	stores, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := stores.MiniAppStore.GetAllMiniApps()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Get All MiniApp Successfully", "")
}

func SearchMiniApp(c *gin.Context) {
	stores, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewMiniAppSearchReq()

	if err := c.ShouldBindJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := stores.MiniAppStore.SearchMiniApps(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "fetched successfully", "")
}

func DeleteMiniApp(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	id := c.Param("id")
	result, err := store.MiniAppStore.DeletMniApp(id)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "MiniApp Deleted Successfully", "")
}

func GetStoresByCategory(c *gin.Context) {

	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	data := request.NewCategoryID()

	if err := c.ShouldBindJSON(data); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.MiniAppStore.GetStoresByCategory(data.CategoryId)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "Get MiniApp Successfully", "")

}

func GenrateSubid(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewGenrateMiniAppSubId()

	if err := c.ShouldBindJSON(request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)

	result, err := store.MiniAppStore.GenrateSubid(request.MiniAppName, userId)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}
	utils.RespondOK(c, result, "Generate SubID Successfully", "")
}
