package response

import (
	"PerkHub/model"
	"database/sql"
	"fmt"
)

type UserDetailResponse struct {
	User_Id   string `json:"user_id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Number    string `json:"number"`
	OTP       string `json:"otp"`
	Gender    string `json:"gender"`
	Dob       string `json:"dob"`
	FCMToken  string `json:"fcm_token"`
	Verified  string `json:"verified"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

func NewUserDetailRes() UserDetailResponse {
	return UserDetailResponse{}
}

func (u *UserDetailResponse) ResponsesBind(dbUser *model.UserDetail) error {
	u.User_Id = dbUser.User_id.String
	u.Name = dbUser.Name.String
	u.Email = dbUser.Email.String
	u.Number = dbUser.Number.String
	u.OTP = dbUser.OTP.String
	u.Gender = dbUser.Gender.String
	u.Dob = dbUser.Dob.String
	u.FCMToken = dbUser.FCMToken.String
	u.Verified = dbUser.Verified.String
	u.CreatedAt = dbUser.CreatedAt.String
	u.UpdatedAt = dbUser.UpdatedAt.String

	return nil
}

type Token struct {
	Token string `json:"token"`
}

func ConvertNullString(nullStr sql.NullString) string {
	if nullStr.Valid {
		return nullStr.String
	}
	return ""
}

func (u *UserDetailResponse) BindMultipleUsers(dbUsers []*model.UserDetail) ([]UserDetailResponse, error) {
	var responses []UserDetailResponse

	for _, dbUser := range dbUsers {
		var response UserDetailResponse
		err := response.ResponsesBind(dbUser)
		if err != nil {
			return nil, fmt.Errorf("error binding user detail: %w", err)
		}
		responses = append(responses, response)
	}

	return responses, nil
}
