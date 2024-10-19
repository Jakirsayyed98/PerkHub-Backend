package response

import (
	"PerkHub/model"
)

type HomePageResponse struct {
	Banner1      []*model.Banner    `json:"banner1"`
	Banner2      []*model.Banner    `json:"banner2"`
	Banner3      []*model.Banner    `json:"banner3"`
	Categories   []*model.Category  `json:"categories"`
	Popular      []model.MiniApp    `json:"popular"`
	Top_cashback []model.MiniApp    `json:"top_cashback"`
	Trending     []model.MiniApp    `json:"trending"`
	CategoryList []CategoryResponse `json:"category_list"`
}

func NewHomePagedata() *HomePageResponse {
	return &HomePageResponse{}
}

func (res *HomePageResponse) Bind(category []*model.Category, banner1 []*model.Banner, banner2 []*model.Banner, banner3 []*model.Banner,
	popular []model.MiniApp, trending []model.MiniApp, top_cashback []model.MiniApp, finalCategory []CategoryResponse) (interface{}, error) {
	var rest []HomePageResponse
	newResponse := HomePageResponse{
		Categories:   category,
		Banner1:      banner1,
		Banner2:      banner2,
		Banner3:      banner3,
		Popular:      popular,
		Trending:     trending,
		Top_cashback: top_cashback,
		CategoryList: finalCategory,
	}

	rest = append(rest, newResponse)
	return rest, nil
}
