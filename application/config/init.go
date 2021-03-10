package config

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var Env struct {
	APP_URL string
}

func Init() {
	loadEnv()
	InitRoute()
	log.Println("Starting server:", Env.APP_URL)
	err := http.ListenAndServe(Env.APP_URL, nil)
	if err != nil {
		log.Println("Can not start server, failed: ", err)
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
	Env.APP_URL = os.Getenv("APP_URL")
}
