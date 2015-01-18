package server

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Serve() {
	http.HandleFunc("/coffee/reviews/", reviewHandler)
	http.HandleFunc("/coffee/assets/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[8:])
		log.Println(r.URL.Path[8:])
	})
	http.ListenAndServe(":8082", nil)
}
