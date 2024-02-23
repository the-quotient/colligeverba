package engine

func ReplaceChar(s string, oldChar byte, newChar byte) string {
	strBytes := []byte(s)

	for i := range strBytes {
		if strBytes[i] == oldChar {
			strBytes[i] = newChar
		}
	}

	return string(strBytes)
}
