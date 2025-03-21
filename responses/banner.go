package responses

import (
	"PerkHub/model"
	"fmt"
	"time"
)

type BannerResponses struct {
	ID        string    `json:"id"`         // Assuming ID is a string (could also be UUID)
	Name      string    `json:"name"`       // Name of the item
	BannerId  string    `json:"banner_id"`  // Name of the item
	Image     string    `json:"image"`      // URL or path to the item's image
	Url       string    `json:"url"`        // URL or path to the item's url
	Status    bool      `json:"status"`     // Status of the item (e.g., active, inactive)
	CreatedAt time.Time `json:"created_at"` // Timestamp when the item was created
	UpdatedAt time.Time `json:"updated_at"` // Timestamp when the item was last updated
}

func NewBanner() *BannerResponses {
	return &BannerResponses{}
}

func (u *BannerResponses) BindMultipleUsers(banners []*model.Banner) ([]BannerResponses, error) {
	var responses []BannerResponses

	for _, banner := range banners {
		var response BannerResponses
		err := response.ResponsesBind(banner)
		if err != nil {
			return nil, fmt.Errorf("error binding user detail: %w", err)
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (u *BannerResponses) ResponsesBind(banner *model.Banner) error {
	u.ID = banner.ID
	u.Name = banner.Name
	u.BannerId = banner.BannerCategoryId
	u.Image = banner.Image
	u.Url = banner.Url
	u.Status = banner.Status
	u.CreatedAt = banner.CreatedAt
	u.UpdatedAt = banner.UpdatedAt

	return nil
}
