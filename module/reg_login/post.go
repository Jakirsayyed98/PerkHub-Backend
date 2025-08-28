package reglogin

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"
	"PerkHub/utils"
	"errors"

	"github.com/gin-gonic/gin"
)

func LoginRegistration(c *gin.Context) {
	s, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	var login request.LoginRequest
	err = c.ShouldBindJSON(&login)
	if err != nil {
		utils.RespondBadRequest(c, errors.New("number is required"), "")
		return
	}

	err = s.LoginStore.RegistrationLogin(login.Number)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, nil, "OTP sent Successfully", "")
}

func GetAuthToken(c *gin.Context) {
	s, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	var request request.GetAuthToken
	err = c.ShouldBindJSON(&request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := s.LoginStore.GetAuthToken(&request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Token Got Successfully", "")
}

func VerifyOTP(c *gin.Context) {
	s, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	var login request.LoginRequest
	err = c.ShouldBindJSON(&login)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
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
		utils.RespondBadRequest(c, err, "")
		return
	}

	request := request.NewSaveUserDetail()

	if err := c.ShouldBindJSON(&request); err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	err = store.LoginStore.SaveUserDetail(c.MustGet("user_id").(string), *request)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, nil, "Details Saved Successfully", "")

}

func GetUserDetail(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.LoginStore.GetUserDetail(c.MustGet("user_id").(string))
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	utils.RespondOK(c, result, "Successfully get User Detail", "")
}
