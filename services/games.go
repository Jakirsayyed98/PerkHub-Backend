package services

import (
	"PerkHub/responses"
	"PerkHub/settings"
	"errors"
	"fmt"
	"io"
	"net/http"
)

type GamesService struct {
	service *settings.HttpService
}

func NewGameService() *GamesService {
	return &GamesService{
		service: settings.NewHttpService("https://pub.gamezop.com/v3/games?id=4625"),
	}
}

func (s *GamesService) GetAllgames() (interface{}, error) {

	response, err := s.service.Get("", nil)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, errors.New("unkown Error")
	}

	res := responses.NewGameResponse()

	result, err := res.Unmarshal(body)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return result, nil
}
