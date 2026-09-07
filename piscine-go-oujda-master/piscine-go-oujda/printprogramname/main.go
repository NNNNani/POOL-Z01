package main

import (
	"os"
	"sort"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	sort.Strings(args)

	for _, arg := range args {
		for _, val := range arg {
			z01.PrintRune(val)
		}
		z01.PrintRune('\n')
	}
}
