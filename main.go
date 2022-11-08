package main

import (
	"shortener/bootstrap"
	"shortener/server/http"
)

func init() {
	bootstrap.InitConfig()
	bootstrap.Migrate("./infrastructure/db/migration/")
}

// @title Url-shortener App API
// @version 1.0m
// @description API Server for Url-shortener Application

// @host localhost:8080

func main() {
	//// calling http server with all configuration
	s := http.NewHttpServer()
	//// listen and server
	s.Run()
}
