package miniapp

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

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
