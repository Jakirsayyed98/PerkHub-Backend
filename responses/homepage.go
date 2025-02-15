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
func (res *HomePageResponse) Bind(category []*model.Category, banner []*model.BannerCategory,
	popular []model.MiniApp, trending []model.MiniApp, top_cashback []model.MiniApp, finalCategory []CategoryResponse) (*HomePageResponse, error) {
	// Initialize the slices to empty arrays if not already initialized
	if category != nil {
		res.Categories = category
	} else {
		res.Categories = []*model.Category{}
	}

	// Assuming banner[0] exists, make sure you check it to prevent panics.
	if len(banner) > 0 && banner[0] != nil {
		res.Banner1 = banner[0].Banner
	} else {
		res.Banner1 = []*model.Banner{}
	}

	// If the rest of the banners (banner2, banner3, ...) are missing, initialize them as empty slices.
	if len(banner) > 1 && banner[1] != nil {
		res.Banner2 = banner[1].Banner
	} else {
		res.Banner2 = []*model.Banner{}
	}
	if len(banner) > 2 && banner[2] != nil {
		res.Banner3 = banner[2].Banner
	} else {
		res.Banner3 = []*model.Banner{}
	}
	if len(banner) > 3 && banner[3] != nil {
		res.Banner4 = banner[3].Banner
	} else {
		res.Banner4 = []*model.Banner{}
	}
	if len(banner) > 4 && banner[4] != nil {
		res.Banner5 = banner[4].Banner
	} else {
		res.Banner5 = []*model.Banner{}
	}
	if len(banner) > 5 && banner[5] != nil {
		res.Banner6 = banner[5].Banner
	} else {
		res.Banner6 = []*model.Banner{}
	}
	if len(banner) > 6 && banner[6] != nil {
		res.Banner7 = banner[6].Banner
	} else {
		res.Banner7 = []*model.Banner{}
	}
	if len(banner) > 7 && banner[7] != nil {
		res.Banner8 = banner[7].Banner
	} else {
		res.Banner8 = []*model.Banner{}
	}
	if len(banner) > 8 && banner[8] != nil {
		res.Banner9 = banner[8].Banner
	} else {
		res.Banner9 = []*model.Banner{}
	}
	if len(banner) > 9 && banner[9] != nil {
		res.Banner10s = banner[9].Banner
	} else {
		res.Banner10s = []*model.Banner{}
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
