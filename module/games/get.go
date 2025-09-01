package games

import (
	"PerkHub/settings"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func GetGameCategories(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.GetGameCategories()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Successfully get games categories", "")
}

func GetGames(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.GetGames()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Successfully get games", "")
}

func GetPopulargames(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.GetPopularGames()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Successfully get games", "")
}

func GetTrendingGames(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.GetTrendingGames()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Successfully get games", "")
}

func RefreshGameCategories(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	err = store.GamesStore.Refreshcategory()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, nil, "Games Category refresh successfully", "")
}
func RefreshGames(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.RefreshGames()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Games refresh successfully", "")
}

func AdminGetGameCategories(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.GamesStore.GetAdminGameCategories()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, result, "Successfully get games categories", "")
}

func AdminGetGames(c *gin.Context) {
	store, err := stores.GetStores(c)

	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	result, err := store.GamesStore.GetAdminGames()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Successfully get games", "")
}
