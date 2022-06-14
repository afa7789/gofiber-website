package server

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gomarkdown/markdown"
)

func (s *Server) githubPage() fiber.Handler {
	return func(c *fiber.Ctx) error {
		username := os.Getenv("GITHUB_NAME")
		url := fmt.Sprintf(
			"https://raw.githubusercontent.com/%s/%s/main/README.md",
			username, username,
		)

		// request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error at get: %s", err)
		}

		// resp body bytes
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error at read all: %s", err)
		}

		// parse mrkdown to html
		mkd := markdown.ToHTML(body, nil, nil)

		// parse html to template
		pc := template.HTML(string(mkd))

		return c.Status(http.StatusOK).Render("github.html", fiber.Map{
			"Title":          "Profile - afa7789",
			"ProfileContent": pc,
			"Link":           "github.com/" + username,
		})
	}
}
