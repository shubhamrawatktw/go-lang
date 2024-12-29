package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Env         string
	StoragePath string
	Port        string
}

func MustLoad() *Config {

	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading env file : %v", err)
	}
	port := os.Getenv("PORT")
	storagePath := os.Getenv("STORAGE_PATH")
	env := os.Getenv("ENV")

	cfg := Config{
		Env:         env,
		StoragePath: storagePath,
		Port:        port,
	}

	fmt.Println(port, storagePath, env)

	return &cfg
}
