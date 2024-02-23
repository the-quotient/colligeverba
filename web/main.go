package main

import (
	"log"
	"net/http"
    "html/template"
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
    testData := []string{"Hallo", "WElt", "lorem", "ipsum" }
    res := Words{Results: testData}
    rootTmpl.ExecuteTemplate(w, "word-list-el", res)
}

func main() {
    http.HandleFunc("/", getRoot)
    http.HandleFunc("/query", query)


	log.Fatal(http.ListenAndServe(":8080", nil))
}
