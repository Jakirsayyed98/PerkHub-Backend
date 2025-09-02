package responses

import (
	"PerkHub/model"
	"time"
)

type BannerResponses struct {
	ID               string    `json:"id"`                 // Assuming ID is a string (could also be UUID)
	Name             string    `json:"name"`               // Name of the item
	BannerCategoryId string    `json:"banner_category_id"` // Name of the item
	Image            string    `json:"image"`              // URL or path to the item's image
	Url              string    `json:"url"`                // URL or path to the item's url
	Status           bool      `json:"status"`             // Status of the item (e.g., active, inactive)
	StartDate        string    `json:"start_date"`         // Start date of the item
	EndDate          string    `json:"end_date"`           // End date of the item
	CreatedAt        time.Time `json:"created_at"`         // Timestamp when the item was created
	UpdatedAt        time.Time `json:"updated_at"`         // Timestamp when the item was last updated
}

func NewBanner() *BannerResponses {
	return &BannerResponses{}
}

func (u *BannerResponses) ResponsesBind(banner *model.Banner) error {
	u.ID = banner.ID
	u.Name = banner.Name
	u.BannerCategoryId = banner.BannerCategoryId
	u.Image = banner.Image
	u.Url = banner.Url
	u.Status = banner.Status
	u.StartDate = banner.StartDate
	u.EndDate = banner.EndDate
	u.CreatedAt = banner.CreatedAt
	u.UpdatedAt = banner.UpdatedAt

	return nil
}
