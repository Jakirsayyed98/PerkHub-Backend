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
	"time"
)

type LoginStore struct {
	db          *sql.DB
	userService *services.UserService
}

func NewLoginStore(dbs *sql.DB) *LoginStore {
	userService := services.NewUserService()
	return &LoginStore{
		db:          dbs,
		userService: userService,
	}
}

func (s *LoginStore) RegistrationLogin(number string) error {
	if number == "" {
		return errors.New("number required")
	}

	otp := utils.GenerateNumber(6)
	_, err := s.userService.SendOTPService(number, otp)
	if err != nil {
		return err
	}

	err = model.InsertOTPRequest(s.db, number, otp)
	// err = model.InsertLoginData(s.db, number, otp)

	if err != nil {
		return err
	}

	return nil
}

func (s *LoginStore) VerifyOTP(login *request.LoginRequest) (interface{}, error) {

	if login.Number == "" {
		return nil, errors.New("number required")
	}
	if login.OTP == "" {
		return nil, errors.New("otp required")
	}

	latestOTP, err := model.GetLatestOTPByNumber(s.db, login.Number)
	if err != nil {
		return nil, err
	}

	if latestOTP != login.OTP {
		return nil, errors.New("invalid otp")
	}

	if err := model.MarkOtpVerified(s.db, login.Number, login.OTP); err != nil {
		return nil, err
	}

	_, err = model.UserDetailByMobileNumber(s.db, login.Number)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, fmt.Errorf("failed to get user details: %w", err)
		}
		if err := model.InsertLoginData(s.db, login.Number); err != nil {
			return nil, fmt.Errorf("failed to insert login data: %w", err)
		}
	}

	// token := responses.Token{}
	// res, err := utils.GenerateJWTToken(userDetail.User_id.String, time.Minute*15)
	// if err != nil {

	// }
	// token.Token = res
	return nil, nil
}

func (s *LoginStore) GetAuthToken(login *request.GetAuthToken) (interface{}, error) {

	if login.Number == "" {
		return nil, errors.New("number required")
	}

	userDetail, err := model.UserDetailByMobileNumber(s.db, login.Number)
	if err != nil {
		return nil, err
	}
	token := responses.Token{}
	res, err := utils.GenerateJWTToken(userDetail.User_id.String, time.Minute*15)
	if err != nil {

	}
	token.Token = res
	return token, nil
}

func (s *LoginStore) SaveUserDetail(user_id string, req request.SaveUserDetailReq) error {

	if err := model.UpdateUserDetail(s.db, user_id, &req); err != nil {
		return err
	}

	return nil

}

func (s *LoginStore) GetUserDetail(user_id string) (interface{}, error) {

	data, err := model.UserDetailByUserID(s.db, user_id)

	if err != nil {
		return nil, err
	}

	res := responses.NewUserDetailRes()

	if err := res.ResponsesBind(data); err != nil {
		return nil, err
	}

	return res, nil
}

func (s *LoginStore) GetAllUserDetail() (interface{}, error) {

	data, err := model.AllUsersDetail(s.db)

	if err != nil {
		return nil, err
	}

	res := responses.NewUserDetailRes()

	result, err := res.BindMultipleUsers(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}
