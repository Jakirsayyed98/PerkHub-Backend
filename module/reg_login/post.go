package reglogin

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"
	"errors"

	"github.com/gin-gonic/gin"
)

func LoginRegistration(c *gin.Context) {
	s, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
	}

	var login request.LoginRequest
	err = c.ShouldBindJSON(&login)
	if err != nil {
		settings.StatusBadRequest(c, errors.New("number is required"), "")
		return
	}

	err = s.LoginStore.RegistrationLogin(login.Number)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	settings.StatusOk(c, nil, "OTP sent Successfully", "")
}

func VerifyOTP(c *gin.Context) {
	s, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	var login request.LoginRequest
	err = c.ShouldBindJSON(&login)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	result, err := s.LoginStore.VerifyOTP(&login)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	settings.StatusOk(c, result, "OTP verified Successfully", "")

}

func SaveUserDetail(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
	}

	request := request.NewSaveUserDetail()

	if err := c.ShouldBindJSON(&request); err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
	}

	err = store.LoginStore.SaveUserDetail(c.MustGet("user_id").(string), *request)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
	}
	settings.StatusOk(c, nil, "Details Saved Successfully", "")
	return

}

func GetUserDetail(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
	}

	result, err := store.LoginStore.GetUserDetail(c.MustGet("user_id").(string))
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
	}

	settings.StatusOk(c, result, "Successfully get User Detail", "")
}
