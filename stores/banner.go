package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"database/sql"
)

type BannerStore struct {
	db *sql.DB
}

func NewBannerStore(dbs *sql.DB) *BannerStore {
	return &BannerStore{
		db: dbs,
	}
}

func (s *BannerStore) SaveBanner(req *request.Banner) (interface{}, error) {

	err := model.InsertBanner(s.db, req)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *BannerStore) UpdateBanner(req *request.Banner) (interface{}, error) {

	err := model.UpdateBanner(s.db, req)

	if err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *BannerStore) DeleteBanner(id string) (interface{}, error) {
	err := model.DeleteBanner(s.db, id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *BannerStore) GetAllBanners() (interface{}, error) {

	data, err := model.GetAllBanners(s.db)
	if err != nil {
		return nil, err
	}

	// res := response.NewBanner()

	// result, err := res.BindMultipleUsers(data)
	// if err != nil {
	// 	return nil, err
	// }

	return data, nil
}

// func (s *BannerStore) GetBannerbyId(id string) (interface{}, error) {

// 	data, err := model.GetAllBanners(s.db)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := response.NewBanner()

// 	result, err := res.BindMultipleUsers(data)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return result, nil

// }
