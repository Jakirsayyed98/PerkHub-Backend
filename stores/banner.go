package stores

import (
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"PerkHub/request"
	"PerkHub/responses"
	"database/sql"
	"time"
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
	startTime := time.Now()
	if req.ID == "" {
		err := model.InsertBanner(s.db, req)
		if err != nil {
			log := logger.LogData{
				Message:   err.Error(),
				StartTime: startTime,
			}
			logger.LogError(log)
			return nil, err
		}
	} else {
		err := model.UpdateBannerData(s.db, req)

		if err != nil {
			log := logger.LogData{
				Message:   err.Error(),
				StartTime: startTime,
			}
			logger.LogError(log)
			return nil, err
		}

	}

	return nil, nil
}

func (s *BannerStore) UpdateBanner(req *request.Banner) (interface{}, error) {
	startTime := time.Now()
	err := model.UpdateBannerData(s.db, req)

	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return nil, nil
}

func (s *BannerStore) DeleteBanner(id string) (interface{}, error) {
	startTime := time.Now()
	err := model.DeleteBanner(s.db, id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return nil, nil
}

func (s *BannerStore) GetBannersByCategoryID(categoryId string) (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetBannersByCategoryID(s.db, categoryId)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return data, nil
}

func (s *BannerStore) GetBannerbyId(id string) (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetBannerbyId(s.db, id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	res := responses.NewBanner()

	err = res.ResponsesBind(data)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return res, nil

}
