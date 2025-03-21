package responses

import (
	"PerkHub/model"
	"fmt"
	"time"
)

type CategoryResponse struct {
	ID              string          `json:"id"`
	Name            string          `json:"name"`
	Description     string          `json:"description"`
	Image           string          `json:"image"`
	Status          string          `json:"status"`
	HomepageVisible bool            `json:"homepage_visible"`
	CreatedAt       time.Time       `json:"created_at"`
	UpdatedAt       time.Time       `json:"updated_at"`
	Data            []model.MiniApp `json:"data"`
}

func NewCategoryRes() CategoryResponse {
	return CategoryResponse{}
}

func (u *CategoryResponse) BindMultipleUsers(categories []*model.Category) ([]CategoryResponse, error) {
	var responses []CategoryResponse

	for _, category := range categories {
		var response CategoryResponse
		err := response.ResponsesBind(category)
		if err != nil {
			return nil, fmt.Errorf("error binding user detail: %w", err)
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (u *CategoryResponse) ResponsesBind(cat *model.Category) error {
	u.ID = cat.ID
	u.Name = cat.Name
	u.Description = cat.Description
	u.Image = cat.Image
	u.Status = cat.Status
	u.HomepageVisible = cat.HomepageVisible
	u.CreatedAt = cat.CreatedAt
	u.UpdatedAt = cat.UpdatedAt

	return nil
}
