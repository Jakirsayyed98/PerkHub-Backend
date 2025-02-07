package request

import (
	"PerkHub/utils"
	"fmt"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type MiniAppRequest struct {
	ID                   uuid.UUID `db:"id" json:"id"`                                         // Unique identifier
	MiniAppCategoryID    string    `db:"miniapp_category_id" json:"miniapp_category_id"`       // Category ID
	MiniAppSubcategoryID string    `db:"miniapp_subcategory_id" json:"miniapp_subcategory_id"` // Subcategory ID
	Name                 string    `db:"name" json:"name"`                                     // Name of the miniapp
	Icon                 string    `db:"icon" json:"icon"`                                     // URL or path to the icon
	Description          string    `db:"description" json:"description"`                       // Description of the miniapp
	CashbackTerms        string    `db:"cashback_terms" json:"cashback_terms"`                 // Terms for cashback
	CashbackRates        string    `db:"cashback_rates" json:"cashback_rates"`                 // Rates for cashback
	Status               bool      `db:"status" json:"status"`                                 // Status: '0' for inactive, '1' for active
	UrlType              string    `db:"url_type" json:"url_type"`                             // Type of URL
	CBActive             bool      `db:"cb_active" json:"cb_active"`                           // Cashback active status
	CBPercentage         string    `db:"cb_percentage" json:"cb_percentage"`                   // Cashback percentage
	Url                  string    `db:"url" json:"url"`                                       // URL of the miniapp
	Label                string    `db:"label" json:"label"`                                   // Label for the miniapp
	Banner               string    `db:"banner" json:"banner"`                                 // Banner URL
	Logo                 string    `db:"logo" json:"logo"`                                     // Logo URL
	MacroPublisher       string    `db:"macro_publisher" json:"macro_publisher"`               // Publisher name
	Popular              bool      `db:"popular" json:"popular"`                               // Popular status
	Trending             bool      `db:"trending" json:"trending"`                             // Trending status
	TopCashback          bool      `db:"top_cashback" json:"top_cashback"`                     // Top cashback status
	About                string    `db:"about" json:"about"`                                   // About information
	HowItsWork           string    `db:"howitswork" json:"howitswork"`                         // How it works information
	HomepageVisible      bool      `db:"homepage_visible" json:"homepage_visible"`             // Visibility on homepage
	CreatedAt            time.Time `db:"created_at" json:"created_at"`                         // Creation timestamp
	UpdatedAt            time.Time `db:"updated_at" json:"updated_at"`
}

func NewMiniAppRequest() MiniAppRequest {
	return MiniAppRequest{}
}

func (req *MiniAppRequest) Bind(c *gin.Context) error {
	if !strings.Contains(c.Request.Header.Get("Content-Type"), "multipart/form-data") {
		return fmt.Errorf("content type not supported %s", c.Request.Header.Get("Content-Type"))
	}
	form, err := c.MultipartForm()
	if err != nil {
		return err
	}
	logo := ""
	banner := ""
	icon := ""
	iconFile := form.File["icon"]
	bannerFile := form.File["banner"]
	logoFile := form.File["logo"]
	if len(iconFile) > 0 {
		file := iconFile[0]
		icon, err = utils.SaveFile(c, file)
		if err != nil {
			return err
		}
	}
	if len(bannerFile) > 0 {
		file := bannerFile[0]
		banner, err = utils.SaveFile(c, file)
		if err != nil {
			return err
		}
	}
	if len(logoFile) > 0 {
		file := logoFile[0]
		logo, err = utils.SaveFile(c, file)
		if err != nil {
			return err
		}
	}
	req.ID, err = uuid.Parse(c.PostForm("id"))
	req.MiniAppCategoryID = c.PostForm("miniapp_category_id")
	req.MiniAppSubcategoryID = c.PostForm("miniapp_subcategory_id")
	req.Name = c.PostForm("name")
	req.Icon = icon
	req.Description = c.PostForm("description")
	req.CashbackTerms = c.PostForm("cashback_terms")
	req.CashbackRates = c.PostForm("cashback_rates")
	// req.Status = c.PostForm("status")
	req.UrlType = c.PostForm("url_type")
	// req.CBActive = c.PostForm("cb_active")
	req.CBPercentage = c.PostForm("cb_percentage")
	req.Url = c.PostForm("url")
	req.Label = c.PostForm("label")
	req.Banner = banner
	req.Logo = logo
	req.MacroPublisher = c.PostForm("macro_publisher")
	// req.Popular = c.PostForm("popular")
	// req.Trending = c.PostForm("trending")
	// req.TopCashback = c.PostForm("top_cashback")
	req.About = c.PostForm("about")
	req.HowItsWork = c.PostForm("howitswork")
	req.CreatedAt = time.Now()
	req.UpdatedAt = time.Now()

	return nil
}

type ActiveDeactiveMiniAppReq struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value string `json:"value"`
	// Top_cashback bool `json:"top_cashback"`
	// Trending     bool `json:"trending"`
	// Popular      bool `json:"popular"`
	// Status       bool `json:"status"`
	// Url_type     bool `json:"url_type"`
	// Cb_active    bool `json:"cb_active"`
}

func NewActiveDeactiveminiAppReq() *ActiveDeactiveMiniAppReq {
	return &ActiveDeactiveMiniAppReq{}
}

type MiniAppSearchReq struct {
	Name string `json:"name"`
}

func NewMiniAppSearchReq() *MiniAppSearchReq {
	return &MiniAppSearchReq{}
}

type DeleteMiniApp struct {
	Id string `json:"id"`
}

func NewDeleteMiniApp() *DeleteMiniApp {
	return &DeleteMiniApp{}
}

type GenrateMiniAppSubId struct {
	MiniAppName string `json:"name"  binding:"required"`
}

func NewGenrateMiniAppSubId() *GenrateMiniAppSubId {
	return &GenrateMiniAppSubId{}
}
