package main

import (
	"database/sql"
	"embed"
	"html/template"
	"log"
	"net/http"
	"os"
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
var useDB bool

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
		
		// Try DB first (if enabled)
		if useDB {
			foundInfo, err := db.GetItemByWord(gdb, word)
			if err == nil {
				log.Printf("Returning WordInfo for '%s' from Cache! \n", word)
				return foundInfo
			}
		}

		// Not in DB (or DB disabled) -> Fetch Online
		log.Printf("Cannot find '%s' in db (or db disabled), trying online \n", word)
		bf, m, fa := engine.GetInformation(word)
		winfo = db.WordInfo{
			Word:         word,
			BasicForm:    bf,
			Meanings:     m,
			FormAnalysis: fa,
		}

		// Save to DB (if enabled)
		if useDB {
			err := db.InsertItem(gdb, winfo)
			if err != nil {
				log.Println("Failed to insert in DB")
				log.Println(err)
			}
		}
		return winfo
	}
	return winfo
}

func main() {
	// Check Environment Variable to toggle Database
	if os.Getenv("DISABLE_DB") == "true" {
		useDB = false
		log.Println("Running in Stateless Mode (Database Disabled)")
	} else {
		useDB = true
		log.Println("Running in Stateful Mode (Database Enabled)")
		gdb = db.InitDb()
		defer db.CloseDb(gdb)
	}

	http.HandleFunc("POST /wordinfo/{word}", wordInfo)
	http.HandleFunc("/query", query)
	http.HandleFunc("/", getRoot)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
