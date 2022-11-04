package http

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
	"shortener/config"
	"shortener/server/http/lib/mux"
	"shortener/server/http/route"

	"time"
)

type server struct {
	httpServer http.Server
}

func NewHttpServer() *server {
	return &server{
		httpServer: http.Server{
			Addr:         ":" + config.GetConfig().ServerParam.Port,
			Handler:      registerRoutesHandles(),
			ReadTimeout:  30 * time.Second,
			WriteTimeout: 30 * time.Second,
		},
	}
}

func (s *server) Run() {
	log.Fatalln(s.httpServer.ListenAndServe())
}

func registerRoutesHandles() *chi.Mux {
	muxServe := chi.NewMux()

	// register all routes
	mux.RegisterRoute(muxServe, route.ShortenerRoutes)

	return muxServe
}
