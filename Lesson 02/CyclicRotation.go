package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	// write your code in Go 1.4
	n := len(A)
	if n == 0 {
		return []int{}
	}

	var result = A
	for i := K; i > 0; i-- {
		result = append(result[n-1:], result[0:n-1]...)
	}
	return result
}
