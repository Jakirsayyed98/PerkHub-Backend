package stores

import (
	"PerkHub/model"
	"PerkHub/responses"
	"database/sql"

	"golang.org/x/sync/errgroup"
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
	category, err := model.GetAllActiveCategories(s.db)
	if err != nil {
		return nil, err
	}
	// Get categories for homepage
	categoryHomePage, err := model.GetAllHomePageActive(s.db)
	if err != nil {
		return nil, err
	}

	banner1, err := model.GetBannersByCategoryID(s.db, "1")
	if err != nil {
		return nil, err
	}

	banner2, err := model.GetBannersByCategoryID(s.db, "1")
	if err != nil {
		return nil, err
	}
	banner3, err := model.GetBannersByCategoryID(s.db, "1")
	if err != nil {
		return nil, err
	}

	categoriesres := responses.NewCategoryRes()
	categories, err := categoriesres.BindMultipleUsers(categoryHomePage)
	if err != nil {
		return nil, err
	}

	// Attach mini apps to categories
	finalCategory := make([]responses.CategoryResponse, 0, len(categories))
	// for _, cat := range categories {
	// 	miniApps, err := model.GetStoresByCategory(s.db, cat.ID)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if miniApps != nil {
	// 		if len(miniApps) > 6 {
	// 			cat.Data = miniApps[:6] // take only first 6 items
	// 		} else {
	// 			cat.Data = miniApps
	// 		}
	// 	} else {
	// 		cat.Data = []model.MiniApp{}
	// 	}
	// 	finalCategory = append(finalCategory, cat)
	// }

	// Run popular/trending/topcashback queries in parallel
	var (
		popular     []model.MiniApp
		trending    []model.MiniApp
		topcashback []model.MiniApp
	)

	g := new(errgroup.Group)

	g.Go(func() error {
		var err error
		popular, err = model.GetMiniAppsPopular(s.db)
		return err
	})

	g.Go(func() error {
		var err error
		trending, err = model.GetMiniAppsTrending(s.db)
		return err
	})

	g.Go(func() error {
		var err error
		topcashback, err = model.GetMiniAppsTopCashback(s.db)
		return err
	})

	// Wait for all goroutines
	if err := g.Wait(); err != nil {
		return nil, err
	}

	// Build final response
	res := responses.NewHomePagedata()
	data, err := res.Bind(category, banner1, banner2, banner3, popular, trending, topcashback, finalCategory)
	if err != nil {
		return nil, err
	}

	return data, nil
}
