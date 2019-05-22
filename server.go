package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	newrelic "github.com/newrelic/go-agent"
	nrgorilla "github.com/newrelic/go-agent/_integrations/nrgorilla/v1"
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
	//Config New Relic
	cfg := newrelic.NewConfig("alexmcosta.com", getNREnv())
	cfg.Logger = newrelic.NewDebugLogger(os.Stdout)
	app, err := newrelic.NewApplication(cfg)
	if nil != err {
		fmt.Println(err)
		os.Exit(1)
	}

	//Define router
	r := mux.NewRouter()

	//Custom 404
	var notFound http.Handler = http.HandlerFunc(notFoundHandler)
	r.NotFoundHandler = notFound

	r.HandleFunc("/", home)
	r.HandleFunc("/projects", projects)
	http.ListenAndServe(":3000", nrgorilla.InstrumentRoutes(r, app))
}
