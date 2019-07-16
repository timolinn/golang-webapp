package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	templates := populateTemplates()
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		requestedFile := r.URL.Path[1:]
		log.Println(requestedFile)
		t := templates[requestedFile+".html"]
		if t != nil {
			err := t.Execute(w, nil)
			if err != nil {
				log.Println(err)
			}
		} else {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Oh noes!! You got lost, my bad!"))
		}
	})
	http.Handle("/img/", http.FileServer(http.Dir("public")))
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.ListenAndServe(":8000", nil)
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
