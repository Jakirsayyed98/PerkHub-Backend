package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"PerkHub/responses"
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
	if req.ID == "" {
		err := model.InsertBanner(s.db, req)
		if err != nil {
			return nil, err
		}
	} else {
		err := model.UpdateBannerData(s.db, req)

		if err != nil {
			return nil, err
		}

	}

	return nil, nil
}

func (s *BannerStore) UpdateBanner(req *request.Banner) (interface{}, error) {

	err := model.UpdateBannerData(s.db, req)

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

func (s *BannerStore) GetBannersByCategoryID(categoryId string) (interface{}, error) {

	data, err := model.GetBannersByCategoryID(s.db, categoryId)
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

func (s *BannerStore) GetBannerbyId(id string) (interface{}, error) {

	data, err := model.GetBannerbyId(s.db, id)
	if err != nil {
		return nil, err
	}

	res := responses.NewBanner()

	err = res.ResponsesBind(data)
	if err != nil {
		return nil, err
	}

	return res, nil

}
