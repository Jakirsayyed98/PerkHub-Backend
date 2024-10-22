package games

import (
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func GetGames(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.GetGames()
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Successfully get games", "")
}

func RefreshGames(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.RefreshGames()
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Games refresh successfully", "")
}
