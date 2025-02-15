package games

import (
	"PerkHub/model"
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func GameByCategory(c *gin.Context) {

	store, err := stores.GetStores(c)

	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	req := model.NewGameCategory()

	if err := c.ShouldBind(&req); err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}
	result, err := store.GamesStore.GetGamesByCategory(req.Id.String())
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	settings.StatusOk(c, result, "Game Get Successfully", "")
}

func GameSearch(c *gin.Context) {

	store, err := stores.GetStores(c)

	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	req := model.NewGameSearch()

	if err := c.ShouldBind(&req); err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	result, err := store.GamesStore.GameSearch(req.Name)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	settings.StatusOk(c, result, "Game Get Successfully", "")
}

func SetGameStatus(c *gin.Context) {

	store, err := stores.GetStores(c)

	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	req := model.NewSetGameStatus()

	if err := c.ShouldBind(&req); err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	if err := store.GamesStore.SetGameStatus(req); err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	settings.StatusOk(c, nil, "Game Get Successfully", "")
}
