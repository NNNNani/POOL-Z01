package piscine

func ToUpper(s string) string {
	r := []rune(s)
	st := ""
	for i := 0; i < len(r); i++ {
		if r[i] >= 'a' && r[i] <= 'z' {
			st += string(r[i] - 'a' + 'A')
		} else {
			st += string(r[i])
		}
	}
	return st
}
