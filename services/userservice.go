package services

import (
	"PerkHub/constants"
	"PerkHub/model"
	"PerkHub/settings"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UserService struct {
	service *settings.HttpService
}

func NewUserService() *UserService {
	return &UserService{
		service: settings.NewHttpService("https://www.fast2sms.com"),
	}
}

func (s *UserService) SendOTPService(number, otp string) (interface{}, error) {
	requestBody := map[string]string{
		"route":            "otp",  // OTP route
		"variables_values": otp,    // OTP value
		"numbers":          number, // Phone number
	}

	jsonMarshal, err := json.Marshal(requestBody)
	if err != nil {
		return nil, fmt.Errorf("error marshaling JSON: %w", err)
	}

	return nil, nil

	fmt.Println("Request Body:", string(jsonMarshal)) // Optional, for debugging

	headers := map[string]string{
		"Authorization": constants.FAST2SMS_API_KEY,
		"Content-Type":  "application/json",
	}

	response, err := s.service.Post("/dev/bulkV2", jsonMarshal, headers)
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

	res := model.NewResponseOTP()
	if err := json.Unmarshal(body, &res); err != nil {
		return nil, fmt.Errorf("error unmarshaling response JSON: %w", err)
	}

	if !res.Return && len(res.Message) > 0 {
		return nil, fmt.Errorf("error from API: %s", res.Message[0])
	}

	return res, nil
}
