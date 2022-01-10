package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

// Server is the definition of a REST Server based on Gin
type Server struct {
	router *gin.Engine
}

// New returns a new Server that takes advantage of zerolog for logging
// and holds a reference to the app configuration
func New() *Server {
	server := &Server{}

	r := gin.New()

	r.LoadHTMLGlob("templates/*")
	r.Static("/public", "./public")
	// Public urls intended to be accessed from client-side need CORS headers.
	// cs := r.Group("/")
	// corsConfig := cors.DefaultConfig()
	// corsConfig.AllowOrigins = []string{}
	// corsConfig.AllowCredentials = true
	// corsConfig.AddAllowHeaders("x-sessionid")
	// corsConfig.AddAllowHeaders("authorization")
	// cs.Use(cors.New(corsConfig))

	r.GET("/", server.mainPage())

	server.router = r

	return server
}

// Start starts the REST Server
func (server *Server) Start() {
	log.Printf("\n SERVER HERE: http://localhost:8080\n")
	err := server.router.Run()
	if err != nil {
		panic(err)
	}
}
