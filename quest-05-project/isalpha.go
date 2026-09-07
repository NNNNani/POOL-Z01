package piscine

func IsAlpha(s string) bool {
	for r := 0; r < len(s); r++ {
		if s[r] < '0' || (s[r] > '9' && s[r] < 'A') || (s[r] > 'Z' && s[r] < 'a') || s[r] > 'z' {
			return false
		}
	}
	return true
}
