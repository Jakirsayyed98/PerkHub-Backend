package mobile

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func UpdateNotificationToken(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	var request struct {
		Token string `json:"fcm_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	if err := store.LoginStore.UpdateNotificationToken(c.MustGet("user_id").(string), request.Token); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, nil, "Notification token updated successfully", "")
}
