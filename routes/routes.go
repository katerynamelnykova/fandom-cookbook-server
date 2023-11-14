package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/katerynamelnykova/fandom-cookbook-server/controllers"
)

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/health", controllers.Health())
	r.Get("/get-short-fandom-info", controllers.ShortFandomsInfo())
	r.Get("/book/{id}", controllers.GetBook())
	r.Get("/book/{id}/recipe/{rid}", controllers.GetRecipe())

	return r
}
