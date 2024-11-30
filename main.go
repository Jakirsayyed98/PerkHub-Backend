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

	app := gin.Default()

	awsKeyId := constants.AWSAccessKeyID
	awsSecretKey := constants.AWSSecretAccessKey
	awsRegion := constants.AWSRegion

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
	fmt.Println("S3 Connected Successfully")
	db, err := connection.MakePotgressConn()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()
	store := stores.NewStores(db)

	app.Use(store.BindStore(awsInstance))
	routes.Endpoints(app)

	// Serve the files directory as a static file server
	fs := http.FileServer(http.Dir("./files"))

	// Define the route for serving files
	// This will map /files/* to the files directory in your project
	http.Handle("/files/", http.StripPrefix("/files/", fs))

	app.Run(fmt.Sprintf("localhost:%d", constants.Port))

}
