package model

import (
	"PerkHub/request"
	"database/sql"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
)

type GamesResponse struct {
	Id               uuid.UUID `json:"id"`
	Code             string    `json:"code"`
	URL              string    `json:"url"`
	Name             string    `json:"name"`
	IsPortrait       bool      `json:"isPortrait"`
	Description      string    `json:"description"`
	GamePreviews     string    `json:"gamePreviews"`
	Assets           Assets    `json:"assets"`
	Categories       string    `json:"category_id"`
	ColorsM          string    `json:"colorMuted"`
	ColorsV          string    `json:"colorVibrant"`
	PrivateAllowed   bool      `json:"privateAllowed"`
	Rating           string    `json:"rating"`
	Status           bool      `json:"status"`
	NumberOfRatings  string    `json:"numberOfRatings"`
	GamePlays        string    `json:"gamePlays"`
	HasIntegratedAds bool      `json:"hasIntegratedAds"`
	Width            string    `json:"width"`
	Height           string    `json:"height"`
	Trending         bool      `json:"trending"`
	Popular          bool      `json:"popular"`
}

type Assets struct {
	Cover     string   `json:"cover"`
	Brick     string   `json:"brick"`
	Thumb     string   `json:"thumb"`
	Wall      string   `json:"wall"`
	Square    string   `json:"square"`
	Screens   []string `json:"screens"`
	CoverTiny string   `json:"coverTiny"`
	BrickTiny string   `json:"brickTiny"`
}

func NewGamesResponse() *GamesResponse {
	return &GamesResponse{}
}

type GameSearch struct {
	Name string `json:"name"`
}

func NewGameSearch() *GameSearch {
	return &GameSearch{}
}

type StatusType string

const (
	StatusTypeStatus   StatusType = "status"
	StatusTypePopular  StatusType = "popular"
	StatusTypeTrending StatusType = "trending"
)

type SetGameStatus struct {
	Id         string     `json:"id"`
	Status     bool       `json:"status"`
	StatusType StatusType `json:"key"`
}

func NewSetGameStatus() *SetGameStatus {
	return &SetGameStatus{}
}

func (s *GamesResponse) Bind(data request.GameResponse, categoryId string) error {
	s.Code = data.Code
	s.URL = data.URL
	s.Name = data.Name.EN
	s.IsPortrait = data.IsPortrait
	s.Description = data.Description.EN
	s.GamePreviews = data.GamePreviews.EN
	s.Assets = Assets(data.Assets)
	s.ColorsM = data.ColorsM
	s.ColorsV = data.Colors
	s.Categories = categoryId
	s.PrivateAllowed = data.PrivateAllowed
	s.Rating = fmt.Sprintf("%.2f", data.Rating)
	s.NumberOfRatings = fmt.Sprintf("%d", data.NumberOfRatings)
	s.GamePlays = fmt.Sprintf("%d", data.GamePlays)
	s.HasIntegratedAds = data.HasIntegratedAds
	s.Width = fmt.Sprintf("%d", data.Width)
	s.Height = fmt.Sprintf("%d", data.Height)
	return nil
}

func (s *GamesResponse) InsertGames(db *sql.DB, item *GamesResponse, category_id string) error {
	datas, err := json.Marshal(item.Assets)
	query := `
		INSERT INTO games_data (code, url, name, isPortrait, description, gamePreviews, assets, category_id, colorMuted, colorVibrant, status, privateAllowed, rating, numberOfRatings, gamePlays, hasIntegratedAds, width, height, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, now(), now())
	`
	_, err = db.Exec(query,
		item.Code,             // $1
		item.URL,              // $2
		item.Name,             // $3
		item.IsPortrait,       // $4 (assuming this is a boolean)
		item.Description,      // $5
		item.GamePreviews,     // $6
		datas,                 // $7
		item.Categories,       // $8
		item.ColorsM,          // $9
		item.ColorsV,          // $10
		true,                  // $11 (status)
		item.PrivateAllowed,   // $12
		item.Rating,           // $13
		item.NumberOfRatings,  // $14
		item.GamePlays,        // $15
		item.HasIntegratedAds, // $16
		item.Width,            // $17
		item.Height,           // $18
	)
	return err
}

