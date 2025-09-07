package stores

import (
	"PerkHub/model"
	"PerkHub/request"
	"PerkHub/responses"
	"PerkHub/services"
	"PerkHub/utils"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strings"
)

type MiniAppStore struct {
	db             *sql.DB
	cueLinkService *services.CueLinkAffiliateService
}

func NewMiniAppStore(dbs *sql.DB) *MiniAppStore {
	cuelinkService := services.NewCueLinkAffiliateService()
	return &MiniAppStore{
		db:             dbs,
		cueLinkService: cuelinkService,
	}
}

func (s *MiniAppStore) CreateMiniApp(req *request.MiniAppRequest) (interface{}, error) {

	if req.ID != "" {
		if err := model.UpdateMiniApp(s.db, req); err != nil {
			return nil, err
		}
	} else {
		if err := model.InsertMiniApp(s.db, req); err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (s *MiniAppStore) ActivateSomekey(req *request.ActiveDeactiveMiniAppReq) (interface{}, error) {
	if err := model.ToggleMiniAppFlag(s.db, req.Key, req.ID, req.Value); err != nil {
		return nil, err
	}

	return nil, nil
}

func (s *MiniAppStore) GetAllMiniApps() (interface{}, error) {

	data, err := model.GetAllMiniApps(s.db)

	if err != nil {
		return nil, err
	}
	res := responses.NewMiniAppRes()

	result, err := res.BindMultipleUsers(data)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *MiniAppStore) GetStoreByID(id string) (interface{}, error) {

	data, err := model.GetMiniAppByID(s.db, id)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func (s *MiniAppStore) GetStoresByCategory(category_id string) (interface{}, error) {

	data, err := model.GetStoresByCategory(s.db, category_id)

	if err != nil {
		return nil, err
	}

	if len(data) == 0 {
		return nil, errors.New("Store not found")
	}
	res := responses.NewMiniAppRes()

	result, err := res.BindMultipleUsers(data)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *MiniAppStore) SearchMiniApps(req *request.MiniAppSearchReq) (interface{}, error) {

	data, err := model.SearchMiniApps(s.db, req.Name)

	if err != nil {
		return nil, err
	}
	res := responses.NewMiniAppRes()
	if len(data) == 0 {
		return nil, errors.New("Store not found")
	}

	result, err := res.BindMultipleUsers(data)
	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *MiniAppStore) DeletMniApp(id string) (interface{}, error) {
	if err := model.DeleteMiniAppByID(s.db, id); err != nil {
		return nil, err
	}
	return nil, nil
}

func (s *MiniAppStore) GenrateSubid(miniAppName, userID string) (interface{}, error) {
	data, err := model.SearchMiniApps(s.db, miniAppName)

	if err != nil {
		return nil, err
	}
	if data == nil {
		return nil, errors.New("App not found")
	}
	subid1, err := utils.GenerateRandomUUID(20)
	if err != nil {
		return nil, err
	}
	subid2 := userID
	subid3 := data[0].Name
	// Encode the original store URL for Cuelinks
	encodedURL := url.QueryEscape(data[0].Url)

	affiliateURL := fmt.Sprintf(
		"https://linksredirect.com/?cid=198215&source=linkkit&url=%s&subid=%s&subid2=%s&subid3=%s",
		encodedURL, subid1, subid2, subid3,
	)

	// url := fmt.Sprintf("%s&subid=%s&subid2=%s&subid3=%s", data[0].Url, subid1, subid2, subid3)

	err = model.InsertGenratedSubId(s.db, miniAppName, userID, subid1, subid2)
	if err != nil {
		return nil, err
	}
	// Wrap with your own domain (short redirect link)
	// finalURL := fmt.Sprintf("https://www.perkhub.in/r?u=%s", url.QueryEscape(affiliateURL))
	// finalURL := fmt.Sprintf("http://3.7.77.135:4215/r?u=%s", url.QueryEscape(affiliateURL))
	// finalURL := fmt.Sprintf("http://localhost:4215/r?u=%s", url.QueryEscape(affiliateURL))
	finalURL := fmt.Sprintf("https://blessed-pretty-mammal.ngrok-free.app/r?u=%s", url.QueryEscape(affiliateURL))

	return finalURL, nil
}

// func (s *MiniAppStore) GenrateSubid(miniAppName, userID string) (interface{}, error) {
// 	data, err := model.SearchMiniApps(s.db, miniAppName)

// 	if err != nil {
// 		return nil, err
// 	}
// 	if data == nil {
// 		return nil, errors.New("App not found")
// 	}
// 	subid1, err := utils.GenerateRandomUUID(20)
// 	if err != nil {
// 		return nil, err
// 	}
// 	subid2 := userID
// 	subid3 := data[0].Name
// 	url := fmt.Sprintf("%s&subid=%s&subid2=%s&subid3=%s", data[0].Url, subid1, subid2, subid3)

// 	err = model.InsertGenratedSubId(s.db, miniAppName, userID, subid1, subid2)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return url, nil
// }

func (s *MiniAppStore) GetStoresRefresh() (interface{}, error) {
	page := 1
	perPage := 100
	affiliateProviderID, err := model.GetAffiliateByName(s.db, "cuelink")
	if err != nil {
		return nil, err
	}
	for {
		// Fetch campaigns for the current page
		data, err := s.cueLinkService.GetAllCampaigns(page, perPage)
		if err != nil {
			return nil, err
		}

		// Stop if no campaigns are returned
		if len(data.Campaigns) == 0 {
			break
		}

		for _, v := range data.Campaigns {
			isExist, err := model.MiniAppExists(s.db, v.Name)
			if err != nil {
				return nil, err
			}
			if !isExist {
				categoryId, err := model.CategoryByName(s.db, v.Categories[0].Name)
				if err != nil {
					return nil, err
				}
				payoutClean := ""
				if strings.Contains(v.PayoutType, "%") {
					v.PayoutType = "Percentage"
					formatted := fmt.Sprintf("%.2f%%", v.Payout) // "9.75%"
					payoutClean = formatted
				} else {
					v.PayoutType = "Fixed"
					payoutClean = fmt.Sprintf("₹%.2f", v.Payout) // e.g., "₹123.45"

				}

				url := ""
				if !strings.HasPrefix(v.Domain, "https") {
					url = fmt.Sprintf("https://%s", v.Domain)
				} else {
					url = v.Domain
				}
				fmt.Println(url)
				err = model.InsertMiniApp(s.db, &request.MiniAppRequest{
					MiniAppCategoryID: categoryId, // default or your desired category ID
					Name:              v.Name,
					Icon:              v.Image, // optional
					Logo:              "",      // optional
					Description:       "",
					About:             "",                  // optional
					CashbackTerms:     v.PayoutType,        // optional
					CBActive:          true,                // default false
					CBPercentage:      payoutClean,         // default 0
					Url:               url,                 // optional
					UrlType:           "internal",          // optional
					MacroPublisher:    affiliateProviderID, // optional
					Status:            true,                // active by default
					Popular:           false,               // default false
					Trending:          false,               // default false
					TopCashback:       false,               // default false
				})
				if err != nil {
					fmt.Println("Error inserting category:", err)
					return nil, err
				}
			}
		}
		// Move to the next page
		page++
	}
	return nil, nil
}
