package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETPage(t *testing.T) {
	t.Run("Returns the Home page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		home(response, request)

		got := response.Body.String()
		want := "<h1>Hello World!</h1>"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})

	t.Run("Returns the Projects page", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/projects", nil)
		response := httptest.NewRecorder()

		projects(response, request)

		got := response.Body.String()
		want := "<h1>Projects</h1>"

		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
}
