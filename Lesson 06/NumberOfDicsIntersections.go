package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	count := 0
	for i := 0; i < len(A)-1; i++ {
		for j := i + 1; j < len(A); j++ {
			if i+A[i] >= j-A[j] {
				count++
			}
			if count > 10000000 {
				return -1
			}
		}
	}
	return count
}
