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

func getAwsRegion() string {
	return settings.Config("AWS_REGION")
}

func getAwsAccessID() string {
	return settings.Config("AWS_ACCESS_KEY_ID")
}
func getAwsSecretAccessKey() string {
	return settings.Config("AWS_SECRET_ACCESS_KEY")
}

func getAwsBucketName() string {
	return settings.Config("AWS_BUCKET_NAME")
}

func getAWSCloudFrontURL() string {
	return settings.Config("AWS_CLOUDFRONT_URL")
}

var (
	Port               = GetPortNumber()
	PostgresHost       = getPostgresHost()
	PostgresPort       = getPostgresPort()
	PostgresUsername   = getPostgresUserName()
	PostgresPassword   = getPostgresPassword()
	PostgresDatabase   = getPostgresDatabase()
	AWSAccessKeyID     = getAwsAccessID()
	AWSSecretAccessKey = getAwsSecretAccessKey()
	AWSRegion          = getAwsRegion()
	AWSBucketName      = getAwsBucketName()
	AWSCloudFrontURL   = getAWSCloudFrontURL()
)
