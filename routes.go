package jeanservices

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type Body struct {
	Name	string
	Data	interface{}
}

func RegisterRoutes(server *fiber.App, client *mongo.Client) {
	server.Use("/api/v1", func(c *fiber.Ctx) error {
		c.Accepts("application/json")
		return c.Next()
	})

	server.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": 200,
			"message": "ok",
			"data": nil,
			"exited_code": 0,
		})
	})

	server.Post("/api/v1/user/create", func(c *fiber.Ctx) error {
		err := c.BodyParser(&Body{})
		if err != nil {
			panic(err)
		}

		fmt.Println(c.Body())

		return c.JSON(fiber.Map{
			"status": 200,
			"message": "ok",
			"data": nil,
			"exited_code": 0,
		})
	})
}
