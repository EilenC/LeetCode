package part1

import "bytes"

func removeComments(source []string) []string {
	var res []string
	var newLine bytes.Buffer
	block := false
	for _, line := range source {
		for i := 0; i < len(line); i++ {
			if block {
				if i+1 < len(line) && line[i] == '*' && line[i+1] == '/' {
					block = false
					i++
				}
			} else {
				if i+1 < len(line) && line[i] == '/' && line[i+1] == '*' {
					block = true
					i++
				} else if i+1 < len(line) && line[i] == '/' && line[i+1] == '/' {
					break
				} else {
					newLine.WriteByte(line[i])
				}
			}
		}
		if !block && newLine.Len() > 0 {
			res = append(res, newLine.String())
			newLine = bytes.Buffer{}
		}
	}
	return res
}
