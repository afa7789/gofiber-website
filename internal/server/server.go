package server

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
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

	// static files
	r.Static("/public", "./web/static")

	r.Get("/", server.mainPage())

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
