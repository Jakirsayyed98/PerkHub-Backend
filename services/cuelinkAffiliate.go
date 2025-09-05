package services

import (
	"PerkHub/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CueLinkAffiliateService struct {
	baseURL   string
	authToken string
}

func NewCueLinkAffiliateService() *CueLinkAffiliateService {
	return &CueLinkAffiliateService{
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
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	// Set headers
	req.Header.Set("Authorization", "Bearer "+s.authToken)
	req.Header.Set("Cookie", `DO-LB="ChAxMC4xMzkuMTIwLjI1OjgwEJrBtxc="; _mkra_stck=9b85040dec486f43e8f6afdad672d662%3A1756976050.725129`)
	fmt.Println(req.Header)
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
