package reglogin

import (
	"PerkHub/middlewear"

	"github.com/gin-gonic/gin"
)

func Routes(api *gin.RouterGroup) {

	app := api.Group("/")
	app.Use(middlewear.UserMiddleware())
	api.POST("/sendOTP", LoginRegistration)
	api.POST("/verifyOTP", VerifyOTP)
	app.POST("/savedetail", SaveUserDetail)
	app.GET("/getUserDetail", GetUserDetail)

}
