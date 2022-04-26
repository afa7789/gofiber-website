package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// mainPage creates a mainPage template
func (s *Server) mainPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).Render("index.html", fiber.Map{
			"Title":        "afa7789 - Computer Engineering Solutions",
			"SharedHeader": false,
		})
	}
}

// blogEditor opens the template
func (s *Server) blogEditor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		post_id := c.Params("post_id")
		if post_id != "" {
			return c.Status(http.StatusOK).Render("blog_editor.html", fiber.Map{
				"Title":        "Blog Post Editor - " + post_id + " - afa7789 ",
				"SharedHeader": true,
				"PostID":       post_id,
			})
		} else {
			return c.Status(http.StatusOK).Render("blog_editor.html", fiber.Map{
				"Title":        "Blog Post Creator - afa7789 ",
				"SharedHeader": true,
			})
		}

	}
}
