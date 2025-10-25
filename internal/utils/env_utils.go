package utils

import (
	"basics/internal/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

// var AppConfig models.EnvModels

var Cfg models.EnvModels

func LoadEnvConfig() {

	// err := godotenv.Load()

	// if err != nil {
	// 	log.Fatal("Failed During Loading of Env...")
	// }

	//For hosting
	_ = godotenv.Load()

	Cfg = models.EnvModels{
		DatabaseHostName:   os.Getenv("DB_HOST"),
		DatabasePortNumber: os.Getenv("DB_PORT"),
		DatabaseUserName:   os.Getenv("DB_USER"),
		DatabasePassword:   os.Getenv("DB_PASSWORD"),
		DatabaseName:       os.Getenv("DB_NAME"),
		DBSSL:              os.Getenv("DB_SSL"),
		JWTSecret:          os.Getenv("JWT_SECRET"),
	}
	if Cfg.JWTSecret == "" {
		log.Fatal("❌ JWT_SECRET is not set")
	}

}
