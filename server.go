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
	server.configServer()
	return server
}

func (s *Server) configServer() {
	s.setRoutes()
}

func (s *Server) setRoutes() {
	s.router.AddRoute("/", "GET", HomeRoute)
}

func (s *Server) Listen() {
	http.Handle("/", s.router)
	ws := NewWebSocket()
	http.HandleFunc("/socket", ws.Listen)
	PrintServerInfo("Server listening on port " + s.port)
	err := http.ListenAndServe(s.port, nil)
	if err != nil {
		fmt.Println("Error creating the server")
	}
}
