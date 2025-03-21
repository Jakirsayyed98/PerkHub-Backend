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

	"github.com/gin-contrib/cors"
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
		// CORS configuration for frontend
		admin.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:5173"},                         // Allow React app's domain
			AllowMethods:     []string{"POST", "GET", "OPTIONS"},                        // Allow POST, GET, OPTIONS
			AllowHeaders:     []string{"Content-Type", "Authorization", "X-Request-Id"}, // Allow x-request-id
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))

		// Serve static files for Admin assets
		admin.Static("/FinalAdmin", "./FinalAdmin")
		admin.Static("/dist", "D:/PerkHub/PerkHub_Go/FinalAdmin/dist")

		// Parse templates
		// tmpl := template.Must(template.ParseFiles("Admin/index.html", "Admin/component/navbar.html"))
		tmpl, err := template.New("base").ParseFiles(
			"FinalAdmin/dist/pages/login/login.html",
			"FinalAdmin/dist/pages/gameslist.html",
			"FinalAdmin/dist/pages/index.html",
			"FinalAdmin/dist/pages/index2.html",
			"FinalAdmin/dist/component/navbar.html",
			"FinalAdmin/dist/component/sidenavbar.html",
		)
		if err != nil {
			log.Fatal("Error parsing templates:", err)
		}

		admin.GET("/:page", func(c *gin.Context) {
			// Get the page name from the URL parameter
			page := c.Param("page")

			// Try to render the specific template (e.g. login, about, contact, etc.)
			err := tmpl.ExecuteTemplate(c.Writer, page, nil)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				log.Println("Error rendering template:", err)
			}
		})

	}

}
