package controller

import (
	"golang-webapp/src/viewmodel"
	"html/template"
	"net/http"
)

type standlocator struct {
	standLocatorTemplate *template.Template
}

func (s standlocator) registerRoutes() {
	http.HandleFunc("/standlocator", s.handleStandLocator)
}

func (s standlocator) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandlocator()
	s.standLocatorTemplate.Execute(w, vm)
}
