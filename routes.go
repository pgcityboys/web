package main

import (
	"net/http"
	"web/templates"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

func GetRoutes() (router http.ServeMux) {
	router.HandleFunc("/", createRootHandler(templates.HandleHomePage, http.NotFound))
	router.HandleFunc("/login", templates.HandleLoginPage)
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	return router
}

func createRootHandler(rootHandler, notFoundHandler HttpHandler) HttpHandler {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			rootHandler(w, r)
		} else {
			notFoundHandler(w, r)
		}
	}
}
