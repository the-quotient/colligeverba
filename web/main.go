package main

import (
	"embed"
	"github.com/the-quotient/vestigiaverbi/engine"
	"html/template"
	"log"
	"net/http"
)

var (
	//go:embed html/*
	content embed.FS
)

var rootTmpl *template.Template

type Words struct {
	Results []string
}

type WordInfo struct {
    Word string
    BasicForm string
    Meanings []string
    FormAnalysis []string
}

func getRoot(w http.ResponseWriter, r *http.Request) {
	log.Printf("got / request\n")
    rootTmpl = template.Must(template.ParseFS(content, "html/index.html"))
	rootTmpl.Execute(w, r.UserAgent())
}

func query(w http.ResponseWriter, r *http.Request) {
    rootTmpl = template.Must(template.ParseFS(content, "html/index.html"))
	query := r.FormValue("query")
    log.Printf("Query %s \n", query)
	res := engine.Search(query)
    log.Printf("The result has length: %d", len(res))
	if len(res) > 10 {
		res = res[:10]
	}
	rootTmpl.ExecuteTemplate(w, "word-list-el", Words{Results: res})
}

func wordInfo(w http.ResponseWriter, r *http.Request) {
    infoTmpl := template.Must(template.ParseFS(content, "html/modal.html"))
    word := r.PathValue("word")
    log.Println("/wordinfo for: ", word)
    bf, m, fa := engine.GetInformation(word)
    winfo := WordInfo{
        Word: word,
        BasicForm: bf,
        Meanings: m,
        FormAnalysis: fa,
    }
    infoTmpl.Execute(w, winfo)
}

func main() {
    http.HandleFunc("POST /wordinfo/{word}", wordInfo)
    http.HandleFunc("/query", query)
	http.HandleFunc("/", getRoot)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
