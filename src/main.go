package main

import (
	"fmt"
	"log"
	"net/http"
)

type myHandler struct {
	greeting string
}

func (mh myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("%v world", mh.greeting)))
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mh := myHandler{greeting: "Halos"}
		w.Write([]byte(fmt.Sprintf("%v world!", mh.greeting)))
	})
	http.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request) {
		mh := myHandler{greeting: "Halos"}
		w.Write([]byte(fmt.Sprintf("%v Foo!", mh.greeting)))
	})
	redirect := http.RedirectHandler("/foo", 302)
	http.HandleFunc("/back-to-foo", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Redirecting to foo")
		redirect.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8000", nil)
}
