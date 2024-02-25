package engine

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
)

/*
* This function takes a word as an argument and returns the basic forms,
* meanings and form analysis of the word. Powered by latein.me
 */
func GetInformation(word string) (string, []string, []string) {

	resp, err := http.Get("https://www.latein.me/mixed/" + word)

	if err != nil {
		fmt.Println("Error:", err)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)

	//Parsing the HTML to extract the required information
	withoutPretext := strings.Split(sb,
		"<div class=\"contentBox\">")[1]
	withoutSynonyms := strings.Split(withoutPretext,
		"<dd class=\"synonymEntry\">")[0]
	basicFormsAndMeaningsAndFormAnalysis := strings.Split(withoutSynonyms,
		"<dd class=\"formAnalysisEntry\">")
	basicFormsAndMeanings := strings.Split(basicFormsAndMeaningsAndFormAnalysis[0],
		"<dd class=\"translationEntry\">")

	basicForms := basicFormsAndMeanings[0]
	meanings := basicFormsAndMeanings[1:]
	formAnalysis := basicFormsAndMeaningsAndFormAnalysis[1:]

	//Further cleaning of the elements
	cleanerRegEx := regexp.MustCompile("(^<.*(\")>)|(</a></dd>)|(</dd>)")
	basicForms = cleanerRegEx.ReplaceAllString(basicForms, "")
	for el := range meanings {
		meanings[el] = cleanerRegEx.ReplaceAllString(meanings[el], "")
	}
	for el := range formAnalysis {
		formAnalysis[el] = cleanerRegEx.ReplaceAllString(formAnalysis[el], "")
	}

	return basicForms, meanings, formAnalysis
}
