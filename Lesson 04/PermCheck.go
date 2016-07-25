/*
https://codility.com/programmers/task/perm_check/
*/
package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//Solution - Codility.com - PermCheck
func Solution(A []int) int {
	// write your code in Go 1.4
	N := len(A)
	positions := make([]int, N)
	var sum int
	for _, v := range A {
		if v > N {
			return 0
		}
		if positions[v-1] != 1 {
			positions[v-1] = 1
			sum += v
		}
	}
	if sum == (N*(N+1))/2 {
		return 1
	}
	return 0
}
