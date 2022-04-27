package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// mainPage creates a mainPage template
func (s *Server) mainPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).Render("index.html", fiber.Map{
			"Title":      "afa7789 - Computer Engineering Solutions",
			"MainHeader": true,
		})
	}
}

// mainPage creates a mainPage template
func (s *Server) thanksPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).Render("thanks.html", fiber.Map{
			"Title": "Thanks for your contact - afa7789",
		})
	}
}
