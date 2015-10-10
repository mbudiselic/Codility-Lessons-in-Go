package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	var offset int
	if A%K == 0 {
		offset++
	}
	return B/K - A/K + offset
}
