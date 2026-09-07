package main

import (
	"fmt"
	"os"
)

func main() {
	input := normalize(readInput())
	if input == "" {
		fmt.Println("Not a quad function")
		return
	}

	width, height := getDimensions(input)
	if width == 0 || height == 0 {
		fmt.Println("Not a quad function")
		return
	}

	quads := []string{"quadA", "quadB", "quadC", "quadD", "quadE"}
	matches := []string{}

	for _, quad := range quads {
		var generated string
		switch quad {
		case "quadA":
			generated = normalize(generateQuadA(width, height))
		case "quadB":
			generated = normalize(generateQuadB(width, height))
		case "quadC":
			generated = normalize(generateQuadC(width, height))
		case "quadD":
			generated = normalize(generateQuadD(width, height))
		case "quadE":
			generated = normalize(generateQuadE(width, height))
		}

		if input == generated {
			matches = append(matches, fmt.Sprintf("[%s] [%d] [%d]", quad, width, height))
		}
	}

	if len(matches) > 0 {
		fmt.Println(joinMatches(matches))
	} else {
		fmt.Println("Not a quad function")
	}
}

func readInput() string {
	var input []byte
	buf := make([]byte, 1024)
	for {
		n, err := os.Stdin.Read(buf)
		if n > 0 {
			input = append(input, buf[:n]...)
		}
		if err != nil {
			break
		}
	}
	return string(input)
}

func normalize(input string) string {
	start, end := 0, len(input)
	for start < end && (input[start] == ' ' || input[start] == '\n' || input[start] == '\r') {
		start++
	}
	for end > start && (input[end-1] == ' ' || input[end-1] == '\n' || input[end-1] == '\r') {
		end--
	}
	return input[start:end]
}

func getDimensions(input string) (int, int) {
	width, height, currentWidth := 0, 0, 0
	for i := 0; i < len(input); i++ {
		if input[i] == '\n' {
			height++
			if currentWidth > width {
				width = currentWidth
			}
			currentWidth = 0
		} else {
			currentWidth++
		}
	}
	if currentWidth > 0 {
		height++
		if currentWidth > width {
			width = currentWidth
		}
	}
	return width, height
}

func joinMatches(matches []string) string {
	result := ""
	for i := 0; i < len(matches); i++ {
		if i > 0 {
			result += " || "
		}
		result += matches[i]
	}
	return result
}

func generateQuadA(x, y int) string {
	return generateQuadGeneric(x, y, 'o', '-', '|', ' ')
}

func generateQuadB(x, y int) string {
	return generateQuadGeneric(x, y, '/', '*', '*', '\\')
}

func generateQuadC(x, y int) string {
	return generateQuadGeneric(x, y, 'A', 'B', 'B', 'C')
}

func generateQuadD(x, y int) string {
	return generateQuadGeneric(x, y, 'A', 'B', 'B', 'A')
}

func generateQuadE(x, y int) string {
	return generateQuadGeneric(x, y, 'A', 'B', 'B', 'C')
}

func generateQuadGeneric(x, y int, corner, edgeH, edgeV, oppositeCorner rune) string {
	if x <= 0 || y <= 0 {
		return ""
	}
	output := make([]rune, 0, x*y+(y-1)) // Preallocate memory for efficiency
	for i := 1; i <= y; i++ {
		for j := 1; j <= x; j++ {
			if i == 1 {
				if j == 1 {
					output = append(output, corner)
				} else if j == x {
					output = append(output, oppositeCorner)
				} else {
					output = append(output, edgeH)
				}
			} else if i == y {
				if j == 1 {
					output = append(output, oppositeCorner)
				} else if j == x {
					output = append(output, corner)
				} else {
					output = append(output, edgeH)
				}
			} else {
				if j == 1 || j == x {
					output = append(output, edgeV)
				} else {
					output = append(output, ' ')
				}
			}
		}
		if i < y {
			output = append(output, '\n')
		}
	}
	return string(output)
}
