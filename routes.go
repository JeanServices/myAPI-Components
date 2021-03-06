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
	
	server.Get("/api/v1/get/user/:id", func (c *fiber.Ctx) error {
		var data = make(map[string]interface{})
		err := client.Database("myAPI").Collection("users").FindOne(context.TODO(), bson.D{{"id", c.Params("id")}}).Decode(&data)
		if err != nil {
			return c.JSON(fiber.Map{
				"status": 200,
				"message": err.Error(),
				"data": nil,
				"exited_code": 1,
			})
		}

		return c.JSON(fiber.Map{
			"status": 200,
			"message": "obtained",
			"data": data,
			"exited_code": 0,
		})
	})

	server.Post("/api/v1/user/create", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": 200,
			"message": "ok",
			"data": nil,
			"exited_code": 0,
		})
	})
}
