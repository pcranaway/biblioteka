package routes

import "github.com/go-chi/chi/v5"

func SetupRoutes(router chi.Router) {
    router.Get("/", Index)
}
