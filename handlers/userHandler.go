package handlers

import (
	"abdul15irsyad/belajar-fiber-gorm/database"
	"abdul15irsyad/belajar-fiber-gorm/models/entity"
	"abdul15irsyad/belajar-fiber-gorm/validators"

	"github.com/gofiber/fiber/v2"
)

func GetUsersHandler(c *fiber.Ctx) error {
	users := []entity.User{}

	err := database.DB.Find(&users).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err,
		})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "get all users",
		"data": users,
	})
}

func GetUserHandler(c *fiber.Ctx) error {
	id := c.Params("id")

	user := entity.User{}

	err := database.DB.First(&user, "id = ?", id).Error
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err,
		})
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "get user",
		"data": user,
	})
}

func PostUserHandler(c *fiber.Ctx) error {
	user := new(entity.User)

	err := c.BodyParser(user)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg": err,
		})
		return err
	}

	errors := validators.ValidateStruct(*user)
	if errors != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":    "errors validation",
			"errors": errors,
		})
	}

	newUser := entity.User{
		Name:    user.Name,
		Email:   user.Email,
		Age:     user.Age,
		Address: user.Address,
	}

	err = database.DB.Create(&newUser).Error
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "get user",
		"data": user,
	})
}
