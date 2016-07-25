/*
https://codility.com/programmers/task/count_div/
*/
package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//Solution - Codility.com - CountDiv
func Solution(A int, B int, K int) int {
	// write your code in Go 1.4
	var offset int
	if A%K == 0 {
		offset++
	}
	return B/K - A/K + offset
}
