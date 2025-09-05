package responses

import "time"

type CampaignResponse struct {
	Campaigns []Campaign `json:"campaigns"`
}

type Campaign struct {
	ID                    int               `json:"id"`
	Name                  string            `json:"name"`
	URL                   string            `json:"url"`
	Domain                string            `json:"domain"`
	PayoutType            string            `json:"payout_type"`
	Payout                float64           `json:"payout"`
	Image                 string            `json:"image"`
	AdditionalInfo        string            `json:"additional_info"`
	AdditionalInfoHTML    string            `json:"additional_info_html"`
	ImportantInfoHTML     string            `json:"important_info_html"`
	LastModified          time.Time         `json:"last_modified"`
	Status                string            `json:"status"`
	ApplicationStatus     string            `json:"application_status"`
	PayoutCategories      []PayoutCategory  `json:"payout_categories"`
	Categories            []Category        `json:"categories"`
	Countries             []Country         `json:"countries"`
	ReportingType         string            `json:"reporting_type"`
	DeeplinkAllowed       bool              `json:"deeplink_allowed"`
	SubIDsAllowed         bool              `json:"sub_ids_allowed"`
	CashbackPublishers    bool              `json:"cashback_publishers_allowed"`
	SocialMediaPublishers bool              `json:"social_media_publishers_allowed"`
	MissingTransactions   bool              `json:"missing_transactions_accepted"`
	CookieDuration        string            `json:"cookie_duration"`
	AllowedPlatforms      AllowedPlatforms  `json:"allowed_platforms"`
	AllowedMediums        AllowedMediums    `json:"allowed_mediums"`
	ConversionFlow        map[string]string `json:"conversion_flow"`
}

type PayoutCategory struct {
	Name       string  `json:"name"`
	PayoutType string  `json:"payout_type"`
	Payout     float64 `json:"payout"`
}

type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Country struct {
	ID   int    `json:"id"`
	ISO  string `json:"iso"`
	Name string `json:"name"`
}

type AllowedPlatforms struct {
	Web       bool `json:"web"`
	MobileWeb bool `json:"mobile_web"`
	Android   bool `json:"android"`
	IOS       bool `json:"ios"`
}

type AllowedMediums struct {
	TextLink           bool `json:"text_link"`
	Banner             bool `json:"banner"`
	Deals              bool `json:"deals"`
	Coupons            bool `json:"coupons"`
	Cashback           bool `json:"cashback"`
	Email              bool `json:"email"`
	CustomEmail        bool `json:"custom_email"`
	PopTraffic         bool `json:"pop_traffic"`
	NativeAds          bool `json:"native_ads"`
	FacebookAds        bool `json:"facebook_ads"`
	SEMBrandKeywords   bool `json:"sem_brand_keywords"`
	SEMGenericKeywords bool `json:"sem_generic_keywords"`
}

// Offers Response Strucher
type OfferResponse struct {
	Status     string  `json:"status"`
	TotalCount int     `json:"total_count"`
	Offers     []Offer `json:"offers"`
}

type Offer struct {
	ID                 int    `json:"id"`
	CampaignID         int    `json:"campaign_id"`
	CampaignName       string `json:"campaign_name"`
	Title              string `json:"title"`
	Description        string `json:"description"`
	TermsAndConditions string `json:"terms_and_conditions,omitempty"`
	CouponCode         string `json:"coupon_code,omitempty"`
	ImageURL           string `json:"image_url"`
	OfferType          string `json:"type"`
	Status             string `json:"status"`
	Url                string `json:"url"`
	AffiliateUrl       string `json:"affiliate_url"`
	StartDate          string `json:"start_date"`
	EndDate            string `json:"end_date"`
}
