package server

import (
	"html/template"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// blogEditor opens the template
// this func returns a page to edit an old post
// or create a newer one
func (s *Server) blogEditor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		post_id := c.Params("post_id")
		if post_id != "" {
			// retrieve post data
			return c.Status(http.StatusOK).Render("editor.html", fiber.Map{
				"Title":       "Post Editor - " + post_id + " - afa7789 ",
				"PostID":      post_id,
				"PostTitle":   "teste",
				"PostContent": template.HTML("<b>teste</b>"),
			})
		} else {
			return c.Status(http.StatusOK).Render("editor.html", fiber.Map{
				"Title": "Post Creator - afa7789 ",
			})
		}

	}
}

// blogView opens the template
// this func returns the blog page
// or specific post
func (s *Server) blogView() fiber.Handler {
	return func(c *fiber.Ctx) error {
		post_id := c.Params("post_id")
		post_title := "Título"
		if post_id != "" {
			// retrieve post data
			// blog post
			return c.Status(http.StatusOK).Render("post.html", fiber.Map{
				"Title":       "Post - " + post_id + " - " + post_title + " - afa7789 ",
				"PostID":      post_id,
				"PostTitle":   post_title,
				"PostContent": template.HTML("<b>teste</b>"),
			})
		} else {
			// blog posts
			return c.Status(http.StatusOK).Render("blog.html", fiber.Map{
				"Title": "Blog - afa7789 ",
			})
		}

	}
}

// blogMissing opens the template
// this func returns the missing blog post page
// or specific post
func (s *Server) blogMissing() fiber.Handler {
	return func(c *fiber.Ctx) error {

		return c.Status(http.StatusOK).Render("post.html", fiber.Map{
			"Title": "Post doesn't exist - afa7789 ",
		})

	}
}
