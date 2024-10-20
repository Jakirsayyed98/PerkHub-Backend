package stores

import (
	"PerkHub/services"
	"database/sql"
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

	data, err := s.gameservice.GetAllgames()

	if err != nil {
		return nil, err
	}

	return data, nil
}
