package main

import (
	"log"
	"net/http"
	"text/template"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/index.html"))
		temp.Execute(w, nil)
	})

	http.HandleFunc("/search", func(w http.ResponseWriter, r *http.Request) {
		temp := template.Must(template.ParseFiles("./templates/fregments"))
		temp.Execute(w, nil)
	})

	log.Println("App runing on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", nil))
}
