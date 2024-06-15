package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting http server @ localhost:2137")
	router := GetRoutes();
	http.ListenAndServe(":2137", &router)
}
