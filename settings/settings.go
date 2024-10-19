package settings

import (
	"log"
	"os"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var once = sync.Once{}

func LoadEnvFile() {

	var envFile string
	switch gin.Mode() {
	case "release":
		envFile = ".env.production"
		log.Println("Loaded .env for production environment")
	case "debug":
		envFile = ".env.development"
		log.Println("Loaded .env for debug environment")
	case "test":
		envFile = ".env.test"
		log.Println("Loaded .env for test environment")
	}

	if envFile != "" {
		if err := godotenv.Load(envFile); err != nil {
			log.Panicf("%s not found\n", envFile)
		}
	}

	err := os.Setenv("TZ", os.Getenv("TZ"))
	if err != nil {
		log.Println("Failed to set timezone for the application")
	}
}

// Config load a specified .env.development variable
func Config(key string) string {
	once.Do(LoadEnvFile)
	return os.Getenv(key)
}
