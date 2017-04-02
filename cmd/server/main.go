package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rcliao/programming-contest/web"
)

func main() {
	// TODO: parse Environment variable for configuration
	var addr = ":8080"

	// set up router
	r := mux.NewRouter()
	r.HandleFunc("/", web.Index()).Methods("GET")
	r.HandleFunc("/api/hello", web.Hello()).Methods("GET")

	// start up web application
	log.Printf("Starting programming-contest at port %v\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Http server failed to start at port "+addr, err)
	}
}
