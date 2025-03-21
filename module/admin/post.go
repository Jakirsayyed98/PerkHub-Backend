package admin

import (
	"PerkHub/request"
	"PerkHub/settings"
	"PerkHub/stores"
	"PerkHub/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func AdminLogin(c *gin.Context) {
	store, err := stores.GetStores(c)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
	}

	var login request.AdminLoginRequest
	err = c.ShouldBindJSON(&login)
	if err != nil {
		settings.StatusBadRequest(c, err, "")
		return
	}

	result, err := store.AdminStore.AdminLogin(&login)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	res, err := utils.GenerateJWTToken(result.UserID, time.Minute*10)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}

	data := gin.H{}
	data["user"] = result

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
		settings.StatusBadRequest(c, err.Error(), "")
	}

	var register request.AdminRegister
	err = c.ShouldBindJSON(&register)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  400,
			"error":   err.Error(), // Capture the error message here
			"message": "Validation failed. Please check the input fields.",
		})
		return
	}

	_, err = store.AdminStore.AdminRegister(&register)
	if err != nil {
		settings.StatusBadRequest(c, err.Error(), "")
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Registration successful"})

}
