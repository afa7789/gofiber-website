package server

import (
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// Server is the definition of a REST server based on Gin
type Server struct {
	router *gin.Engine
}

// New returns a new server that takes advantage of zerolog for logging
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

// Start starts the REST server
func (server *Server) Start(PORT int) {
	fmt.Println("")
	log.Printf("SERVER HERE: http://localhost:%d\n", PORT)
	fmt.Println("")

	err := server.router.Run(fmt.Sprintf(":%d", PORT))
	if err != nil {
		// Using this error treatment to try again on next port
		if strings.Contains(err.Error(), "address already in use") {
			fmt.Println("")
			log.Printf("PORT ALREADY IN USE::%d", PORT)
			PORT++
			log.Printf("TRYING NEXT PORT:%d\n", PORT)
			server.Start(PORT)
		} else {
			panic(err)
		}
	}

}
