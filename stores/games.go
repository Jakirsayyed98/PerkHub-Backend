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

func (s *GamesStore) GetGames() (interface{}, error) {

	data, err := model.GetAllGames(s.db)

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	return data, nil
}

func (s *GamesStore) RefreshGames() (interface{}, error) {

	data, err := s.gameservice.GetAllgames()

	if err != nil {
		fmt.Println("Status :=", err.Error())
		return nil, err
	}
	resu := []model.GamesResponse{}
	for _, v := range data {

		gamedata, err := model.NewGamesResponse().FindGameByCode(s.db, v.Code)
		if err != nil {
			fmt.Println("already Exist Code 1:- ", err.Error())
			if err == sql.ErrNoRows {
				fmt.Println("already Exist Code 2:- ", err.Error())
				res := model.NewGamesResponse()
				if err := res.Bind(v); err != nil {
					fmt.Println("Error:= ", err.Error())
				}

				// if err == nil {
				resu = append(resu, *res)
				if err := res.InsertGames(s.db, res, "testID"); err != nil {
					fmt.Println("Error:= ", err.Error())
				}
				// }
			}
		}

		if err == nil {
			fmt.Println("already Exist Code :- ", gamedata.Code)
		}

	}

	return resu, nil
}
