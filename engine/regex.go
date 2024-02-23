package engine

func TransformToRegEx(pattern string) string {

	return ReplaceChar(pattern, '?', '.')

}
