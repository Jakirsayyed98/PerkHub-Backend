package responses

import (
	"PerkHub/model"
)

type HomePageResponse struct {
	Banner1      []*model.Banner    `json:"banner1"`
	Banner2      []*model.Banner    `json:"banner2"`
	Banner3      []*model.Banner    `json:"banner3"`
	Banner4      []*model.Banner    `json:"banner4"`
	Banner5      []*model.Banner    `json:"banner5"`
	Banner6      []*model.Banner    `json:"banner6"`
	Banner7      []*model.Banner    `json:"banner7"`
	Banner8      []*model.Banner    `json:"banner8"`
	Banner9      []*model.Banner    `json:"banner9"`
	Banner10s    []*model.Banner    `json:"banner10"`
	Categories   []*model.Category  `json:"categories"`
	Popular      []model.MiniApp    `json:"popular"`
	Top_cashback []model.MiniApp    `json:"top_cashback"`
	Trending     []model.MiniApp    `json:"trending"`
	CategoryList []CategoryResponse `json:"category_list"`
}

func NewHomePagedata() *HomePageResponse {
	return &HomePageResponse{}
}
func (res *HomePageResponse) Bind(category []*model.Category, banner1, banner2, banner3 []*model.Banner,
	popular []model.MiniApp, trending []model.MiniApp, top_cashback []model.MiniApp, finalCategory []CategoryResponse) (*HomePageResponse, error) {
	// Initialize the slices to empty arrays if not already initialized
	if category != nil {
		res.Categories = category
	} else {
		res.Categories = []*model.Category{}
	}

	if banner1 != nil {
		res.Banner1 = banner1
	} else {
		res.Banner1 = []*model.Banner{}
	}

	if banner2 != nil {
		res.Banner2 = banner2
	} else {
		res.Banner2 = []*model.Banner{}
	}

	if banner3 != nil {
		res.Banner3 = banner3
	} else {
		res.Banner3 = []*model.Banner{}
	}

	// Popular and trending apps can also be checked to prevent `nil` issues.
	if popular != nil {
		res.Popular = popular
	} else {
		res.Popular = []model.MiniApp{}
	}

	if trending != nil {
		res.Trending = trending
	} else {
		res.Trending = []model.MiniApp{}
	}

	if top_cashback != nil {
		res.Top_cashback = top_cashback
	} else {
		res.Top_cashback = []model.MiniApp{}
	}

	if finalCategory != nil {
		res.CategoryList = finalCategory
	} else {
		res.CategoryList = []CategoryResponse{}
	}

	// Return the populated response
	return res, nil
}
