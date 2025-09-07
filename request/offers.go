package request

type SearchOffers struct {
	BrandName string `json:"brand_name"`
}

func NewSearchOffers() *SearchOffers {
	return &SearchOffers{}
}
