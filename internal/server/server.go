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

	engine.AddFunc("add", func(a, b int) int {
		return a + b
	})

	engine.AddFunc("sub", func(a, b int) int {
		return a - b
	})

	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false
	r := fiber.New(fiber.Config{
		Views:             engine,
		EnablePrintRoutes: false,
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

	// Some pages
	r.Get("/", server.mainPage())
	r.Get("/ltda", server.enterpriseMainPage())

	r.Get("/thanks", server.thanksPage())
	r.Get("/failed", server.failedPage())
	r.Get("/profile", server.githubPage())
	r.Get("/github", server.githubPage())
	r.Get("/vue", server.demoBlockiesPage())
	r.Get("/blockies-vue-demo", server.demoBlockiesPage())
	r.Get("/gradient-demo", server.demoGradientPage())

	message := r.Group("/message")
	message.Get("/",
		basicauth.New(bac), // Basic Auth Middleware
		server.messagesView())

	message.Delete("delete/:id", server.deleteMessage())

	// links
	link := r.Group("/link")
	link.Get("edit/:link_id?",
		basicauth.New(bac), // Basic Auth Middleware
		server.linkEditor())

	link.Delete("delete/:link_id?", server.deleteLink())
	link.Get("swap/:source_id/:target_id", server.swapIndexes())

	link.Post("/", server.receiveLink())
	// link view
	link.Get("/", server.linksView())

	blog := r.Group("/blog")
	// editor is exclusive to admin authentification
	blog.Get("edit/:post_id?",
		basicauth.New(bac), // Basic Auth Middleware
		server.blogEditor())

	// missing page
	blog.Get("/missing", server.blogMissing())
	// post view
	blog.Get("/:post_id", server.postView())
	// blog view
	blog.Get("/", server.blogView())

	// Post Auth Middleware ?
	pc := newPostsController(si.Reps.PostRep)
	blog.Post("/post", pc.receivePost())

	// Mail routes
	mailController := newMailerController(server.reps.MessageRep)
	r.Post("/mail", mailController.send())

	server.router = r
	return server
}

// StartTLS starts the server with TLS connection
func (s *Server) StartTLS(port int) {

	err := s.router.ListenTLS(
		fmt.Sprintf(":%d", port),
		os.Getenv("CERTIFICATE"),
		os.Getenv("PRIVATE_KEY"),
	)

	if err != nil {
		// Using this error treatment to try again on next port
		if strings.Contains(err.Error(), "address already in use") {
			fmt.Println("")
			log.Printf("PORT ALREADY IN USE::%d", port)
			port++
			log.Printf("TRYING NEXT PORT:%d\n", port)
			s.Start(port)
		} else {
			fmt.Printf("%s\n", err)
			panic(err)
		}
	}
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
