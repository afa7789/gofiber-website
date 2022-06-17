package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) demoBlockiesPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).Render("vue.html", fiber.Map{
			"Title": "BlockiesVue – Demo",
		})
	}
}
