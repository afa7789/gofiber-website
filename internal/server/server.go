package server

import (
	"afa7789/site/internal/domain"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/template/html"
)

// Server is the definition of a REST server based on Gin
type Server struct {
	router *fiber.App
	reps   *domain.Repositories
}

// New returns a new server that takes advantage of zerolog for logging
// and holds a reference to the app configuration
func New(si *domain.ServerInput) *Server {
	server := &Server{
		reps: si.Reps,
	}

	// https://github.com/gofiber/template
	engine := html.New("./web/templates", "")

	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false
	r := fiber.New(fiber.Config{
		Views: engine,
	})

	// Basic Auth configuration
	bac := basicauth.Config{
		Users: map[string]string{
			os.Getenv("ADMIN_LOGIN"): os.Getenv("ADMIN_PASSWORD"),
		},
	}

	// ================ROUTES====================
	// Static Files
	r.Static("/public", "./web/static")
	r.Get("/", server.mainPage())
	r.Get("/thanks", server.thanksPage())
	r.Get("/failed", server.failedPage())

	blog := r.Group("/blog")
	// editor is exclusive to admin authentification
	blog.Get("edit/:post_id?",
		basicauth.New(bac), // Basic Auth Middleware
		server.blogEditor())
	// missing page
	blog.Get("/missing", server.blogMissing())
	// post viewer or blog list
	blog.Get("/:post_id?", server.blogView())

	// Post Auth Middleware ?
	pc := NewPostsController(&si.Reps.PostRep)
	blog.Post("/post", pc.ReceivePost())

	// Mail routes
	mailController := NewMailerController()
	r.Post("/mail", mailController.send())

	server.router = r

	return server
}

// Start starts the server
func (s *Server) Start(port int) {
	err := s.router.Listen(fmt.Sprintf(":%d", port))
	if err != nil {
		// Using this error treatment to try again on next port
		if strings.Contains(err.Error(), "address already in use") {
			fmt.Println("")
			log.Printf("PORT ALREADY IN USE::%d", port)
			port++
			log.Printf("TRYING NEXT PORT:%d\n", port)
			s.Start(port)
		} else {
			panic(err)
		}
	}
}
