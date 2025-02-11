package request

import (
	"PerkHub/connection"
	"PerkHub/utils"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type BannerCategory struct {
	Title string `json:"title"` // Name of the item
}

func NewBannerCategory() *BannerCategory {
	return &BannerCategory{}
}

type BannerRequest struct {
	CategoryId string `json:"category_id"` // Name of the item
}

func NewBannerRequest() *BannerRequest {
	return &BannerRequest{}
}

type Banner struct {
	ID               string    `json:"id"`                 // Assuming ID is a string (could also be UUID)
	Name             string    `json:"name"`               // Name of the item
	BannerCategoryId string    `json:"banner_category_id"` // Name of the item
	Image            string    `json:"image"`              // URL or path to the item's image
	Url              string    `json:"url"`                // URL or path to the item's url
	StartDate        string    `json:"start_date"`         // URL or path to the item's url
	EndDate          string    `json:"end_date"`           // URL or path to the item's url
	Status           bool      `json:"status"`             // Status of the item (e.g., active, inactive)
	CreatedAt        time.Time `json:"created_at"`         // Timestamp when the item was created
	UpdatedAt        time.Time `json:"updated_at"`         // Timestamp when the item was last updated
}

func NewBanner() *Banner {
	return &Banner{}
}

func (banner *Banner) Bind(c *gin.Context, awsInstance *connection.Aws) error {
	if !strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
		return fmt.Errorf("content type not supported %s", c.Request.Header.Get("Content-Type"))
	}
	fmt.Println(c.PostForm("id"))

	if len(strings.Split(c.PostForm("image"), "http")) > 0 {
		fmt.Println("Update")
	} else {
		fmt.Println("New")
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}

		image, _ := utils.UploadFileOnServer(form.File["image"], awsInstance)
		banner.Image = image
	}

	banner.ID = c.PostForm("id")
	banner.Name = c.PostForm("name")
	banner.BannerCategoryId = c.PostForm("banner_category_id")
	banner.StartDate = c.PostForm("start_date")
	banner.EndDate = c.PostForm("end_date")

	banner.Url = c.PostForm("url")

	statusStr := c.PostForm("status")
	status, err := strconv.ParseBool(statusStr)
	if err != nil {
		return fmt.Errorf("invalid value for popular: %s", statusStr)
	}
	banner.Status = status

	return nil
}
