package stores

import (
	"PerkHub/model"
	"PerkHub/services"
	"database/sql"
	"fmt"
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

	data, err := model.GetGameCategories(s.db)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (s *GamesStore) GetGames() (interface{}, error) {

	data, err := model.GetAllGames(s.db)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (s *GamesStore) GetPopularGames() (interface{}, error) {

	data, err := model.GetPopularGames(s.db)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (s *GamesStore) GetTrendingGames() (interface{}, error) {

	data, err := model.GetTrendingGames(s.db)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (s *GamesStore) GetGamesByCategory(category_id string) (interface{}, error) {

	data, err := model.GetAllGamesBycategory(s.db, category_id)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (s *GamesStore) GameSearch(search string) (interface{}, error) {

	data, err := model.GetGameSearch(s.db, search)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (s *GamesStore) Refreshcategory() error {
	data, err := s.gameservice.GetAllgames()

	if err != nil {
		return err
	}
	for _, v := range data {
		for _, cat := range v.Categories.EN {
			_, err = model.NewGameCategory().FindGameCategoryByNameOrId(s.db, "", cat)

			if err != nil {
				fmt.Println(err.Error())
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

	data, err := s.gameservice.GetAllgames()

	if err != nil {
		fmt.Println("Status :=", err.Error())
		return nil, err
	}
	resu := []model.GamesResponse{}
	for _, v := range data {
		categorydata, err := model.NewGameCategory().FindGameCategoryByNameOrId(s.db, "", v.Categories.EN[0])
		gamedata, err := model.NewGamesResponse().FindGameByCode(s.db, v.Code)
		if err != nil {
			if err == sql.ErrNoRows {
				res := model.NewGamesResponse()

				if err := res.Bind(v, categorydata.Id.String()); err != nil {
					fmt.Println("Error:= ", err.Error())
				}
				resu = append(resu, *res)
				if err := res.InsertGames(s.db, res, "testID"); err != nil {
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
