package server

import (
	"afa7789/site/internal/domain"
	"io"
	"log"
	"mime/multipart"
	"os"

	"github.com/gofiber/fiber/v2"
)

type PostsController struct{}

func NewPostsController() *PostsController {
	return &PostsController{}
}

// Receive post receives a multi-form from page.
func (pc *PostsController) ReceivePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post domain.Post

		// a := c.Body()
		// print(string(a))
		if err := c.BodyParser(&post); err != nil {

			return c.Status(fiber.StatusInternalServerError).JSON(struct {
				Message string `json:"message"`
			}{
				Message: "Error at content parsing " + err.Error(),
			})
		}

		file, err := c.FormFile("document")

		if err != nil {
			log.Default().Printf("No document found in request body %v", err)
			return c.Status(fiber.StatusOK).JSON( fiber.{post )
		} else {
			go func(file *multipart.FileHeader) {
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

				// store the data TODO, receive post
				// upload or create the post
				// if it's create will be sent with post id 0.
			}(file)
		}

		return c.Status(fiber.StatusOK).JSON(post)
	}
}
