package piscine

func IsNumeric(s string) bool {
	for r := 0; r < len(s); r++ {
		if s[r] < '0' || s[r] > '9' {
			return false
		}
	}
	return true
}
