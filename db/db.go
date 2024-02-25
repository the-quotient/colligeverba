package db

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

type WordInfo struct {
	Word         string
	BasicForm    string
	Meanings     []string
	FormAnalysis []string
}

func CloseDb(db *sql.DB) {
	db.Close()
}

func InitDb() *sql.DB {
	var gdb *sql.DB
	// Open a connection to the SQLite database
	gdb, err := sql.Open("sqlite3", "/data/db.db")
	if err != nil {
		panic(err)
	}

	// Create a table to store items
	createTableSQL := `
        CREATE TABLE IF NOT EXISTS wordInfo (
            word TEXT PRIMARY KEY,
            basicform TEXT,
            meanings TEXT,
            formanalysis TEXT
        );
    `
	_, err = gdb.Exec(createTableSQL)
	if err != nil {
		panic(err)
	}
	log.Println("db initialized")
	return gdb
}

// Function to insert an item into the database
func InsertItem(db *sql.DB, item WordInfo) error {
	_, err := db.Exec("INSERT INTO wordInfo(word, basicform, meanings, formanalysis) VALUES(?, ?, ?, ?)", item.Word, item.BasicForm, arrayToString(item.Meanings), arrayToString(item.FormAnalysis))
	return err
}

// Function to retrieve an item from the database by word
func GetItemByWord(db *sql.DB, word string) (WordInfo, error) {
	var item WordInfo
	var meanings, formAnalysis string
	err := db.QueryRow("SELECT word, basicform, meanings, formanalysis FROM wordInfo WHERE word = ?", word).Scan(&item.Word, &item.BasicForm, &meanings, &formAnalysis)
	item.Meanings = stringToArray(meanings)
	item.FormAnalysis = stringToArray(formAnalysis)
	return item, err
}

// Function to update an item in the database
func UpdateItem(db *sql.DB, item WordInfo) error {
	_, err := db.Exec("UPDATE wordInfo SET basicform = ?, meanings = ?, formanalysis = ? WHERE word = ?", item.BasicForm, arrayToString(item.Meanings), arrayToString(item.FormAnalysis), item.Word)
	return err
}

// Function to delete an item from the database by word
func DeleteItem(db *sql.DB, word string) error {
	_, err := db.Exec("DELETE FROM wordInfo WHERE word = ?", word)
	return err
}

// Helper function to convert array to string
func arrayToString(arr []string) string {
	result := ""
	for i, v := range arr {
		if i > 0 {
			result += ","
		}
		result += v
	}
	return result
}

// Helper function to convert string to array
func stringToArray(s string) []string {
	return strings.Split(s, ",")
}
