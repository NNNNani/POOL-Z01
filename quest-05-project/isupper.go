package piscine

func IsUpper(s string) bool {
	for r := 0; r < len(s); r++ {
		if s[r] < 'A' || s[r] > 'Z' {
			return false
		}
	}
	return true
}
