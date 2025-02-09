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
	// Enable CORS for all origins (this allows all websites to access your resources)
	// corsHandler := handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(http.DefaultServeMux)
	// http.ListenAndServe(fmt.Sprintf("localhost:%d", constants.Port), corsHandler) // Start the Gin server on the specified port
	app.Run(fmt.Sprintf("localhost:%d", constants.Port))
}
