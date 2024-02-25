package main

import (
	"database/sql"
	"embed"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/the-quotient/colligeverba/db"
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

var gdb *sql.DB

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

func getWordInfo(word string) db.WordInfo {
	var winfo db.WordInfo
	if word != "Invalid input" && word != "No matches found" {
		winfo, err := db.GetItemByWord(gdb, word)
		if err != nil {
			log.Printf("Cannot find '%s' in db, trying online \n", word)
			bf, m, fa := engine.GetInformation(word)
			winfo = db.WordInfo{
				Word:         word,
				BasicForm:    bf,
				Meanings:     m,
				FormAnalysis: fa,
			}
            err = db.InsertItem(gdb, winfo)
            if err != nil {
                log.Println("Failed to insert in DB")
                log.Println(err)
            }
            return winfo
		} else {
            log.Printf("Returning WordInfo for '%s' from Cache! \n", word)
            return winfo
        }
	}
	//try to find in db
	return winfo
}

func main() {
	http.HandleFunc("POST /wordinfo/{word}", wordInfo)
	http.HandleFunc("/query", query)
	http.HandleFunc("/", getRoot)

	gdb = db.InitDb()
    defer db.CloseDb(gdb)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
