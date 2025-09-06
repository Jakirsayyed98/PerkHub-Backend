package stores

import (
	"PerkHub/model"
	"PerkHub/services"
	"database/sql"
	"errors"
	"strconv"
	"time"
)

type OffersStore struct {
	db             *sql.DB
	cueLinkService *services.CueLinkAffiliateService
}

func NewOffersStore(dbs *sql.DB) *OffersStore {
	cuelinkService := services.NewCueLinkAffiliateService()
	return &OffersStore{
		db:             dbs,
		cueLinkService: cuelinkService,
	}
}

func (s *OffersStore) GetOffersRefresh() (interface{}, error) {
	page := 1
	perPage := 50
	now := time.Now()
	startDate := now.Format("2006-01-02")
	oneMonthLater := now.AddDate(0, 1, 0) // add 1 month
	endDate := oneMonthLater.Format("2006-01-02")
	for {
		// Fetch campaigns for the current page
		data, err := s.cueLinkService.RefreshAllOffers(startDate, endDate, 1, page, perPage)
		if err != nil {
			return nil, err
		}

		// Stop if no campaigns are returned
		if len(data.Offers) == 0 {
			break
		}

		for _, v := range data.Offers {
			store, err := model.SearchMiniApps(s.db, v.CampaignName)
			if err != nil {
				return nil, err
			}
			if len(store) > 0 {
				isExist, err := model.OfferExists(s.db, strconv.Itoa(v.ID))
				if err != nil {
					return nil, err
				}
				if !isExist {
					offerType := "offer"
					if v.CouponCode != "" {
						offerType = "coupon"
					}

					err = model.InsertOffer(s.db, &model.Offer{
						OfferID:           strconv.Itoa(v.ID),
						StoreID:           store[0].ID,
						StoreName:         v.CampaignName,
						Title:             v.Title,
						Description:       v.Description,
						TermsAndCondition: "",
						CouponCode:        v.CouponCode,
						Image:             v.ImageURL,
						Type:              offerType,
						Status:            v.Status == "live",
						URL:               v.AffiliateUrl,
						StartDate:         v.StartDate, // must be *time.Time
						EndDate:           v.EndDate,   // must be *time.Time
					})
					if err != nil {
						return nil, err
					}

				}
			}

		}
		// Move to the next page
		page++
	}
	return nil, nil
}

func (s *OffersStore) GetAllActiveOffersList(offerType string) ([]model.Offer, error) {
	offers, err := model.GetAllOfferList(s.db, offerType)
	if err != nil {
		return nil, err
	}
	return offers, nil
}

func (s *OffersStore) SearchOffersByStoreName(storeName string) ([]model.Offer, error) {
	if storeName == "" {
		return nil, errors.New("please pass store name")
	}

	offers, err := model.SearchOffersByStoreName(s.db, storeName)
	if err != nil {
		return nil, err
	}

	return offers, nil
}

func (s *OffersStore) OffersForHomePage() (interface{}, error) {
	offers, err := model.GetRandomOffers(s.db)
	if err != nil {
		return nil, err
	}
	return offers, nil
}
