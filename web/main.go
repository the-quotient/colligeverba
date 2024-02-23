package main

import (
	"log"
	"net/http"
    "html/template"
    "github.com/the-quotient/vestigiaverbi/engine"
)


var rootTmpl *template.Template

type Words struct{
    Results []string
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request\n")
    rootTmpl = template.Must(template.ParseFiles("html/index.html"))
    rootTmpl.Execute(w, r.UserAgent())
}

func query(w http.ResponseWriter, r *http.Request) {
    rootTmpl = template.Must(template.ParseFiles("html/index.html"))
    query := r.FormValue("query")
    res := engine.Search(query)
    if len(res) > 10 {
        res = res[:10]
    }
    rootTmpl.ExecuteTemplate(w, "word-list-el", Words{Results: res})
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/query", query)


	log.Fatal(http.ListenAndServe(":8080", nil))
}
