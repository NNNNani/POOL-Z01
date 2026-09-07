package piscine

func IsLower(s string) bool {
	for r := 0; r < len(s); r++ {
		if s[r] < 'a' || s[r] > 'z' {
			return false
		}
	}
	return true
}
