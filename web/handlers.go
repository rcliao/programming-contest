package web

import (
	"fmt"
	"html/template"
	"net/http"

	contest "github.com/rcliao/programming-contest"
)

// Hello simply return the hello world message
func Hello() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello, Programming Contest!")
	})
}

// Handle index page
func Index() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse templates
		indexTmpl := template.Must(template.ParseFiles("templates/index.html"))
		indexTmpl.Execute(w, struct{ Contest []contest.Contest }{})
	})
}

func UserSignUp(params) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})
}
