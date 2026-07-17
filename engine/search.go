package engine

import (
    "bufio"
    "bytes"
    "container/list"
    "embed"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
    "net/url"
    "regexp"
)

type apiResponse struct {
    Lemma string `json:"lemma"`
}

var (
    //go:embed assets/*
    content embed.FS
)

func Search(pattern string, language string, stage string, dictionary string) []string {
    if !InputValidation(pattern, language) {
        return []string{"Invalid input"}
    }

    var matches []string

    if language == "latin" {
        matches = SearchLatin(pattern, stage, dictionary)
    } else if language == "german" {
        matches = SearchGerman(pattern, stage, dictionary)
    }

    return matches
}

func SearchLatin(pattern string, stage string, dictionary string) []string {
    regex := TransformToRegEx("^" + pattern + "$")
    matches := SearchInFile(regex)

    return TransformToArray(matches)
}

func SearchGerman(pattern string, stage string, dictionary string) []string {
    dictionary = "GWB"
    searchPatterns := TransformToWoerterbuchPattern(pattern)
    var allLemmata []string

    for _, pattern := range searchPatterns {
        apiURL := fmt.Sprintf(
            "https://api.woerterbuchnetz.de/open-api/dictionaries/%s/lemmata/%s",
            url.PathEscape(dictionary),
            url.PathEscape(pattern),
        )

        resp, err := http.Get(apiURL)
        if err != nil {
            fmt.Printf("failed to fetch pattern '%s': %v\n", pattern, err)
            continue
        }

        if resp.StatusCode != http.StatusOK {
            resp.Body.Close()
            fmt.Printf("API returned status %d for pattern '%s'\n", resp.StatusCode, pattern)
            continue
        }

        bodyBytes, err := io.ReadAll(resp.Body)
        resp.Body.Close()
        if err != nil {
            fmt.Printf("failed to read body for pattern '%s': %v\n", pattern, err)
            continue
        }

        trimmed := bytes.TrimSpace(bodyBytes)
        if len(trimmed) == 0 {
            continue
        }

        if trimmed[0] == '[' {
            var results []apiResponse
            if err := json.Unmarshal(trimmed, &results); err != nil {
                fmt.Printf("failed to parse JSON array for pattern '%s': %v\n", pattern, err)
                continue
            }
            for _, res := range results {
                allLemmata = append(allLemmata, res.Lemma)
            }
        } else if trimmed[0] == '{' {
            var root map[string]interface{}
            if err := json.Unmarshal(trimmed, &root); err != nil {
                fmt.Printf("failed to parse JSON object for pattern '%s': %v\n", pattern, err)
                continue
            }
            
            if lemma, ok := root["lemma"].(string); ok {
                allLemmata = append(allLemmata, lemma)
            }
            
            for _, val := range root {
                if arr, ok := val.([]interface{}); ok {
                    for _, item := range arr {
                        if obj, ok := item.(map[string]interface{}); ok {
                            if lemma, ok := obj["lemma"].(string); ok {
                                allLemmata = append(allLemmata, lemma)
                            }
                        }
                    }
                }
            }
        }
    }

    return allLemmata
}

func SearchInFile(regex *regexp.Regexp) *list.List {
    matches := list.New()
    errorMessage := list.New()

    file, err := content.Open("assets/latin.txt")
    if err != nil {
        fmt.Println("Error:", err)
        errorMessage.PushBack("Error: Problem opening the file")
        return errorMessage
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)

    for scanner.Scan() {
        line := scanner.Text()

        if regex.MatchString(line) {
            matches.PushBack(line)
        }
    }

    if err := scanner.Err(); err != nil {
        fmt.Println("Error:", err)
        errorMessage.PushBack("Error: Problem reading the file")
        return errorMessage
    }

    if matches.Len() == 0 {
        errorMessage.PushBack("No matches found")
        return errorMessage
    }

    return matches
}

func InputValidation(pattern string, language string) bool {
    if len(pattern) == 0 {
        return false
    }

    if language == "latin" {
        regex := regexp.MustCompile("^([a-zA-Z]+|(\\?{})+" +
            "|(\\?{([a-zA-Z],)+|[a-zA-Z]+})+|(\\?\\?)+)+$")

        return regex.MatchString(pattern)
    }

    //Add German Input Validation
    return true
}
