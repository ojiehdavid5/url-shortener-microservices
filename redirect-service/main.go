package main

import (
	"encoding/json"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/:code", func(c *fiber.Ctx) error {
		code := c.Params("code")

		resp, err := http.Get("http://localhost:8081/urls/" + code)
		if err != nil || resp.StatusCode != 200 {
			return c.Status(404).SendString("Invalid URL")
		}

		var data map[string]string
		json.NewDecoder(resp.Body).Decode(&data)

		return c.Redirect(data["original_url"], 302)
	})

	app.Listen(":8082")
}
