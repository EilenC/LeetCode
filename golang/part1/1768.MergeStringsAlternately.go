package part1

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	var s strings.Builder
	var a int
	if len(word1) < len(word2) {
		a = len(word1)
	} else {
		a = len(word2)
	}

	for i := 0; i < a; i++ {
		s.WriteByte(word1[i])
		s.WriteByte(word2[i])
	}
	s.WriteString(word1[a:])
	s.WriteString(word2[a:])
	return s.String()
}
