package stores

import (
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"PerkHub/request"
	"PerkHub/services"
	"database/sql"
	"time"
)

type CategoryStore struct {
	db             *sql.DB
	cueLinkService *services.CueLinkAffiliateService
}

func NewCategoryStore(dbs *sql.DB) *CategoryStore {
	cuelinkService := services.NewCueLinkAffiliateService()
	return &CategoryStore{
		db:             dbs,
		cueLinkService: cuelinkService,
	}
}

func (s *CategoryStore) SaveCategory(req *request.Category) (interface{}, error) {
	startTime := time.Now()
	if err := model.InsertCategory(s.db, req); err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return nil, nil
}

func (s *CategoryStore) UpdateCategory(req *request.Category) (interface{}, error) {
	startTime := time.Now()
	if err := model.UpdateCategory(s.db, req); err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return nil, nil
}

func (s *CategoryStore) DeleteCategory(id string) (interface{}, error) {
	startTime := time.Now()
	if err := model.DeleteCategoryByID(s.db, id); err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return nil, nil
}

func (s *CategoryStore) GetAllCategory() (interface{}, error) {
	startTime := time.Now()
	result, err := model.GetAllCategory(s.db)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}

	return result, nil
}

func (s *CategoryStore) GetCategoryByID(id string) (interface{}, error) {
	startTime := time.Now()
	result, err := model.GetCategoryByID(s.db, id)
	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return result, nil
}

func (s *CategoryStore) ActiveDeactiveCategory(req *request.CategoryActiveDeactive) (interface{}, error) {
	startTime := time.Now()
	if err := model.ActivateDeactiveCategorykey(s.db, req.Key, req.CategoryId, req.Value); err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	return nil, nil
}

func (s *CategoryStore) GetStoresCategoryRefresh() (interface{}, error) {
	startTime := time.Now()
	page := 1
	perPage := 100

	for {
		// Fetch campaigns for the current page
		data, err := s.cueLinkService.GetAllCampaigns(page, perPage)
		if err != nil {
			log := logger.LogData{
				Message:   err.Error(),
				StartTime: startTime,
			}
			logger.LogError(log)
			return nil, err
		}

		// Stop if no campaigns are returned
		if len(data.Campaigns) == 0 {
			break
		}

		for _, v := range data.Campaigns {
			for _, c := range v.Categories {
				isExist, err := model.CategoryExists(s.db, c.Name)
				if err != nil {
					log := logger.LogData{
						Message:   err.Error(),
						StartTime: startTime,
					}
					logger.LogError(log)
					return nil, err
				}
				if !isExist {
					err := model.InsertCategory(s.db, &request.Category{
						Name:        c.Name,
						Description: "",
						Image:       "",   // optional
						Status:      true, // optional if InsertCategory sets default
					})
					if err != nil {
						log := logger.LogData{
							Message:   err.Error(),
							StartTime: startTime,
						}
						logger.LogError(log)
						return nil, err
					}
				}
			}
		}
		// Move to the next page
		page++
	}
	return nil, nil
}
