package routes

import (
	"PerkHub/module/affiliates"
	"PerkHub/module/banner"
	"PerkHub/module/category"
	"PerkHub/module/games"
	miniapp "PerkHub/module/miniApp"
	"PerkHub/module/mobile"
	reglogin "PerkHub/module/reg_login"
	"PerkHub/module/transactions"

	"github.com/gin-gonic/gin"
)

func Endpoints(app *gin.Engine) {

	api := app.Group("/api")
	{
		reglogin.Routes(api)
		category.Routes(api)
		miniapp.Routes(api)
		banner.Routes(api)

		mobile.Routes(api)
		affiliates.Routes(api)
		transactions.Routes(api)
		games.Routes(api)

	}

}
