package affiliates

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"

	"github.com/gin-gonic/gin"
)

func CueLinkCallBack(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	request := request.NewCueLinkCallBackRequest()

	if err := request.Bind(c); err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.AffiliatesStore.CueLinkCallBack(request)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	settings.StatusOk(c, result, "Callback get Successfully", "")
}
