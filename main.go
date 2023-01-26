package main

import (
	"github.com/MicBun/go-file-manager/config"
	"github.com/MicBun/go-file-manager/route"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default env")
	}

	database := config.ConnectDataBase()
	sqlDB, _ := database.DB()
	defer sqlDB.Close()
	r := route.SetupRouter(database)
	r.Run()
}
