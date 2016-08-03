package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	// write your code in Go 1.4
	var openBrackets int
	for _, c := range S {
		if c == '(' {
			openBrackets++
		} else {
			if openBrackets > 0 {
				openBrackets--
			} else {
				return 0
			}
		}
	}
	if openBrackets == 0 {
		return 1
	}
	return 0
}
