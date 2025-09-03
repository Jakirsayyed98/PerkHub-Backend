package admin

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func GetAdminDashBoard(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AdminStore.GetAdminDashBoardData()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "Get Dashboard Data Successfully", "")
}

func GetUserList(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.LoginStore.GetAllUserDetail()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Get User List Successfully", "")
}
