package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Game struct {
	Name string
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func helloRouterHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello %s!", name)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	// parse templates
	indexTmpl := template.Must(template.ParseFiles("templates/index.html"))
	games := []Game{
		{"Tic-tac-toe"},
		{"Rock paper scissor"},
	}
	indexTmpl.Execute(w, struct{ Games []Game }{games})
}

func main() {
	// TODO: parse Environment variable for configuration
	var addr = ":8080"

	// set up router
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler).Methods("GET")
	r.HandleFunc("/hello", helloHandler).Methods("GET")
	r.HandleFunc("/hello/{name}", helloRouterHandler).Methods("GET")

	// start up web application
	log.Printf("Start programming-contest at port %v\n", addr)
	if err := http.ListenAndServe(addr, r); err != nil {
		log.Fatal("Http server failed to start at port "+addr, err)
	}
}
