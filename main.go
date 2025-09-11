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
	"net/http"
	"net/url"

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
	app.Use(CORSMiddleware())

	// Debug logging middleware
	app.Use(func(c *gin.Context) {
		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

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
		fmt.Printf("Failed to initialize AWS: %v\n", err)
		return
	}

	// Connect to the database
	db, err := connection.MakePotgressConn()
	if err != nil {
		fmt.Printf("Failed to connect to database: %v\n", err)
		return
	}
	defer db.Close()

	store := stores.NewStores(db)

	// Bind the store with AWS instance
	app.Use(store.BindStore(awsInstance))

	// Initialize API routes
	routes.Endpoints(app)

	// Example redirect route
	app.GET("/r", RedirectHandler)

	// Serve static files (like images, CSS, JS)

	app.Static("/files", "./files")

	// Start Gin server
	if err := app.Run(fmt.Sprintf("localhost:%d", constants.Port)); err != nil {
		fmt.Printf("Failed to start server: %v\n", err)
	}
}

func RedirectHandler(c *gin.Context) {
	raw := c.Query("u")
	if raw == "" {
		c.String(http.StatusBadRequest, "missing u param")
		return
	}

	decoded, err := url.QueryUnescape(raw)
	if err != nil {
		c.String(http.StatusBadRequest, "invalid url encoding")
		return
	}

	c.Redirect(http.StatusFound, decoded)
}
func CORSMiddleware() gin.HandlerFunc {
	allowedOrigins := map[string]bool{
		"http://localhost:5173":                        true,
		"https://blessed-pretty-mammal.ngrok-free.app": true,
	}

	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if allowedOrigins[origin] {
			c.Header("Access-Control-Allow-Origin", origin)
			c.Header("Vary", "Origin")
			c.Header("Access-Control-Allow-Credentials", "true")
			c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
			c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		}

		// ✅ Preflight requests must return early
		if c.Request.Method == http.MethodOptions {
			c.AbortWithStatus(http.StatusNoContent) // 204
			return
		}

		c.Next()
	}
}
