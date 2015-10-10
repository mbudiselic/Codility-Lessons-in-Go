package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A) + 1
	fullSum := (N * (N + 1)) / 2
	sum := 0
	for _, v := range A {
		sum += v
	}
	return fullSum - sum
}
