package server;

import "net/http"
import "github.com/go-chi/chi/v5"
import "github.com/pcranaway/biblioteka/env"
import "github.com/pcranaway/biblioteka/routes"

type Server struct {
    router      chi.Router
    environment env.Environment
}

func NewServer(environment env.Environment) *Server {
    s := new(Server)

    s.environment = environment

    // setup router
    s.router = chi.NewRouter()
    routes.SetupRoutes(s.router)

    return s
}

func (s *Server) Start() {
    http.ListenAndServe(":8080", s.router)
}
