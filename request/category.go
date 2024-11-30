package request

import (
	"PerkHub/connection"
	"PerkHub/constants"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type Category struct {
	ID              string    `json:"id"`               // Assuming ID is a string (could also be UUID)
	Name            string    `json:"name"`             // Name of the item
	Description     string    `json:"description"`      // Description of the item
	Image           string    `json:"image"`            // URL or path to the item's image
	Status          string    `json:"status"`           // Status of the item (e.g., active, inactive)
	HomepageVisible bool      `json:"homepage_visible"` // Visibility on the homepage
	CreatedAt       time.Time `json:"created_at"`       // Timestamp when the item was created
	UpdatedAt       time.Time `json:"updated_at"`       // Timestamp when the item was last updated
}

func NewCategory() *Category {
	return &Category{}
}

func (category *Category) Bind(c *gin.Context, awsInstance *connection.Aws) error {
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

		fileHeader := files[0]
		f, err := fileHeader.Open()
		if err != nil {
			return err
		}
		defer f.Close()

		image, err = awsInstance.UploadFile(f, fileHeader.Filename, constants.AWSBucketName, constants.AWSSecretAccessKey, constants.AWSCloudFrontURL)
		// image, err = utils.SaveFile(c, file)
		if err != nil {
			return err
		}
	}

	category.ID = c.PostForm("id")
	category.Name = c.PostForm("name")
	category.Description = c.PostForm("description")
	category.Status = c.PostForm("status")
	category.HomepageVisible = false
	category.Image = image
	category.CreatedAt = time.Now()
	category.UpdatedAt = time.Now()

	return nil
}

type CategoryID struct {
	CategoryId string `json:"category_id"`
}

func NewCategoryID() *CategoryID {
	return &CategoryID{}
}
