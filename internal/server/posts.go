package server

import (
	"afa7789/site/internal/domain"
	"io"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

type PostsController struct{}

func NewPostsController() *PostsController {
	return &PostsController{}
}

func (pc *PostsController) ReceivePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post *domain.Post

		if err := c.BodyParser(&post); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(struct {
				Message string `json:"message"`
			}{
				Message: "Error at content parsing",
			})
		}

		file, err := c.FormFile("document")
		if !(err == nil) {
			// store the file
			f, err := os.OpenFile(domain.StaticPathToFile, os.O_WRONLY|os.O_CREATE, 0o666)
			if err != nil {
				log.Default().Printf("Error at saving file: %v for Post %v", err, post.ID)
			}
			defer f.Close()

			fio, _ := file.Open()
			_, err = io.Copy(f, fio)
			if err != nil {
				log.Default().Printf("Error at saving file: %v for Post %v", err, post.ID)
			}

			// and add it to the post
			post.Image = domain.PathToFile + file.Filename
		}

		// store the data TODO, receive post
		// upload or create the post
		// if it's create will be sent with post id 0.

		return c.Status(fiber.StatusOK).JSON(post)
	}
}
