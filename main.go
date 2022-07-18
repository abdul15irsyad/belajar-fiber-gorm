package main

import (
	"abdul15irsyad/belajar-fiber-gorm/database"
	"abdul15irsyad/belajar-fiber-gorm/database/migration"
	"abdul15irsyad/belajar-fiber-gorm/routes"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	// load env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// connect to database
	_, err = database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database")

	// run migration
	migration.RunMigration()

	// fiber init
	app := fiber.New(fiber.Config{
		StrictRouting: true,
		CaseSensitive: true,
	})

	// set public folder to static folder
	app.Static("/", "./public")

	// routes
	routes.Routes(app)

	app.Listen("localhost:3000")
}
