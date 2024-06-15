package main

import (
	"fmt"
	"log"
	"net/http"
	"web/rabbit"
)

func main() {
	log.Println("Connecting to rmq client")
	go rabbit.IntializeRMQClient()
	fmt.Println("Starting http server @ localhost:2137")
	router := GetRoutes();
	http.ListenAndServe(":2137", &router)
}
