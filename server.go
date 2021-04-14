package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	port   string
	router *Router
}

func NewServer() *Server {
	server := &Server{
		":3000",
		NewRouter(),
	}
	server.ConfigServer()
	return server
}

func (s *Server) ConfigServer() {
	s.SetRoutes()
}

func (s *Server) SetRoutes() {
	s.router.AddRoute("/", "GET", HomeRoute)
}

func (s *Server) CreateServer() {
	http.Handle("/", s.router)
	PrintServerInfo("Server listening on port " + s.port)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		fmt.Println("Error creating the server")
	}
}
