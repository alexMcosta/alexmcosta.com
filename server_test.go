package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPage(t *testing.T) {
	t.Run("returns the Home page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		home(response, request)

		got := response.Body.String()
		want := "<h1>Hello World!</h1>"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("returns the Projects page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/projects", nil)
		response := httptest.NewRecorder()

		projects(response, request)

		got := response.Body.String()
		want := "<h1>Projects</h1>"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("return custom 404", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/asidhffsd", nil)
		response := httptest.NewRecorder()

		notFoundHandler(response, request)

		got := response.Body.String()
		want := "<h1>Page Not Found</h1>"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

}
