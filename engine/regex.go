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
