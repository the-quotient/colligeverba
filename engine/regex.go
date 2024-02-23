package engine

import (
	"strings"
)

func TransformToRegEx(pattern string) string {

	step1 := strings.Replace(pattern, "?{}", ".", -1)
	step2 := strings.Replace(step1, "?{", "[", -1)
	step3 := strings.Replace(step2, "}", "]", -1)

	return step3

}
