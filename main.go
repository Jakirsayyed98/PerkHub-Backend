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
	"net/url"
	"time"

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

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // your React dev server
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	app.Use(func(c *gin.Context) {
		println("Request:", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})
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

	// Connect to the database
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
	app.GET("/r", RedirectHandler)

	// Serve static files (like images, CSS, JS)
	app.Static("/files", "./files") // This will map /files/* to the ./files directory
	// Enable CORS for all origins (this allows all websites to access your resources)
	// corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(http.DefaultServeMux)
	// http.ListenAndServe(fmt.Sprintf("localhost:%d", constants.Port), corsHandler) // Start the Gin server on the specified port
	app.Run(fmt.Sprintf(":%d", constants.Port))
}

func RedirectHandler(c *gin.Context) {
	raw := c.Query("u")
	if raw == "" {
		c.String(400, "missing u param")
		return
	}

	decoded, err := url.QueryUnescape(raw)
	if err != nil {
		c.String(400, "invalid url encoding")
		return
	}

	c.Redirect(302, decoded)
}
