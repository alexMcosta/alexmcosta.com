package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Hello World!</h1>")
}

func projects(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	fmt.Fprint(w, "<h1>Projects</h1>")
}

func notFoundHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<h1>Page Not Found</h1>")
}

func main() {

	//Define router
	r := mux.NewRouter()

	//Custom 404
	var notFound http.Handler = http.HandlerFunc(notFoundHandler)
	r.NotFoundHandler = notFound

	r.HandleFunc("/", home)
	r.HandleFunc("/projects", projects)
	http.ListenAndServe(":3000", r)
}
