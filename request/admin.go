package request

type AdminAffiliateTransactionsRequest struct {
	PageNo int `json:"page_no"`
	Limit  int `json:"limit"` // Corrected spelling
}

func NewAdminAffiliateTransactionsRequest() *AdminAffiliateTransactionsRequest {
	return &AdminAffiliateTransactionsRequest{}
}
