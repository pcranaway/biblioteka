package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/pcranaway/biblioteka/db"
	"github.com/pcranaway/biblioteka/env"
	"github.com/pcranaway/biblioteka/routes"
	"gorm.io/gorm"
)

type Server struct {
    router      chi.Router
    environment env.Environment
    database    gorm.DB
}

func NewServer(environment env.Environment) Server {
    s := new(Server)

    s.environment = environment

    // setup router
    s.router = chi.NewRouter()
    routes.SetupRoutes(s.router)

    // setup db
    s.database = *db.ConnectDatabase(&environment)

    return *s
}

func (s *Server) Start() {
    http.ListenAndServe(":8080", s.router)
}
