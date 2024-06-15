package main

import (
	"net/http"
	"web/api"
	"web/templates"
)

type HttpHandler func(http.ResponseWriter, *http.Request)

func GetRoutes() (router http.ServeMux) {
	router.HandleFunc("/", createRootHandler(templates.HandleHomePage, http.NotFound))
	router.HandleFunc("/call", templates.HandleCallPage)
	router.HandleFunc("/friends", templates.HandleFriendsPage)
	router.HandleFunc("/timetable", templates.HandleTimeTablePage)
  
	router.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Api endpoints
	router.HandleFunc("GET /api/matchreq", api.HandleMatchRequest)
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
