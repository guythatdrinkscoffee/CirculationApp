package config

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"regexp"
)

type Config struct {
	APP_KEY  string
	GIN_MODE string
}

const dirName = "CirculationApp"

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + dirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)

	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func GetConfig() *Config {
	loadEnv()

	appKey := os.Getenv("APP_KEY")
	ginMode := os.Getenv("GIN_MODE")
	return &Config{
		APP_KEY:  appKey,
		GIN_MODE: ginMode,
	}
}
