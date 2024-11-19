package miniapp

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func CreateMiniApp(c *gin.Context) {

	stores, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewMiniAppRequest()

	if err := request.Bind(c); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := stores.MiniAppStore.CreateMiniApp(&request)

	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "MiniApp Created Successfully", "")

}

func UpdateActivateAndDeactive(c *gin.Context) {
	stores, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewActiveDeactiveminiAppReq()

	if err := c.ShouldBindBodyWithJSON(request); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := stores.MiniAppStore.ActivateSomekey(request)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Updated Successfully", "")
}

func GetAllMiniApps(c *gin.Context) {
	stores, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := stores.MiniAppStore.GetAllMiniApps()
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Get All MiniApp Successfully", "")
}

func SearchMiniApp(c *gin.Context) {
	stores, err := stores.GetStores(c)
	if err != nil {
		settings.StatusForbidden(c, err)
		return
	}

	request := request.NewMiniAppSearchReq()

	if err := c.ShouldBindJSON(request); err != nil {
		settings.StatusNotFound(c, err, "")
		return
	}

	result, err := stores.MiniAppStore.SearchMiniApps(request)
	if err != nil {
		settings.StatusUnauthorized(c, err)
		return
	}

	settings.StatusOk(c, result, "Get All MiniApp Successfully", "")
}

func DeleteMiniApp(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewDeleteMiniApp()

	if err := c.ShouldBindJSON(request); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.MiniAppStore.DeletMniApp(request.Id)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "MiniApp Deleted Successfully", "")
}

func GetMiniAppBycategory(c *gin.Context) {

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

	result, err := store.MiniAppStore.GetMiniAppsBycategoryID(data.CategoryId)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}
	settings.StatusOk(c, result, "Get MiniApp Successfully", "")

}

func GenrateSubid(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewGenrateMiniAppSubId()

	if err := c.ShouldBindJSON(request); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id").(string)

	result, err := store.MiniAppStore.GenrateSubid(request.MiniAppId, userId)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}
	settings.StatusOk(c, result, "Generate SubID Successfully", "")
}
