package engine

import (
    "regexp"
    "strings"
)

func TransformToRegEx(pattern string) *regexp.Regexp {

    step1 := strings.Replace(pattern, "?{}", ".", -1)
    step2 := strings.Replace(step1, "?{", "[", -1)
    step3 := strings.Replace(step2, "}", "]", -1)
    step4 := strings.Replace(step3, "??", ".{1,2}", -1)

    return regexp.MustCompile(step4)
}

func TransformToWoerterbuchPattern(pattern string) []string {

    step1 := strings.Replace(pattern, "?{}", "*", -1)
    patternRe := regexp.MustCompile(`\?\{([^}]+)\}`)
    return ExpandPatterns(step1, patternRe) 

}

// ExpandPatterns takes a string with ?{a,b,c...} patterns and returns all combinations.
func ExpandPatterns(s string, patternRe *regexp.Regexp) []string {
    loc := patternRe.FindStringSubmatchIndex(s)

    if loc == nil {
        return []string{s}
    }

    // Extract the block of characters and split them by commas
    variants := strings.Split(s[loc[2]:loc[3]], ",")
    var result []string

    // Loop through each variant
    for _, v := range variants {
        // Construct the new string with the current variant
        newStr := s[:loc[0]] + v + s[loc[1]:]

        // Recursively expand any remaining patterns and append
        result = append(result, ExpandPatterns(newStr, patternRe)...)
    }

    return result
}
