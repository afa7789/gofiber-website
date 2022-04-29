package server

import (
	"afa7789/site/internal/domain"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PostsController struct {
	pr domain.PostRepository
}

func newPostsController(pr domain.PostRepository) *PostsController {
	return &PostsController{
		pr: pr,
	}
}

// Receive post receives a multi-form from page.
func (pc *PostsController) receivePost() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var post domain.Post

		// parsinsg the post that's in the form coming from the request
		if err := c.BodyParser(&post); err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(struct {
				Message string `json:"message"`
			}{
				Message: "Error at content parsing " + err.Error(),
			})
		}

		// getting the file from the form
		file, err := c.FormFile("document")
		if err != nil {
			log.Default().Printf("No document found in request body? %v", err)
			return c.Status(fiber.StatusOK).JSON(fiber.Map{
				"post": post,
				"err":  err.Error(),
			})
		}

		// create a file at the place we want to store it
		f, err := os.OpenFile(
			filepath.Join(domain.StaticPathToFile, file.Filename),
			os.O_WRONLY|os.O_CREATE, 0o666)
		if err != nil {
			log.Default().Printf("Error at saving file: %v for Post %v", err, post.ID)
		}
		defer f.Close()

		// copying the file bytes to the file we created above
		fio, _ := file.Open()
		_, err = io.Copy(f, fio)
		if err != nil {
			log.Default().Printf("Error at saving file: %v for Post %v", err, post.ID)
		}

		// and add it to the post
		post.Image = domain.PathToFile + file.Filename

		// upload or create the post
		// if it's create will be sent with post id 0.
		pc.pr.AddPost(&post)

		return c.Status(fiber.StatusOK).JSON(post)
	}
}

func (s *Server) getPost(id string) (*domain.Post, error) {
	// parse from string to uint
	ID_int, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		log.Default().Printf("Error parsing post ID = %s, not an integer", id)
		return nil, err
	}
	// get post data
	// retrieve post data
	post, err := s.reps.PostRep.RetrievePost(uint(ID_int))
	if err != nil {
		log.Default().Printf("Couldn't retrieve post ID = %s", id)
		return nil, err
	}
	return post, nil
}

// blogEditor opens the template
// this func returns a page to edit an old post
// or create a newer one
func (s *Server) blogEditor() fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID := c.Params("post_id")
		post, err := s.getPost(postID)
		if err != nil {
			log.Default().Printf("Error with post ID = %s : %s", postID, err.Error())
		}

		if postID != "" {
			// retrieve post data
			return c.Status(http.StatusOK).Render("editor.html", fiber.Map{
				"Title":            "Post Editor - " + postID + " - afa7789 ",
				"PostID":           postID,
				"PostTitle":        post.Title,
				"PostContent":      post.Content,
				"PostSynopsis":     post.Synopsis,
				"PostRelatedPosts": post.RelatedPosts,
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
func (s *Server) postView() fiber.Handler {
	return func(c *fiber.Ctx) error {
		postID := c.Params("post_id")
		// get post data
		post, err := s.getPost(postID)
		if err != nil {
			log.Default().Printf("Error with post ID = %s : %s", postID, err.Error())
		}

		if postID != "" {
			// blog post
			return c.Status(http.StatusOK).Render("post.html", fiber.Map{
				"Title":            post.Title + " - " + postID + " - afa7789 ",
				"PostID":           postID,
				"PostTitle":        post.Title,
				"PostContent":      post.Content,
				"PostSynopsis":     post.Synopsis,
				"PostRelatedPosts": post.RelatedPosts,
			})
		}

		return c.Status(fiber.StatusOK).Redirect("/missing")

	}
}

func (s *Server) blogView() fiber.Handler {
	return func(c *fiber.Ctx) error {
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
