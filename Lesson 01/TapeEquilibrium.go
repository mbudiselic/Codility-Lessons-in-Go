package solution

// you can also use imports, for example:
// import "fmt"
import (
	"math"
)

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//Solution Codility.com - TapeEquilibrium
func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	min := math.MaxUint32
	startSum := A[0]
	endSum := sumSlice(A[1:])
	for P := 1; P < N; P++ {
		diff := startSum - endSum
		if diff < 0 {
			diff = -diff
		}
		if diff < min {
			min = diff
		}
		startSum += A[P]
		endSum -= A[P]
	}
	return min
}

func sumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}
