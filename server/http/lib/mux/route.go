package mux

import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

type Routes struct {
	HeadUrl     string
	Routes      []Route
	Middlewares []MiddlewareFunc
}

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func (r *Routes) chainMiddlewares(handler http.HandlerFunc) http.HandlerFunc {
	var wrapped = handler
	for _, middleware := range r.Middlewares {
		wrapped = middleware(wrapped)
	}
	return wrapped
}

func RegisterRoute(mux *chi.Mux, r Routes) {
	for _, route := range r.Routes {
		mux.HandleFunc(r.HeadUrl+route.Path, route.checkMethod(r.chainMiddlewares(route.Handler)))
	}
}

func (route Route) checkMethod(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != route.Method {
			w.WriteHeader(http.StatusMethodNotAllowed)

			_, err := w.Write([]byte("\"error\":\"" + r.Method + " not allowed.\""))
			if err != nil {
				log.Printf("middlerware error: %v", err.Error())
			}
			return
		}
		next.ServeHTTP(w, r)
	}
}
