package main

import (
	"database/sql"
	"fmt"
	"golang-webapp/src/controller"
	"golang-webapp/src/middleware"
	"golang-webapp/src/model"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	_ "net/http/pprof"

	_ "github.com/lib/pq"
)

func main() {
	templates := populateTemplates()
	db := connectToDatabase()
	defer db.Close()
	controller.StartUp(templates)
	go http.ListenAndServe(":8080", nil)
	http.ListenAndServeTLS(":8000", "cert.pem", "key.pem", &middleware.TimeoutMiddleware{new(middleware.GzipMiddleware)})
}

func connectToDatabase() *sql.DB {
	db, err := sql.Open("postgres", "postgres://timothyonyiuke:@localhost/lss?sslmode=disable")
	if err != nil {
		log.Fatalln(fmt.Errorf("unable to connect to database: %v", err))
	}
	log.Println("Connected to database")
	model.SetDatabase(db)
	return db
}

func populateTemplates() map[string]*template.Template {
	result := make(map[string]*template.Template)
	const basePath = "templates"
	var layout = template.Must(template.ParseFiles(basePath + "/_layout.html"))
	template.Must(layout.ParseFiles(basePath+"/_header.html", basePath+"/_footer.html"))
	dir, err := os.Open(basePath + "/content")
	if err != nil {
		panic("Failed to open template directory: " + err.Error())
	}
	files, err := dir.Readdir(-1)
	if err != nil {
		panic("Failed to read template directory: " + err.Error())
	}
	for _, file := range files {
		f, err := os.Open(basePath + "/content/" + file.Name())
		if err != nil {
			panic("Failed to open template '" + file.Name() + "'")
		}
		content, err := ioutil.ReadAll(f)
		if err != nil {
			panic("Failed to read file '" + file.Name() + "'")
		}
		f.Close()
		tmpl := template.Must(layout.Clone())
		_, err = tmpl.Parse(string(content))
		if err != nil {
			panic("Failed to parse template: '" + file.Name() + "'")
		}
		result[file.Name()] = tmpl
	}
	return result
}
