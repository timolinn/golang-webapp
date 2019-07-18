package controller

import (
	"fmt"
	"golang-webapp/src/viewmodel"
	"html/template"
	"log"
	"net/http"
)

type home struct {
	homeTemplate         *template.Template
	loginTemplate        *template.Template
	standLocatorTemplate *template.Template
}

func (h home) registerRoutes() {
	http.HandleFunc("/", h.handleHome)
	http.HandleFunc("/home", h.handleHome)
	http.HandleFunc("/login", h.handleLogin)
	http.HandleFunc("/standlocator", h.handleStandLocator)
}

func (h home) handleStandLocator(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewStandlocator()
	h.standLocatorTemplate.Execute(w, vm)
}

func (h home) handleLogin(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewLogin()
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			log.Println(fmt.Errorf("An error occured in: %v", err))
		}
		email := r.Form.Get("email")
		password := r.Form.Get("password")
		if email == "test@gmail.com" && password == "password" {
			http.Redirect(w, r, "/home", http.StatusTemporaryRedirect)
			return
		}
		vm.Password = password
		vm.Email = email
	}
	h.loginTemplate.Execute(w, vm)
}

func (h home) handleHome(w http.ResponseWriter, r *http.Request) {
	vm := viewmodel.NewHome()
	h.homeTemplate.Execute(w, vm)
}