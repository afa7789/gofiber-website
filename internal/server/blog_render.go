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
		potsID := c.Params("post_id")
		if potsID != "" {
			// retrieve post data
			return c.Status(http.StatusOK).Render("editor.html", fiber.Map{
				"Title":       "Post Editor - " + potsID + " - afa7789 ",
				"PostID":      potsID,
				"PostTitle":   "teste",
				"PostContent": template.HTML("<b>teste</b>"),
			})
		}

		return c.Status(http.StatusOK).Render("editor.html", fiber.Map{
			"Title": "Post Creator - afa7789 ",
		})
	}
}

// blogView opens the template
// this func returns the blog page
// or specific post
func (s *Server) blogView() fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID := c.Params("post_id")
		// get post data
		postTitle := "Título"
		if postID != "" {
			// retrieve post data
			// blog post
			return c.Status(http.StatusOK).Render("post.html", fiber.Map{
				"Title":       "Post - " + postID + " - " + postTitle + " - afa7789 ",
				"PostID":      postID,
				"PostTitle":   postTitle,
				"PostContent": template.HTML("<b>teste</b>"),
			})
		}

		// blog posts
		return c.Status(http.StatusOK).Render("blog.html", fiber.Map{
			"Title": "Blog - afa7789 ",
		})
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
