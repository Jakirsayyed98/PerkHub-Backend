package games

import (
	"PerkHub/model"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func GameByCategory(c *gin.Context) {

	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	req := model.NewGameCategory()

	if err := c.ShouldBind(&req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	result, err := store.GamesStore.GetGamesByCategory(req.Id.String())
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Game Get Successfully", "")
}

func GameSearch(c *gin.Context) {

	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	req := model.NewGameSearch()

	if err := c.ShouldBind(&req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.GameSearch(req.Name)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Game Search Successfully", "")
}

func SetGameStatus(c *gin.Context) {

	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	req := model.NewSetGameStatus()

	if err := c.ShouldBind(&req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	if err := store.GamesStore.SetGameStatus(req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, nil, "Status Updated Successfully", "")
}
func SetGameCategoryStatus(c *gin.Context) {

	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	req := model.NewSetGameStatus()

	if err := c.ShouldBind(&req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	if err := store.GamesStore.SetGameCategoryStatus(req); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, nil, "Status Updated Successfully", "")
}
