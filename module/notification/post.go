package notification

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"

	"github.com/gin-gonic/gin"
)

func CreateNotification(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewNotificationRequest()

	if err := request.Bind(c, awsInstance); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	err = store.NotificationStore.CreateNotification(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, nil, "Notification created successfully", "")
}

func UpdateNotification(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	awsInstance, err := stores.GetAwsInstance(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewNotificationRequest()

	if err := request.Bind(c, awsInstance); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	err = store.NotificationStore.UpdateNotification(request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, nil, "Notification updated successfully", "")
}
