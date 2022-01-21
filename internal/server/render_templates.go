package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Template creates a template
func (s *Server) mainPage() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.HTML(http.StatusOK, "index.html", gin.H{
			"Title": "afa7789 - Computer Engineering Solutions",
		})
	}
}
