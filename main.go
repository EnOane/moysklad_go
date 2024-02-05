package main

import (
	"github.com/joho/godotenv"
	"log"
	"moysklad/app"
	"moysklad/config"
	"moysklad/infr"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	config.NewConfig()
	infr.ConnectToDb()
	app.CreateApp()
}
