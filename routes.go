package jeanservices

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(server *fiber.App, client *mongo.Client) {
	server.Use("/api/v1", func(c *fiber.Ctx) error {
		c.Accepts("application/json")

	   return c.Next()
   })

   server.Get("/api/v1", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"status": 200,
			"message": "ok",
			"exited_code": 0,
		})
   })

   server.Post("/api/v1/create", func(c *fiber.Ctx) error {
	   insertUser(client, "352905372375823", "a")

	   fmt.Println(c.Body())

	   return c.JSON(fiber.Map{
		   "jean": "jean",
	   })
   })
}
