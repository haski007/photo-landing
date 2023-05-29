package server

import (
	"github.com/haski007/photo-landing/internal/app/web-server/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	router  *gin.Engine
	handler *handler.Handler

	port string
}

func NewServer(handler *handler.Handler, port string) *Server {
	router := gin.New()
	server := &Server{
		router:  router,
		handler: handler,
		port:    port,
	}

	return server
}

func (s *Server) setupRoutes() {
	// serve html files
	s.router.LoadHTMLGlob("public/*.html") // Only consider .html files
	s.router.StaticFS("/static", http.Dir("public"))
	//s.router.Static("/static", "./public")

	s.router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	s.router.POST("/contact", s.handler.HandleContactForm)
}

func (s *Server) Run() error {
	s.setupRoutes()
	return s.router.Run(s.port)
}
