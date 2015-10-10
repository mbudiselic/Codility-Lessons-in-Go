package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	var sum, n int
	for i := len(A) - 1; i >= 0; i-- {
		if A[i] == 1 {
			n++
		} else {
			sum += n
		}
	}
	if sum > 1000000000 {
		return -1
	}
	return sum
}
