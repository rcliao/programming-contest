package web

import (
	"fmt"
	"net/http"
)

func Hello() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello Contest!")
	})
}
