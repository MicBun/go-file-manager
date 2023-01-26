package main

import (
	"github.com/MicBun/go-file-manager/config"
	"github.com/MicBun/go-file-manager/docs"
	"github.com/MicBun/go-file-manager/route"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file, using default env")
	}

	description := "This is a file manager API.\n\n" +
		"To get Bearer Token, first you need to login. \n\n" +
		"Reset user database by POST /resetUserDatabase \n\n" +
		"Login by POST /login username: user1, password: password1 \n\n" +
		"Checkout my Github: https://github.com/MicBun\n\n" +
		"Checkout my Linkedin: https://www.linkedin.com/in/MicBun\n\n"

	docs.SwaggerInfo.Title = "Go File Manager"
	docs.SwaggerInfo.Description = description

	database := config.ConnectDataBase()
	sqlDB, _ := database.DB()
	defer sqlDB.Close()
	r := route.SetupRouter(database)
	r.Run()
}
