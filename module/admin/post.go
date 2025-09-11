package admin

import (
	"PerkHub/request"
	"PerkHub/stores"
	"PerkHub/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	var login request.AdminLoginRequest
	err = c.ShouldBindJSON(&login)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	result, err := store.AdminStore.AdminLogin(&login)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	res, err := utils.GenerateJWTToken(result.UserID, time.Minute*30)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	data := gin.H{}
	data["user"] = result
	fmt.Println("data:", result)
	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"token":   res,
		"data":    data,
		"message": "Login Successfully",
	})

	// settings.StatusOk(c, result, "Login Successfully", "")
}

func RegisterAdmin(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	var register request.AdminRegister
	err = c.ShouldBindJSON(&register)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}

	_, err = store.AdminStore.AdminRegister(&register)
	if err != nil {
		utils.RespondBadRequest(c, err, "")
		return
	}
	utils.RespondOK(c, nil, "Registration successful", "")
}
