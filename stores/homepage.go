package stores

import (
	"PerkHub/model"
	"PerkHub/responses"
	"database/sql"
)

type HomePageStore struct {
	db *sql.DB
}

func NewHomePageStore(dbs *sql.DB) *HomePageStore {
	return &HomePageStore{
		db: dbs,
	}
}

func (s *HomePageStore) GetHomePagedata() (*responses.HomePageResponse, error) {
	bannerCategory, err := model.GetAllBannersCategory(s.db)
	if err != nil {
		return nil, err
	}

	for _, v := range bannerCategory {
		banner, err := model.GetBannersByCategoryID(s.db, v.ID)
		if err != nil {
			return nil, err
		}
		v.Banner = banner

	}

	// banner1, err := model.GetBannerbyId(s.db, "1")
	// if err != nil {
	// 	return nil, err
	// }

	// banner2, err := model.GetBannerbyId(s.db, "2")
	// if err != nil {
	// 	return nil, err
	// }

	// banner3, err := model.GetBannerbyId(s.db, "3")
	// if err != nil {
	// 	return nil, err
	// }

	category, err := model.GetAllCategory(s.db)
	if err != nil {
		return nil, err
	}

	categoryHomePage, err := model.GetAllHomePageActive(s.db)
	if err != nil {
		return nil, err
	}
	categoriesres := responses.NewCategoryRes()
	categories, err := categoriesres.BindMultipleUsers(categoryHomePage)
	if err != nil {
		return nil, err
	}

	finalCategory := []responses.CategoryResponse{}
	for _, categorys := range categories {
		miniApps, err := model.GetMiniAppsByCategoryID(s.db, categorys.ID)
		if err != nil {
			return nil, err
		}

		if miniApps != nil {
			categorys.Data = miniApps
		} else {
			categorys.Data = []model.MiniApp{}
		}

		finalCategory = append(finalCategory, categorys)
	}

	popular, err := model.GetMiniAppsPopular(s.db)
	if err != nil {
		return nil, err
	}

	trending, err := model.GetMiniAppsTrending(s.db)
	if err != nil {
		return nil, err
	}

	topcashback, err := model.GetMiniAppsTopCashback(s.db)
	if err != nil {
		return nil, err
	}

	res := responses.NewHomePagedata()

	data, err := res.Bind(category, bannerCategory, popular, trending, topcashback, finalCategory)
	if err != nil {
		return nil, err
	}
	return data, nil
}
