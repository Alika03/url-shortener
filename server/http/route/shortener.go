package route

import (
	httpSwagger "github.com/swaggo/http-swagger"
	"net/http"
	_ "shortener/docs"
	"shortener/server/http/handler/add"
	"shortener/server/http/handler/get"
	"shortener/server/http/lib/mux"
)

var ShortenerRoutes = mux.Routes{
	HeadUrl: "", // @Failure 400

	Routes: []mux.Route{
		{
			Method:  http.MethodPost,
			Path:    "/add",
			Handler: add.Handle,
		},
		{
			Method:  http.MethodGet,
			Path:    "/{code}",
			Handler: get.Handle,
		},
		{
			Method:  http.MethodGet,
			Path:    "/swagger/*",
			Handler: httpSwagger.Handler(),
		},
	},
	Middlewares: nil,
}
