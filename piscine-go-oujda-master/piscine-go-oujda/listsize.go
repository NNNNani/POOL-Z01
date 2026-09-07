package piscine


func ListSize(l *List) int {
	if l.Head == nil {
		return 0
	}
	return 1 + ListSize(&List{Head: l.Head.Next})
}
