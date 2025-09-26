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

func getJWT_KEY() string {
	return settings.Config("JWT_KEY")
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

func getBaseUrl() string {
	return settings.Config("BASE_URL")
}

func getAdminBaseUrl() string {
	return settings.Config("ADMIN_BASE_URL")
}

func getImageBaseUrl() string {
	return settings.Config("IMAGE_BASE_URL")
}

func getFast2SMSKey() string {
	return settings.Config("FAST2SMS_API_KEY")
}

func getFirebaseProjectID() string {
	return settings.Config("FCM_PROJECT_ID")
}

func getFirebaseFilePath() string {
	return settings.Config("FCM_FILE_PATH")
}

func getGameCron() string {
	return settings.Config("GAME_CRON")
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
	JWT_KEY            = getJWT_KEY()
	BASE_URL           = getBaseUrl()
	ADMIN_BASE_URL     = getAdminBaseUrl()
	IMAGE_BASE_URL     = getImageBaseUrl()
	FAST2SMS_API_KEY   = getFast2SMSKey()
	FirebaseProjectID  = getFirebaseProjectID()
	FireBaseFilePath   = getFirebaseFilePath()
	GameCron           = getGameCron()
)

var StatusMap = map[string]string{
	"pending":   "pending",
	"payable":   "verified",
	"validated": "verified",
	"rejected":  "rejected",
}
