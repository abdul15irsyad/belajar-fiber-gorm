package v1

import (
	"abdul15irsyad/belajar-fiber-gorm/handlers"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(v1 fiber.Router) {
	v1.Route("/user", func(user fiber.Router) {
		user.Get("/", handlers.GetUsersHandler)
		user.Get("/:id", handlers.GetUserHandler)
		user.Post("/", handlers.PostUserHandler)
	}, "user.")
}
