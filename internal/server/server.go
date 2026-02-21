package server

import (
	"linkyard/internal/imports"
	"linkyard/internal/links"
	"net/http"
)

type Server struct {
	Mux            *http.ServeMux
	linksHandler   *links.Handler
	importsHandler *imports.Handler
}

func NewServer(linksHandler *links.Handler, importsHandler *imports.Handler) *Server {
	srv := &Server{
		Mux:            http.NewServeMux(),
		linksHandler:   linksHandler,
		importsHandler: importsHandler,
	}

	srv.registerRoutes()

	return srv
}

func (s *Server) registerRoutes() {
	//register handleGetLinks as an http.Handler by wrapping it in http.HandlerFunc.
	//mux.Handle("/", http.HandlerFunc(srv.handleGetLinks))
	s.Mux.HandleFunc("GET /", s.linksHandler.HandleGetLinks)
	s.Mux.HandleFunc("POST /", s.linksHandler.HandleCreateLink)
	s.Mux.HandleFunc("POST /import", s.importsHandler.HandleImportLink)
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Mux.ServeHTTP(w, r)
}
