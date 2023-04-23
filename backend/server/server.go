package server

import (
	"net/http"
)

func ServerHanlder() http.Handler{
	mux := http.NewServeMux()
	dir := http.Dir("../backend/static/")
	
	// static file
  fs := http.FileServer(dir)
	mux.Handle("/", fs)

	// graphql

	// rest api

	return mux
}