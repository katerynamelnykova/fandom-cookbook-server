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
	r.Get("/api/fandoms/fandoms-highlights", controllers.ShortFandomsInfo())
	r.Get("/api/fandoms/fandom/{fandom}", controllers.GetBook())
	r.Get("/api/fandoms/fandom/{fandom}/{recipe}", controllers.GetRecipe())

	return r
}
