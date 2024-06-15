package main

import (
	"fmt"
	"log"
	"net/http"
	"web/rabbit"
	"web/wsserver"
)

func main() {
	hub := wsserver.NewHub()
	go hub.Run()
	log.Println("Connecting to rmq client")
	go rabbit.IntializeRMQClient()
	fmt.Println("Starting http server @ localhost:2137")
	router := GetRoutes();
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		wsserver.ServeWs(hub, w, r)
	})
	http.ListenAndServe(":2137", &router)
}
