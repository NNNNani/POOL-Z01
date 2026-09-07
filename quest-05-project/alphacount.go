package piscine

func AlphaCount(s string) int {
	c := 0
	for _, r := range s {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			c++
		}
	}
	return c
}
