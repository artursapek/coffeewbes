package server

import (
	"fmt"
	"html"
	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func Serve() {
	http.HandleFunc("/", rootHandler)
	http.ListenAndServe(":8082", nil)
}
