package routes

import (
	"PerkHub/module/admin"
	"PerkHub/module/affiliates"
	"PerkHub/module/banner"
	"PerkHub/module/category"
	"PerkHub/module/games"
	miniapp "PerkHub/module/miniApp"
	"PerkHub/module/mobile"
	reglogin "PerkHub/module/reg_login"
	"PerkHub/module/transactions"
	"PerkHub/module/withdrawal"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Endpoints(app *gin.Engine) {

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

	}
	admin := app.Group("/admin")
	{

		// Serve static files from the "dist" directory
		admin.Static("/dist", "./FinalAdmin/dist")

		// Parse templates
		tmpl, err := template.New("base").ParseFiles(
			"FinalAdmin/dist/pages/login/login.html",
			"FinalAdmin/dist/pages/gameslist.html",
			"FinalAdmin/dist/pages/miniapp.html",
			"FinalAdmin/dist/pages/index.html",
			"FinalAdmin/dist/component/navbar.html",
			"FinalAdmin/dist/component/sidenavbar.html",
		)
		if err != nil {
			log.Fatal("Error parsing templates:", err)
		}

		// Render HTML pages dynamically
		admin.GET("/:page", func(c *gin.Context) {
			page := c.Param("page")
			err := tmpl.ExecuteTemplate(c.Writer, page, nil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Println("Error rendering template:", err)
			}
		})
	}

}
