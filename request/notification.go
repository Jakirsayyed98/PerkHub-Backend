package request

import (
	"PerkHub/connection"
	"PerkHub/utils"
	"fmt"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

type NotificationRequest struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Message     string `json:"message"`
	Image       string `json:"image"`
	EventType   string `json:"event_type"`
	StartDate   string `json:"start_date"`
	EndDate     string `json:"end_date"`
	Frequency   string `json:"frequency"`
	ClickAction string `json:"click_action"`
	Status      bool   `json:"status"`
}

func NewNotificationRequest() *NotificationRequest {
	return &NotificationRequest{}
}

func (req *NotificationRequest) Bind(c *gin.Context, awsInstance *connection.Aws) error {
	if !strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
		return fmt.Errorf("content type not supported %s", c.Request.Header.Get("Content-Type"))
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}

	image, _ := utils.UploadFileOnServer(form.File["image"], awsInstance)

	req.Id = c.PostForm("id")
	req.Title = c.PostForm("title")
	req.Message = c.PostForm("message")
	req.EventType = c.PostForm("event_type")
	req.StartDate = c.PostForm("start_date")
	req.EndDate = c.PostForm("end_date")
	req.Frequency = c.PostForm("frequency")
	req.ClickAction = c.PostForm("click_action")

	statusStr := c.PostForm("status")
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		return fmt.Errorf("invalid value for popular: %s", statusStr)
	}
	req.Status = status
	req.Image = image

	return nil
}