func (s *GamesResponse) FindGameByCode(db *sql.DB, code string) (*GamesResponse, error) {
	query := `SELECT id, code, url, name, assets, isPortrait, description, gamepreviews, category_id, colorMuted, colorVibrant, status, privateAllowed, rating, numberOfRatings, gamePlays, hasIntegratedAds, width, height FROM games_data WHERE code = $1;`
	row := db.QueryRow(query, code)

	var game GamesResponse
	var assetsJSON string

	err := row.Scan(
		&game.Id,
		&game.Code,
		&game.URL,
		&game.Name,
		&assetsJSON,
		&game.IsPortrait,
		&game.Description,
		&game.GamePreviews,
		&game.Categories,
		&game.ColorsM,
		&game.ColorsV,
		&game.Status,
		&game.PrivateAllowed,
		&game.Rating,
		&game.NumberOfRatings,
		&game.GamePlays,
		&game.HasIntegratedAds,
		&game.Width,
		&game.Height,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, err
		}
		return nil, err
	}

	err = json.Unmarshal([]byte(assetsJSON), &game.Assets)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling assets JSON: %w", err)
	}

	return &game, nil
}

func GetAllGames(db *sql.DB) ([]GamesResponse, error) {
	query := `SELECT id, code, url, name, isPortrait, description, gamepreviews, assets, category_id, colorMuted, colorVibrant, status, privateAllowed, rating, numberOfRatings, gamePlays, hasIntegratedAds, width, height,trending,popular FROM games_data WHERE status=true`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var games []GamesResponse

	for rows.Next() {
		var game GamesResponse
		var assetsJSON string
		err := rows.Scan(
			&game.Id,
			&game.Code,
			&game.URL,
			&game.Name,
			&game.IsPortrait,
			&game.Description,
			&game.GamePreviews,
			&assetsJSON,
			&game.Categories,
			&game.ColorsM,
			&game.ColorsV,
			&game.Status,
			&game.PrivateAllowed,
			&game.Rating,
			&game.NumberOfRatings,
			&game.GamePlays,
			&game.HasIntegratedAds,
			&game.Width,
			&game.Height,
			&game.Trending,
			&game.Popular,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		err = json.Unmarshal([]byte(assetsJSON), &game.Assets)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return games, nil
}

func GetPopularGames(db *sql.DB) ([]GamesResponse, error) {
	query := `SELECT id, code, url, name, isPortrait, description, gamepreviews, assets, category_id, colorMuted, colorVibrant, status, privateAllowed, rating, numberOfRatings, gamePlays, hasIntegratedAds, width, height,trending,popular FROM games_data WHERE status=true AND popular=true`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var games []GamesResponse

	for rows.Next() {
		var game GamesResponse
		var assetsJSON string
		err := rows.Scan(
			&game.Id,
			&game.Code,
			&game.URL,
			&game.Name,
			&game.IsPortrait,
			&game.Description,
			&game.GamePreviews,
			&assetsJSON,
			&game.Categories,
			&game.ColorsM,
			&game.ColorsV,
			&game.Status,
			&game.PrivateAllowed,
			&game.Rating,
			&game.NumberOfRatings,
			&game.GamePlays,
			&game.HasIntegratedAds,
			&game.Width,
			&game.Height,
			&game.Trending,
			&game.Popular,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		err = json.Unmarshal([]byte(assetsJSON), &game.Assets)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return games, nil
}

func GetTrendingGames(db *sql.DB) ([]GamesResponse, error) {
	query := `SELECT id, code, url, name, isPortrait, description, gamepreviews, assets, category_id, colorMuted, colorVibrant, status, privateAllowed, rating, numberOfRatings, gamePlays, hasIntegratedAds, width, height,trending,popular FROM games_data WHERE status=true AND trending=true`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var games []GamesResponse

	for rows.Next() {
		var game GamesResponse
		var assetsJSON string
		err := rows.Scan(
			&game.Id,
			&game.Code,
			&game.URL,
			&game.Name,
			&game.IsPortrait,
			&game.Description,
			&game.GamePreviews,
			&assetsJSON,
			&game.Categories,
			&game.ColorsM,
			&game.ColorsV,
			&game.Status,
			&game.PrivateAllowed,
			&game.Rating,
			&game.NumberOfRatings,
			&game.GamePlays,
			&game.HasIntegratedAds,
			&game.Width,
			&game.Height,
			&game.Trending,
			&game.Popular,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		err = json.Unmarshal([]byte(assetsJSON), &game.Assets)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return games, nil
}

func GetAllGamesBycategory(db *sql.DB, category_id string) ([]GamesResponse, error) {
	query := `SELECT id, code, url, name, isPortrait, description, gamepreviews, assets, category_id, colorMuted, colorVibrant, status, privateAllowed, rating, numberOfRatings, gamePlays, hasIntegratedAds, width, height,trending,popular FROM games_data WHERE status=true AND category_id=$1;`

	rows, err := db.Query(query, category_id)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var games []GamesResponse

	for rows.Next() {
		var game GamesResponse
		var assetsJSON string
		err := rows.Scan(
			&game.Id,
			&game.Code,
			&game.URL,
			&game.Name,
			&game.IsPortrait,
			&game.Description,
			&game.GamePreviews,
			&assetsJSON,
			&game.Categories,
			&game.ColorsM,
			&game.ColorsV,
			&game.Status,
			&game.PrivateAllowed,
			&game.Rating,
			&game.NumberOfRatings,
			&game.GamePlays,
			&game.HasIntegratedAds,
			&game.Width,
			&game.Height,
			&game.Trending,
			&game.Popular,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		err = json.Unmarshal([]byte(assetsJSON), &game.Assets)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return games, nil
}

func GetGameSearch(db *sql.DB, search string) ([]GamesResponse, error) {
	query := `SELECT id, code, url, name, isPortrait, description, gamepreviews, assets, category_id, colorMuted, colorVibrant, status, privateAllowed, rating, numberOfRatings, gamePlays, hasIntegratedAds, width, height,trending,popular FROM games_data WHERE status=true AND (name ILIKE '%' || $1 || '%' OR name ILIKE $1);`

	rows, err := db.Query(query, search)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var games []GamesResponse

	for rows.Next() {
		var game GamesResponse
		var assetsJSON string
		err := rows.Scan(
			&game.Id,
			&game.Code,
			&game.URL,
			&game.Name,
			&game.IsPortrait,
			&game.Description,
			&game.GamePreviews,
			&assetsJSON,
			&game.Categories,
			&game.ColorsM,
			&game.ColorsV,
			&game.Status,
			&game.PrivateAllowed,
			&game.Rating,
			&game.NumberOfRatings,
			&game.GamePlays,
			&game.HasIntegratedAds,
			&game.Width,
			&game.Height,
			&game.Trending,
			&game.Popular,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		err = json.Unmarshal([]byte(assetsJSON), &game.Assets)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return games, nil
}

func UpdateGameStatus(db *sql.DB, key, id string, value bool) error {

	query := fmt.Sprintf("UPDATE games_data SET %s = $1 WHERE id = $2", key)
	if _, err := db.Exec(query, value, id); err != nil {
		return err
	}
	return nil
}

func AdminGetAllGames(db *sql.DB) ([]GamesResponse, error) {
	query := `SELECT id, code, url, name, isPortrait, description, gamepreviews, assets, category_id, colorMuted, colorVibrant, status, privateAllowed, rating, numberOfRatings, gamePlays, hasIntegratedAds, width, height,trending,popular FROM games_data`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("query failed: %w", err)
	}
	defer rows.Close()

	var games []GamesResponse

	for rows.Next() {
		var game GamesResponse
		var assetsJSON string
		err := rows.Scan(
			&game.Id,
			&game.Code,
			&game.URL,
			&game.Name,
			&game.IsPortrait,
			&game.Description,
			&game.GamePreviews,
			&assetsJSON,
			&game.Categories,
			&game.ColorsM,
			&game.ColorsV,
			&game.Status,
			&game.PrivateAllowed,
			&game.Rating,
			&game.NumberOfRatings,
			&game.GamePlays,
			&game.HasIntegratedAds,
			&game.Width,
			&game.Height,
			&game.Trending,
			&game.Popular,
		)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		err = json.Unmarshal([]byte(assetsJSON), &game.Assets)
		if err != nil {
			return nil, fmt.Errorf("row scan failed: %w", err)
		}
		games = append(games, game)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return games, nil
}
