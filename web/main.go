package main

import (
    "embed"
    "html/template"
    "log"
    "net/http"
    "strings"

    "github.com/the-quotient/colligeverba/engine"
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
	Word         string
	BasicForm    string
	Meanings     []string
	FormAnalysis []string
}

func getRoot(w http.ResponseWriter, r *http.Request) {
    rootTmpl = template.Must(template.ParseFS(content, "html/index.html"))
    rootTmpl.Execute(w, r.UserAgent())
}

func query(w http.ResponseWriter, r *http.Request) {
    rootTmpl = template.Must(template.ParseFS(content, "html/index.html"))
    query := r.FormValue("query")
    log.Printf("Query '%s' \n", query)
    res := engine.Search(strings.ToLower(query))
    if len(res) > 10 {
        res = res[:10]
    }
    rootTmpl.ExecuteTemplate(w, "word-list-el", Words{Results: res})
}

func wordInfo(w http.ResponseWriter, r *http.Request) {
    infoTmpl := template.Must(template.ParseFS(content, "html/modal.html"))
    word := r.PathValue("word")
    winfo := getWordInfo(word)
    infoTmpl.Execute(w, winfo)
}

func getWordInfo(word string) WordInfo {
    var winfo WordInfo

    if word != "Invalid input" && word != "No matches found" {

        bf, m, fa := engine.GetInformation(word)
        winfo = WordInfo{
            Word:         word,
            BasicForm:    bf,
            Meanings:     m,
            FormAnalysis: fa,
        }

    }
    return winfo
}

func main() {

    http.HandleFunc("POST /wordinfo/{word}", wordInfo)
    http.HandleFunc("/query", query)
    http.HandleFunc("/", getRoot)

    log.Println("Server starting on :8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}
