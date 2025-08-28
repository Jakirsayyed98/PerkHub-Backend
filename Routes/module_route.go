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
	"time"

	"github.com/gin-contrib/cors"
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
	}

	// ---------- ADMIN PANEL ROUTES ----------
	adminGroup := app.Group("/admin")
	{
		// Apply CORS (best practice: allow only your production domain later)
		adminGroup.Use(cors.New(cors.Config{
			AllowOrigins:     []string{"http://localhost:4215"}, // React/Vite dev URL
			AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
			AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
			MaxAge:           12 * time.Hour,
		}))

		// // Serve static files (CSS, JS, Images)
		// adminGroup.Static("/assets", "./FinalAdmin/dist/assets")

		// // Pre-parse templates (ensures error at startup, not runtime)
		// tmpl, err := template.ParseFiles(
		// 	"FinalAdmin/dist/pages/login/login.html",
		// 	"FinalAdmin/dist/pages/gameslist.html",
		// 	"FinalAdmin/dist/pages/miniapp.html",
		// 	"FinalAdmin/dist/pages/miniapp_categories.html",
		// 	"FinalAdmin/dist/pages/banner_category_list.html",
		// 	"FinalAdmin/dist/pages/banner_list.html",
		// 	"FinalAdmin/dist/pages/add_update_banner.html",
		// 	"FinalAdmin/dist/pages/affiliate_transaction_list.html",
		// 	"FinalAdmin/dist/pages/index.html",
		// 	"FinalAdmin/dist/pages/add_update_miniApp.html",
		// 	"FinalAdmin/dist/pages/add_update_miniApp_categories.html",
		// 	"FinalAdmin/dist/component/navbar.html",
		// 	"FinalAdmin/dist/component/sidenavbar.html",
		// )
		// if err != nil {
		// 	log.Fatal("Error parsing templates:", err)
		// }

		// // Dynamic page renderer
		// adminGroup.GET("/:page", func(c *gin.Context) {
		// 	page := c.Param("page") + ".html" // auto append .html for cleaner URLs
		// 	err := tmpl.ExecuteTemplate(c.Writer, page, nil)
		// 	if err != nil {
		// 		log.Printf("Error rendering template (%s): %v", page, err)
		// 		c.JSON(http.StatusNotFound, gin.H{"error": "Page not found"})
		// 	}
		// })

		// // Default: redirect /admin → /admin/index
		// adminGroup.GET("/", func(c *gin.Context) {
		// 	c.Redirect(http.StatusFound, "/admin/index")
		// })
	}
}
