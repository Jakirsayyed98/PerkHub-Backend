package responses

import (
	"PerkHub/model"
	"fmt"
	"time"
)

type MiniAppRes struct {
	ID                string    `db:"id" json:"id"`                                   // Unique identifier
	MiniAppCategoryID string    `db:"miniapp_category_id" json:"miniapp_category_id"` // Category ID
	Name              string    `db:"name" json:"name"`                               // Name of the miniapp
	Icon              string    `db:"icon" json:"icon"`                               // URL or path to the icon
	Description       string    `db:"description" json:"description"`                 // Description of the miniapp
	CashbackTerms     string    `db:"cashback_terms" json:"cashback_terms"`           // Terms for cashback
	CashbackRates     string    `db:"cashback_rates" json:"cashback_rates"`           // Rates for cashback
	Status            bool      `db:"status" json:"status"`                           // Status: '0' for inactive, '1' for active
	UrlType           string    `db:"url_type" json:"url_type"`                       // Type of URL
	CBActive          bool      `db:"is_cb_active" json:"is_cb_active"`               // Cashback active status
	CBPercentage      string    `db:"cb_percentage" json:"cb_percentage"`             // Cashback percentage
	Url               string    `db:"url" json:"url"`                                 // URL of the miniapp
	Label             string    `db:"label" json:"label"`                             // Label for the miniapp
	Banner            string    `db:"banner" json:"banner"`                           // Banner URL
	Logo              string    `db:"logo" json:"logo"`                               // Logo URL
	MacroPublisher    string    `db:"macro_publisher" json:"macro_publisher"`         // Publisher name
	Popular           bool      `db:"popular" json:"popular"`                         // Popular status
	Trending          bool      `db:"trending" json:"trending"`                       // Trending status
	TopCashback       bool      `db:"top_cashback" json:"top_cashback"`               // Top cashback status
	About             string    `db:"about" json:"about"`                             // About information
	HowItsWork        string    `db:"howitswork" json:"howitswork"`                   // How it works information
	CreatedAt         time.Time `db:"created_at" json:"created_at"`                   // Creation timestamp
	UpdatedAt         time.Time `db:"updated_at" json:"updated_at"`
}

func NewMiniAppRes() *MiniAppRes {
	return &MiniAppRes{}
}

func (u *MiniAppRes) BindMiniAppResponse(miniapps []model.MiniApp) ([]MiniAppRes, error) {
	var responses []MiniAppRes

	if len(miniapps) == 0 {
		return responses, nil
	}

	for _, dbUser := range miniapps {
		var response MiniAppRes
		err := response.ResponsesBind(&dbUser)
		if err != nil {
			return nil, fmt.Errorf("error binding user detail: %w", err)
		}
		responses = append(responses, response)
	}

	return responses, nil
}
func (m *MiniAppRes) ResponsesBind(dbMiniApp *model.MiniApp) error {
	m.ID = dbMiniApp.ID
	m.Name = dbMiniApp.Name
	m.Icon = dbMiniApp.Icon
	m.Logo = dbMiniApp.Logo
	m.MiniAppCategoryID = dbMiniApp.MiniAppCategoryID
	m.Description = dbMiniApp.Description
	m.About = dbMiniApp.About
	m.CashbackTerms = dbMiniApp.CashbackTerms
	m.CBActive = dbMiniApp.CBActive
	m.CBPercentage = dbMiniApp.CBPercentage
	m.Url = dbMiniApp.Url
	m.UrlType = dbMiniApp.UrlType
	m.MacroPublisher = dbMiniApp.MacroPublisher
	m.Status = dbMiniApp.Active
	m.Popular = dbMiniApp.Popular
	m.Trending = dbMiniApp.Trending
	m.TopCashback = dbMiniApp.TopCashback
	m.CreatedAt = dbMiniApp.CreatedAt
	m.UpdatedAt = dbMiniApp.UpdatedAt
	m.Banner = dbMiniApp.Banner

	return nil
}
