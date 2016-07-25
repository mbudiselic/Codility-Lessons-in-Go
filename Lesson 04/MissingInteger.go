/*
https://codility.com/programmers/task/missing_integer/
*/
package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//Solution - Codility.com - MissingInteger
func Solution(A []int) int {
	// write your code in Go 1.4
	m := make(map[int]int)
	for _, v := range A {
		if v > 0 {
			m[v]++
		}
	}
	var i int
	for m[i+1] != 0 {
		i++
	}
	return i + 1
}
