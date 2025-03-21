package request

import (
	"PerkHub/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Banner struct {
	ID        string    `json:"id"`         // Assuming ID is a string (could also be UUID)
	Name      string    `json:"name"`       // Name of the item
	BannerId  string    `json:"banner_id"`  // Name of the item
	Image     string    `json:"image"`      // URL or path to the item's image
	Url       string    `json:"url"`        // URL or path to the item's url
	Status    string    `json:"status"`     // Status of the item (e.g., active, inactive)
	CreatedAt time.Time `json:"created_at"` // Timestamp when the item was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the item was last updated
}

func NewBanner() *Banner {
	return &Banner{}
}

func (banner *Banner) Bind(c *gin.Context) error {
	if !strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
		return fmt.Errorf("content type not supported %s", c.Request.Header.Get("Content-Type"))
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	image := ""
	files := form.File["image"]
	if len(files) > 0 {

		file := files[0]

		image, err = utils.SaveFile(c, file)
		if err != nil {
			return err
		}
	}

	banner.ID = c.PostForm("id")
	banner.Name = c.PostForm("name")
	banner.BannerId = c.PostForm("banner_id")
	banner.Image = image
	banner.Url = c.PostForm("url")
	banner.Status = c.PostForm("status")

	return nil
}
