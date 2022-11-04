package mux

import "net/http"

type MiddlewareFunc func(next http.HandlerFunc) http.HandlerFunc
