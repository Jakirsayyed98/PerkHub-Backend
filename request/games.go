package request

import "encoding/json"

type GameResponse struct {
	Code             string       `json:"code"`
	URL              string       `json:"url"`
	Name             Name         `json:"name"`
	IsPortrait       bool         `json:"isPortrait"`
	Description      Description  `json:"description"`
	GamePreviews     GamePreviews `json:"gamePreviews"` // New field
	Assets           Assets       `json:"assets"`
	Categories       Categories   `json:"categories"`
	Tags             Tags         `json:"tags"`         // Combine width and height into this struct
	ColorsM          string       `json:"colorMuted"`   // Combine colorMuted and colorVibrant into this struct
	Colors           string       `json:"colorVibrant"` // Combine colorMuted and colorVibrant into this struct
	PrivateAllowed   bool         `json:"privateAllowed"`
	Rating           float64      `json:"rating"`
	NumberOfRatings  int          `json:"numberOfRatings"`
	GamePlays        int          `json:"gamePlays"`
	HasIntegratedAds bool         `json:"hasIntegratedAds"`
	Width            int          `json:"width"`
	Height           int          `json:"height"`
}

type Name struct {
	EN string `json:"en"`
}

type Description struct {
	EN string `json:"en"`
}

type GamePreviews struct {
	EN string `json:"en"`
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

type Categories struct {
	EN []string `json:"en"`
}

type Tags struct {
	EN []string `json:"en"`
}

func NewGameResponse() *GameResponse {
	return &GameResponse{}
}

func (s *GameResponse) Unmarshal(data []byte) ([]GameResponse, error) {
	var wrapper struct {
		Games []GameResponse `json:"games"`
	}
	if err := json.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}
	return wrapper.Games, nil
}
