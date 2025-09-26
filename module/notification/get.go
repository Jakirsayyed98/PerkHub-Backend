package notification

import (
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func GetAdminNotificationList(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.NotificationStore.GetAdminNotificationList()
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Successfully get notification list", "")
}

func SendNotification(c *gin.Context) {
	id := c.Param("id")
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.NotificationStore.AdminSendNotification(id)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Notification send successfullly", "")
}

func GetNotificationById(c *gin.Context) {
	id := c.Param("id")
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.NotificationStore.GetNotificationById(id)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Notification got successfullly", "")
}

func GetUserNotificationHistory(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	userId := c.MustGet("user_id")
	result, err := store.NotificationStore.GetUserNotificationHistory(userId.(string))
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Successfully get notification list", "")
}
