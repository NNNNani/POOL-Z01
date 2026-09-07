package piscine

func TrimAtoi(s string) int {
	n := 0
	si := 1
	h := false

	for _, r := range s {
		if r >= '0' && r <= '9' {
			h = true
			n = n*10 + int(r-'0')
		} else if r == '-' && !h {
			si = -1
		}
	}
	return n * si
}
