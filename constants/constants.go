package constants

import (
	"PerkHub/settings"
	"strconv"
)

func GetPortNumber() int {
	portnumber := settings.Config("PORT")
	golangPort, _ := strconv.Atoi(portnumber)
	return golangPort
}

func getPostgresHost() string {
	return settings.Config("POSTGRES_HOST")
}

func getPostgresPort() int {
	postgresPORT := settings.Config("POSTGRES_PORT")

	port, _ := strconv.Atoi(postgresPORT)

	return port
}

func getPostgresUserName() string {
	return settings.Config("POSTGRES_USER_NAME")
}

func getPostgresPassword() string {
	return settings.Config("POSTGRES_PASSWORD")
}
func getPostgresDatabase() string {
	return settings.Config("POSTGRES_DATABASE")
}

var (
	Port             = GetPortNumber()
	PostgresHost     = getPostgresHost()
	PostgresPort     = getPostgresPort()
	PostgresUsername = getPostgresUserName()
	PostgresPassword = getPostgresPassword()
	PostgresDatabase = getPostgresDatabase()
)
