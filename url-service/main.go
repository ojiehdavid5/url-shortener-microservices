package main

import (
	"math/rand"
	"time"

	"github.com/gofiber/fiber/v2"
)

type URL struct {
	Code        string `json:"code"`
	OriginalURL string `json:"original_url"`
}

var urls = make(map[string]string)

func generateCode() string {
	const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())
	code := make([]byte, 6)
	for i := range code {
		code[i] = chars[rand.Intn(len(chars))]
	}
	return string(code)
}

func main() {
	app := fiber.New()

	app.Post("/urls", func(c *fiber.Ctx) error {
		var req struct {
			OriginalURL string `json:"original_url"`
		}
		if err := c.BodyParser(&req); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		code := generateCode()
		urls[code] = req.OriginalURL

		return c.JSON(URL{Code: code, OriginalURL: req.OriginalURL})
	})

	app.Get("/urls/:code", func(c *fiber.Ctx) error {
		code := c.Params("code")
		url, ok := urls[code]
		if !ok {
			return c.Status(404).SendString("Not found")
		}
		return c.JSON(fiber.Map{"original_url": url})
	})

	app.Listen(":8081")
}
