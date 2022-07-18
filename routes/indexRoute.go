package routes

import (
	"abdul15irsyad/belajar-fiber-gorm/handlers"
	v1 "abdul15irsyad/belajar-fiber-gorm/routes/v1"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	apiRouter := app.Group("/api")

	apiRouter.Get("/", handlers.RootHandler)

	v1Router := apiRouter.Group("/v1")
	v1.UserRoute(v1Router)
}
