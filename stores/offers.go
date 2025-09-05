package stores

import (
	"PerkHub/model"
	"PerkHub/services"
	"database/sql"
	"fmt"
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
	perPage := 100
	now := time.Now()
	startDate := now.Format("2006-01-02")
	oneMonthLater := now.AddDate(0, 1, 0) // add 1 month
	endDate := oneMonthLater.Format("2006-01-02")

	fmt.Printf("startDate:%s,endDate:%s", startDate, endDate)

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

					err = model.InsertOffers(s.db, &model.Offer{
						OfferID:     int64(v.ID),
						StoreID:     store[0].ID,
						StoreName:   v.CampaignName,
						Title:       v.Title,
						Description: v.Description,
						TermsAndCondition: func() string {
							if v.TermsAndConditions != "" {
								return v.TermsAndConditions
							}
							return ""
						}(),
						CouponCode: v.CouponCode,
						Image:      v.ImageURL,
						Type:       offerType,
						Status:     v.Status == "live",
						URL:        v.AffiliateUrl,
						StartDate:  v.StartDate, // must be *time.Time
						EndDate:    v.EndDate,   // must be *time.Time
					})
					if err != nil {
						fmt.Println("Error inserting offer:", err)
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
