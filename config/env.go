package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetMongoURI() string {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}

	return os.Getenv("MONGO_URI")
}
