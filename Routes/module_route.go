package routes

import (
	"PerkHub/module/admin"
	"PerkHub/module/affiliates"
	"PerkHub/module/banner"
	"PerkHub/module/category"
	"PerkHub/module/games"
	miniapp "PerkHub/module/miniApp"
	"PerkHub/module/mobile"
	"PerkHub/module/notification"
	offers "PerkHub/module/offers"
	reglogin "PerkHub/module/reg_login"
	"PerkHub/module/tickets"
	"PerkHub/module/transactions"
	"PerkHub/module/withdrawal"

	"github.com/gin-gonic/gin"
)

func Endpoints(app *gin.Engine) {
	// ---------- API ROUTES ----------
	api := app.Group("/api")
	{
		admin.Routes(api)
		reglogin.Routes(api)
		category.Routes(api)
		miniapp.Routes(api)
		banner.Routes(api)
		mobile.Routes(api)
		affiliates.Routes(api)
		transactions.Routes(api)
		games.Routes(api)
		withdrawal.Routes(api)
		tickets.RegisterRoutes(api)
		offers.Routes(api)
		notification.Routes(api)
	}
}
