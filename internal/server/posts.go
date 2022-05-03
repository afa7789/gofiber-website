package server

import (
	"afa7789/site/internal/domain"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/avelino/slugify"
	"github.com/gofiber/fiber/v2"
	"github.com/gomarkdown/markdown"
)

type PostsController struct {
	pr domain.PostRepository
}

func newPostsController(pr domain.PostRepository) *PostsController {
	return &PostsController{
		pr: pr,
	}
}

/////////////////API ENDPOINTS

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

		if !validatePost(&post) {
			return c.Status(fiber.StatusBadRequest).JSON(post)
		}

		// slugfy the title
		post.Slug = slugify.Slugify(post.Title)

		// getting the file from the form
		file, err := c.FormFile("document")
		if err != nil {
			log.Default().Printf("No document found in request body? %v", err)

		} else {
			// create a file at the place we want to store it
			f, err := os.OpenFile(
				path.Join(domain.StaticPathToFile, file.Filename),
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
		}

		// threat the related posts
		post.RelatedPosts = relatedPostsFixer(post.RelatedPosts)

		// upload or create the post
		// if it's create will be sent with post id 0.
		pc.pr.AddPost(&post)

		return c.Status(fiber.StatusOK).JSON(post)
	}
}

/////////////////RENDER FUNCTIONS

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
				"PostImage":        post.Image,
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
			if err.Error() == "no_slug" {
				return c.Status(fiber.StatusNoContent).Redirect("/blog/" + postID + "-" + post.Slug)
			}
			log.Default().Printf("Error with post ID = %s : %s", postID, err.Error())
			return c.Status(fiber.StatusNoContent).Redirect("/blog/missing")
		}

		// get data from related posts
		splited := strings.Split(post.RelatedPosts, ",")

		// create arrays to be used
		RelatedPostsIDs := []uint{}
		RelatedPostsTitles := []string{}
		RelatedPostsImages := []string{}
		RelatedPostsSynopsies := []string{}
		RelatedPostsSlugs := []string{}

		for _, postIDStr := range splited {
			// check if it's an integer
			if postID, err := strconv.ParseUint(postIDStr, 10, 64); err == nil {
				// adding value.
				RelatedPostsIDs = append(RelatedPostsIDs, uint(postID))
			}
		}
		relatedPosts, err := s.reps.PostRep.RetrievePosts(RelatedPostsIDs)
		if err != nil {
			log.Default().Printf("Error at related post querry : %s", err.Error())
		}

		for _, relatedPost := range relatedPosts {
			RelatedPostsTitles = append(RelatedPostsTitles, relatedPost.Title)
			RelatedPostsImages = append(RelatedPostsImages, relatedPost.Image)
			RelatedPostsSynopsies = append(RelatedPostsSynopsies, relatedPost.Synopsis)
			RelatedPostsSlugs = append(RelatedPostsSlugs, relatedPost.Slug)
		}

		fmt.Printf("\n%+v\n", RelatedPostsSlugs)
		// blog post
		return c.Status(http.StatusOK).Render("post.html", fiber.Map{
			"Title":           post.Title + " - " + postID + " - afa7789 ",
			"PostID":          postID,
			"PostDate":        post.CreatedAt.Format("01-02-2006"),
			"PostImage":       post.Image,
			"PostTitle":       post.Title,
			"PostContent":     template.HTML(string(markdown.ToHTML([]byte(post.Content), nil, nil))),
			"RPostsID":        RelatedPostsIDs,
			"RPostsSlugs":     RelatedPostsSlugs,
			"RPostsImages":    RelatedPostsImages,
			"RPostsTitles":    RelatedPostsTitles,
			"RPostsSynopsies": RelatedPostsSynopsies,
		})

	}
}

func (s *Server) blogView() fiber.Handler {
	return func(c *fiber.Ctx) error {
		// get page
		page := c.Query("page")
		// if the page is empty it will be first page (0)
		pageInt := 0
		var err error
		if page != "" {
			pageInt, err = strconv.Atoi(page)
			if err != nil {
				return c.Status(fiber.StatusNoContent).Redirect("/blog/")
			}
		}

		// mount the stuff to show
		posts, maxPostNumbers := s.reps.PostRep.PageResult(pageInt)
		// number of pages
		numberOfPages := maxPostNumbers / domain.PageLimit

		PostsIds := []uint{}
		PostsTitles := []string{}
		PostsImages := []string{}
		PostsSynopsies := []string{}
		PostsSlugs := []string{}

		for _, p := range posts {
			PostsIds = append(PostsIds, p.ID)
			PostsTitles = append(PostsTitles, p.Title)
			PostsImages = append(PostsImages, p.Image)
			PostsSynopsies = append(PostsSynopsies, p.Synopsis)
			PostsSlugs = append(PostsSlugs, p.Slug)
		}

		// blog posts
		return c.Status(http.StatusOK).Render("blog.html", fiber.Map{
			"Title":          "Blog - afa7789 ",
			"PostsID":        PostsIds,
			"PostsSlugs":     PostsSlugs,
			"PostsImages":    PostsImages,
			"PostsTitles":    PostsTitles,
			"PostsSynopsies": PostsSynopsies,
			"TotalPages":     int(numberOfPages),
			"Page":           pageInt,
		})
	}
}

// blogMissing opens the template
// this func returns the missing blog post page
// or specific post
func (s *Server) blogMissing() fiber.Handler {
	return func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).Render("missing.html", fiber.Map{
			"Title": "Post doesn't exist - afa7789 ",
		})
	}
}

/////////////////HELPER FUNCTIONS

func validatePost(p *domain.Post) bool {
	if p.Content != "" && p.Title != "" && p.Synopsis != "" {
		return true
	}
	return false
}

func (s *Server) getPost(str string) (*domain.Post, error) {

	index := strings.Index(str, "-")
	var postID uint64
	var err error
	var noSlug bool

	// without slug;
	if index == -1 {
		// parse from string to uint
		postID, err = strconv.ParseUint(str, 10, 64)
		noSlug = true
	} else {
		postID, err = strconv.ParseUint(str[:index], 10, 64)
	}

	// if there's an error
	if err != nil {
		if postID != 0 {
			log.Default().Printf("Error parsing post id in = %s.", str)
		}
		return nil, err
	}

	// get post data
	// retrieve post data
	fmt.Printf("Post ID = %d\n", postID)
	post, err := s.reps.PostRep.RetrievePost(uint(postID))
	// fmt.Printf("%+v\n", *post)
	// print("\n")
	if err != nil {
		log.Default().Printf("Couldn't retrieve post ID = %d", postID)
		return nil, err
	}

	if noSlug {
		return post, fmt.Errorf("no_slug")
	}

	return post, nil
}

func relatedPostsFixer(related_posts string) string {
	splited := strings.Split(related_posts, ",")
	concated_result := ""
	for _, post_id := range splited {
		// check if it's an integer
		if _, err := strconv.ParseUint(post_id, 10, 64); err == nil {
			// if not just don't add it
			concated_result += post_id + ","
		}
	}

	if len(concated_result) == 0 {
		return ""
	}

	// remove last comma
	f2 := []rune(concated_result)
	for f2[len(f2)-1] == ',' {
		f2 = f2[:len(f2)-1]
	}
	concated_result = string(f2)

	return concated_result
}
