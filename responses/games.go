package responses

import "encoding/json"

type GameResponse struct {
	Code             string      `json:"code"`
	URL              string      `json:"url"`
	Name             Name        `json:"name"`
	IsPortrait       bool        `json:"isPortrait"`
	Description      Description `json:"description"`
	Assets           Assets      `json:"assets"`
	Categories       Categories  `json:"categories"`
	Tags             Tags        `json:"tags"`
	Dimensions       Dimensions  `json:"dimensions"`
	Colors           Colors      `json:"colors"`
	PrivateAllowed   bool        `json:"privateAllowed"`
	Rating           float64     `json:"rating"`
	NumberOfRatings  int         `json:"numberOfRatings"`
	GamePlays        int         `json:"gamePlays"`
	HasIntegratedAds bool        `json:"hasIntegratedAds"`
}

type Name struct {
	EN string `json:"en"`
}

type Description struct {
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

type Dimensions struct {
	Width  int `json:"width"`
	Height int `json:"height"`
}

type Colors struct {
	Muted   string `json:"muted"`
	Vibrant string `json:"vibrant"`
}

func NewGameResponse() *GameResponse {
	return &GameResponse{}
}

func (s *GameResponse) Unmarshal(data []byte) ([]GameResponse, error) {
	var res []GameResponse

	if err := json.Unmarshal(data, res); err != nil {
		return nil, err
	}
	return res, nil
}
