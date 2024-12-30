package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	PORT           string
	MONGODB_URL    string
	ALLOWED_EMAILS string
}

var Env EnvConfig

func LoadEnv() {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading env file : %v", err)
	}

	port := os.Getenv("PORT")
	mongoUrl := os.Getenv("MONGODB_URL")
	allowedEmails := os.Getenv("ALLOWED_EMAILS")

	Env = EnvConfig{
		PORT:           port,
		MONGODB_URL:    mongoUrl,
		ALLOWED_EMAILS: allowedEmails,
	}

}
