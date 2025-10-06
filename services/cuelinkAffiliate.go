package services

import (
	"PerkHub/request"
	"PerkHub/responses"
	"PerkHub/settings"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CueLinkAffiliateService struct {
	baseURL   string
	authToken string
	service   *settings.HttpService
}

func NewCueLinkAffiliateService() *CueLinkAffiliateService {
	return &CueLinkAffiliateService{
		service:   settings.NewHttpService("https://www.cuelinks.com/api/v2/"),
		baseURL:   "https://www.cuelinks.com/api/v2/",
		authToken: "ffNEpbLQb5dkXTpgk5EgappLEslQNsvoOfq7KG4zkrk",
	}
}

// page and perPage are dynamic
func (s *CueLinkAffiliateService) GetAllCampaigns(page, perPage int) (responses.CampaignResponse, error) {
	var result responses.CampaignResponse

	// Build URL with query parameters dynamically
	url := fmt.Sprintf("%sall_campaigns.json?sort_column=id&sort_direction=asc&page=%d&per_page=%d&country_id=%d", s.baseURL, page, perPage, 252)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+s.authToken)
	req.Header.Set("Cookie", `DO-LB="ChAxMC4xMzkuMTIwLjI1OjgwEJrBtxc="; _mkra_stck=9b85040dec486f43e8f6afdad672d662%3A1756976050.725129`)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// Handle HTTP status codes
	switch resp.StatusCode {
	case http.StatusOK:
		// Unmarshal JSON into CampaignResponse struct
		err = json.Unmarshal(body, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	case http.StatusNoContent:
		// Return empty campaigns slice for 204 No Content
		return responses.CampaignResponse{Campaigns: []responses.Campaign{}}, nil
	default:
		return result, fmt.Errorf("unknown error: status code %d", resp.StatusCode)
	}
}

func (s *CueLinkAffiliateService) RefreshAllOffers(startDate, endDate string, offerType, page, perPage int) (responses.OfferResponse, error) {
	var result responses.OfferResponse

	// Build URL with query parameters dynamically
	url := fmt.Sprintf("%soffers.json?start_date=%s&end_date=%s&offer_types=%d&page=%d&per_page=%d", s.baseURL, startDate, endDate, offerType, page, perPage)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+s.authToken)
	req.Header.Set("Cookie", `DO-LB="ChAxMC4xMzkuMTIwLjI1OjgwEJrBtxc="; _mkra_stck=9b85040dec486f43e8f6afdad672d662%3A1756976050.725129`)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}

	// Handle HTTP status codes
	switch resp.StatusCode {
	case http.StatusOK:
		// Unmarshal JSON into CampaignResponse struct
		err = json.Unmarshal(body, &result)
		if err != nil {
			return result, err
		}
		return result, nil
	case http.StatusNoContent:
		// Return empty campaigns slice for 204 No Content
		return responses.OfferResponse{Offers: []responses.Offer{}}, nil
	default:
		return result, fmt.Errorf("unknown error: status code %d", resp.StatusCode)
	}
}

func (s *CueLinkAffiliateService) CheckAffiliateTransaction(request *request.CueLinkTransactionCheckRequest) (*responses.AffiliateTransaction, error) {
	jsonMarshal, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	headers := map[string]string{
		"Authorization": "Bearer " + s.authToken,
		"Content-Type":  "application/json",
	}

	response, err := s.service.Post("transactions.json", jsonMarshal, headers)
	if err != nil {
		return nil, fmt.Errorf("error sending HTTP request: %w", err)
	}

	defer response.Body.Close()
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non-OK HTTP status: %d", response.StatusCode)
	}

	res := responses.NewAffiliateTransactionResponse()
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}

	return res, nil
}
