package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"

	"news_api/config"
	"news_api/database"
	"news_api/router"
)

func main() {
	config, err := config.LoadConfig("./")
	if err != nil {
		log.Fatalln("APP: Config not found")
	}

	err = database.Connect()
	if err != nil {
		log.Println(err.Error())
	}
	defer database.Close()

	app := fiber.New()

	router.SetupRoutes(app)

	log.Println(app.Listen("localhost:" + strconv.Itoa(config.AppPort)))
}
