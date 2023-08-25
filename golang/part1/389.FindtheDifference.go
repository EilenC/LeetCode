package part1

func findTheDifference(s string, t string) byte {
	sm := make(map[byte]int)
	for i, _ := range s {
		sm[s[i]] = sm[s[i]] + 1
	}
	for i := range t {
		sm[t[i]] = sm[t[i]] + 1
	}
	for i := range sm {
		if sm[i]%2 != 0 {
			return i
		}
	}
	return s[0]
}
