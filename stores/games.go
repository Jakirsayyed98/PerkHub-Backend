package stores

import (
	"PerkHub/model"
	"PerkHub/pkg/logger"
	"PerkHub/services"
	"database/sql"
	"fmt"
	"time"
)

type GamesStore struct {
	db          *sql.DB
	gameservice *services.GamesService
}

func NewGameStore(db *sql.DB) *GamesStore {
	gameservice := services.NewGameService()

	return &GamesStore{
		db:          db,
		gameservice: gameservice,
	}
}

func (s *GamesStore) GetGameCategories() (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetGameCategories(s.db)

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

func (s *GamesStore) GetGames() (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetAllGames(s.db)

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

func (s *GamesStore) GetPopularGames() (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetPopularGames(s.db)

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

func (s *GamesStore) GetTrendingGames() (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetTrendingGames(s.db)

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

func (s *GamesStore) GetGamesByCategory(category_id string) (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetAllGamesBycategory(s.db, category_id)

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

func (s *GamesStore) GameSearch(search string) (interface{}, error) {
	startTime := time.Now()
	data, err := model.GetGameSearch(s.db, search)

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

func (s *GamesStore) Refreshcategory() error {
	startTime := time.Now()
	data, err := s.gameservice.GetAllgames()

	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return err
	}
	for _, v := range data {
		for _, cat := range v.Categories.EN {
			_, err = model.NewGameCategory().FindGameCategoryByNameOrId(s.db, "", cat)

			if err != nil {
				log := logger.LogData{
					Message:   err.Error(),
					StartTime: startTime,
				}
				logger.LogError(log)
				if err == sql.ErrNoRows {
					model.InsertGamesCategory(s.db, cat, "")
				}
			}

			if err == nil {
				fmt.Println("already Exist Category :- ", cat)
			}
		}

	}

	return nil
}

func (s *GamesStore) RefreshGames() (interface{}, error) {
	startTime := time.Now()
	data, err := s.gameservice.GetAllgames()

	if err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return nil, err
	}
	resu := []model.GamesResponse{}
	for _, v := range data {
		categorydata, err := model.NewGameCategory().FindGameCategoryByNameOrId(s.db, "", v.Categories.EN[0])
		gamedata, err := model.NewGamesResponse().FindGameByCode(s.db, v.Code)
		if err != nil {
			log := logger.LogData{
				Message:   err.Error(),
				StartTime: startTime,
			}
			logger.LogError(log)
			if err == sql.ErrNoRows {
				res := model.NewGamesResponse()

				if err := res.Bind(v, categorydata.Id.String()); err != nil {
					fmt.Println("Error:= ", err.Error())
				}
				resu = append(resu, *res)
				if err := res.InsertGames(s.db, res); err != nil {
					fmt.Println("Error:= ", err.Error())
				}
			}
		}

		if err == nil {
			fmt.Println("already Exist Code :- ", gamedata.Code)
		}

	}

	return resu, nil
}

func (s *GamesStore) SetGameStatus(game *model.SetGameStatus) error {
	startTime := time.Now()
	if err := model.UpdateGameStatus(s.db, string(game.StatusType), game.Id, game.Status); err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return err
	}

	return nil
}

func (s *GamesStore) SetGameCategoryStatus(game *model.SetGameStatus) error {
	startTime := time.Now()
	if err := model.ActivateDeactiveGameCategoryKey(s.db, game.Id, game.Status); err != nil {
		log := logger.LogData{
			Message:   err.Error(),
			StartTime: startTime,
		}
		logger.LogError(log)
		return err
	}

	return nil
}

func (s *GamesStore) GetAdminGameCategories() (interface{}, error) {
	startTime := time.Now()

	result, err := model.AdminGetGameCategories(s.db)
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

func (s *GamesStore) GetAdminGames() (interface{}, error) {
	startTime := time.Now()
	result, err := model.AdminGetAllGames(s.db)
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
