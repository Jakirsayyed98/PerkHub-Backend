package services

import (
	"PerkHub/constants"
	"PerkHub/request"
	"PerkHub/settings"
	"errors"
	"io"
	"net/http"
)

type GamesService struct {
	service *settings.HttpService
}

func NewGameService() *GamesService {
	return &GamesService{
		service: settings.NewHttpService(constants.GET_GAME_BASE_URL),
	}
}

func (s *GamesService) GetAllgames() ([]request.GameResponse, error) {

	response, err := s.service.Get("", nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("unkown Error")
	}

	res := request.NewGameResponse()

	result, err := res.Unmarshal(body)
	if err != nil {
		return nil, err
	}
	return result, nil
}
