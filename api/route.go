package api

import (
	"github.com/go-chi/chi/v5"
	_ "github.com/otie173/odncore/docs/api"
	_ "github.com/swaggo/http-swagger"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Odncore API
// @version 1.0
// @description API server for Odinbit game
// @license.name  MIT
// @license.url   https://opensource.org/license/mit
// @host localhost:8080
// @BasePath /api/v1
func InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"),
	))

	r.Route("/api", func(api chi.Router) {
		api.Route("/v1", func(v1 chi.Router) {
			v1.Route("/server", func(server chi.Router) {
				// @Tag.name server
				// @Tag.description Server management endpoints
				server.Get("/info", InfoHandler)
				server.Get("/status", StatusHandler)
			})

			v1.Route("/world", func(world chi.Router) {
				// @Tag.name world
				// @Tag.description World management endpoints
				world.Get("/getworld", GetWorldHandler)
				world.Post("/loadid", LoadIdHandler)
				world.Post("/loadworld", LoadWorldHandler)
			})

			v1.Route("/player", func(player chi.Router) {
				// @Tag.name player
				// @Tag.description Player management endpoints
				player.Get("/getpdata", GetPDataHandler)
				player.Post("/auth", AuthHandler)
				player.Post("/loadpdata", LoadPDataHandler)
			})
		})
	})

	return r
}
