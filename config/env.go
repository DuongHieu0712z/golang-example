package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	MONGO_URI = "MONGO_URI"
	DB_NAME   = "DB_NAME"
	PORT      = "PORT"
)

func getEnv(key string) string {
	if err := godotenv.Load(); err != nil {
		log.Println(err)
	}
	return os.Getenv(key)
}

func GetMongoURI() string {
	return getEnv(MONGO_URI)
}

func GetDbName() string {
	return getEnv(DB_NAME)
}

func GetPort() string {
	return getEnv(PORT)
}
