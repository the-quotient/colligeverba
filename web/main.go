package main

import (
	"log"
	"net/http"
    "html/template"
)

var rootTmpl *template.Template

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request\n")
    rootTmpl.Execute(w, r.UserAgent())
}

func query(w http.ResponseWriter, r *http.Request) {
    log.Printf(r.PostFormValue("query"))
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/query", query)
    rootTmpl = template.Must(template.ParseFiles("html/index.html"))


	log.Fatal(http.ListenAndServe(":8080", nil))
}
