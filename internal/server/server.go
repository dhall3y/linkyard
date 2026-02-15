package server

import (
	"linkyard/internal/links"
	"net/http"
)

type Server struct {
	Mux          *http.ServeMux
	linksHandler *links.Handler
}

func NewServer(linksHandler *links.Handler) *Server {
	srv := &Server{
		Mux:          http.NewServeMux(),
		linksHandler: linksHandler,
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) registerRoutes() {
	//register handleGetLinks as an http.Handler by wrapping it in http.HandlerFunc.
	//mux.Handle("/", http.HandlerFunc(srv.handleGetLinks))
	s.Mux.HandleFunc("/", s.linksHandler.HandleGetLinks)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}
