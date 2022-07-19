package server

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

// mainPage creates a mainPage template
func (s *Server) mainPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var LastPostsTitles []string
		var LastPostsImages []string
		var LastPostsSynopsies []string
		var LastPostsIds []uint
		var LastPostsSlugs []string

		posts, _ := s.reps.PostRep.LastThreePosts()
		for _, post := range posts {
			// check if it's an integer
			LastPostsTitles = append(LastPostsTitles, post.Title)
			LastPostsImages = append(LastPostsImages, post.Image)
			LastPostsSynopsies = append(LastPostsSynopsies, post.Synopsis)
			LastPostsIds = append(LastPostsIds, post.ID)
			LastPostsSlugs = append(LastPostsSlugs, post.Slug)
		}

		return c.Status(http.StatusOK).Render("index_blog.html", fiber.Map{
			"Title":              "afa7789 - Computer Wizzard Tech Blog",
			"MainHeader":         true,
			"LastPostsTitles":    LastPostsTitles,
			"LastPostsSynopsies": LastPostsSynopsies,
			"LastPostsImages":    LastPostsImages,
			"LastPostsSlugs":     LastPostsSlugs,
			"LastPostsIds":       LastPostsIds,
		})
	}
}

// mainPage creates a mainPage template
func (s *Server) enterpriseMainPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var LastPostsTitles []string
		var LastPostsImages []string
		var LastPostsSynopsies []string
		var LastPostsIds []uint
		var LastPostsSlugs []string

		posts, _ := s.reps.PostRep.LastThreePosts()
		for _, post := range posts {
			// check if it's an integer
			LastPostsTitles = append(LastPostsTitles, post.Title)
			LastPostsImages = append(LastPostsImages, post.Image)
			LastPostsSynopsies = append(LastPostsSynopsies, post.Synopsis)
			LastPostsIds = append(LastPostsIds, post.ID)
			LastPostsSlugs = append(LastPostsSlugs, post.Slug)
		}

		return c.Status(http.StatusOK).Render("index.html", fiber.Map{
			"Title":              "afa7789 - Computer Engineering Solutions",
			"MainHeader":         true,
			"LastPostsTitles":    LastPostsTitles,
			"LastPostsSynopsies": LastPostsSynopsies,
			"LastPostsImages":    LastPostsImages,
			"LastPostsSlugs":     LastPostsSlugs,
			"LastPostsIds":       LastPostsIds,
		})
	}
}

// thanksPage creates a thanksPage template
func (s *Server) thanksPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).Render("thanks.html", fiber.Map{
			"Title": "Thanks for your contact - afa7789",
		})
	}
}

// failedPage creates a failedPage template
func (s *Server) failedPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).Render("failed.html", fiber.Map{
			"Title": "Contact failed - afa7789",
		})
	}
}
