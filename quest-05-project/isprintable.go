package piscine

func IsPrintable(s string) bool {
	for r := 0; r < len(s); r++ {
		if s[r] < ' ' || s[r] > '~' {
			return false
		}
	}
	return true
}
