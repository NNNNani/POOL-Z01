package piscine

func ToLower(s string) string {
	r := []rune(s)
	st := ""
	for i := 0; i < len(r); i++ {
		if r[i] >= 'A' && r[i] <= 'Z' {
			st += string(r[i] + 'a' - 'A')
		} else {
			st += string(r[i])
		}
	}
	return st
}
