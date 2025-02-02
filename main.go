package main

import (
	routes "PerkHub/Routes"
	"PerkHub/connection"
	amazon "PerkHub/connection"
	"PerkHub/constants"
	"PerkHub/settings"
	"PerkHub/stores"
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var (
	appCtx    context.Context
	appCancel context.CancelFunc
)

func main() {
	appCtx, appCancel = context.WithCancel(context.Background())
	settings.LoadEnvFile()
	defer appCancel()

	// Initialize Gin
	app := gin.Default()

	// CORS middleware setup
	app.Use(cors.Default())

	// AWS setup
	awsKeyId := constants.AWSAccessKeyID
	awsSecretKey := constants.AWSSecretAccessKey
	awsRegion := constants.AWSRegion

	// Initialize AWS instance
	awsInstance, err := amazon.NewAws(
		awsRegion,
		awsKeyId,
		awsSecretKey,
		constants.AWSBucketName,
		constants.AWSCloudFrontURL,
	)
	if err != nil {
		fmt.Println(err)
	}

	// CORS configuration for frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},                         // Allow React app's domain
		AllowMethods:     []string{"POST", "GET", "OPTIONS"},                        // Allow POST, GET, OPTIONS
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Request-Id"}, // Allow x-request-id
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Serve static files for Admin assets
	app.Static("/Admin", "./Admin")

	// Parse templates
	// tmpl := template.Must(template.ParseFiles("Admin/index.html", "Admin/component/navbar.html"))
	tmpl, err := template.New("base").ParseFiles(
		// "Admin/login.html",
		"Admin/index.html",
		"Admin/login.html",
		"Admin/Pages/userlist.html",
		"Admin/Pages/MiniApps/miniapp.html",
		"Admin/Pages/MiniApps/AddAndEditMiniApp.html",
		"Admin/component/navbar.html",
	)
	if err != nil {
		log.Fatal("Error parsing templates:", err)
	}

	app.GET("/:page", func(c *gin.Context) {
		// Get the page name from the URL parameter
		page := c.Param("page")

		// Try to render the specific template (e.g. login, about, contact, etc.)
		err := tmpl.ExecuteTemplate(c.Writer, page, nil)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			log.Println("Error rendering template:", err)
		}
	})

	// Connect to the database
	fmt.Println("S3 Connected Successfully")
	db, err := connection.MakePotgressConn()
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()
	store := stores.NewStores(db)

	// Bind the store with AWS instance
	app.Use(store.BindStore(awsInstance))

	// Initialize API routes
	routes.Endpoints(app)

	// Serve static files (like images, CSS, JS)
	app.Static("/files", "./files") // This will map /files/* to the ./files directory

	// Start the Gin server on the specified port
	app.Run(fmt.Sprintf("localhost:%d", constants.Port))
}
