package main

import (
	routes "PerkHub/Routes"
	"PerkHub/connection"
	"PerkHub/constants"
	"PerkHub/settings"
	"PerkHub/stores"
	"context"
	"fmt"

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

	db, err := connection.MakePotgressConn()

	if err != nil {
		fmt.Println(err.Error())
	}

	defer db.Close()
	store := stores.NewStores(db)

	app.Use(store.BindStore())
	routes.Endpoints(app)

	app.Run(fmt.Sprintf("localhost:%d", constants.Port))

}
