package server

import (
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
}

// New returns a new server that takes advantage of zerolog for logging
// and holds a reference to the app configuration
func New() *Server {
	server := &Server{}

	// https://github.com/gofiber/template
	engine := html.New("./web/templates", "")

	// Reload the templates on each render, good for development
	engine.Reload(true) // Optional. Default: false
	r := fiber.New(fiber.Config{
		Views: engine,
	})

	// Static Files
	r.Static("/public", "./web/static")
	r.Get("/", server.mainPage())

	// mail routes
	mailController := NewMailerController()
	r.Post("/mail", mailController.send())
	// Optional parameter

	r.Get("/hi/:name?", func(c *fiber.Ctx) error {
		name := c.Params("name")
		if name != "" {
			return c.SendString("OI, " + name)
		} else {
			return c.SendString("Quem é você")
		}
	})

	// Basic Auth configuration
	bac := basicauth.Config{
		Users: map[string]string{
			os.Getenv("ADMIN_LOGIN"): os.Getenv("ADMIN_PASSWORD"),
		},
	}
	// Basic Auth Middleware
	rAuth := r.Group("", basicauth.New(bac))
	// editor
	rAuth.Get("/blog-editor/:post_id?", server.blogEditor())

	server.router = r

	return server
}

// Start starts the REST server
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
