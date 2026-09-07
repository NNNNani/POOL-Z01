package piscine

func Capitalize(s string) string {
	st := ""
	is := true
	var c rune
	for i := 0; i < len(s); i++ {
		c = rune(s[i])
		if c >= 'A' && c <= 'Z' && !is {
			st += string(c + 'a' - 'A')
		} else if c >= 'a' && c <= 'z' && is {
			st += string(c + 'A' - 'a')
			is = false
		} else {
			if (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9') {
				is = false
			} else if c < 'a' || c > 'z' {
				is = true
			}
			st += string(c)
		}
	}
	return st
}
