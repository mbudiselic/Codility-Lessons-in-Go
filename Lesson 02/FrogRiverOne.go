package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, A []int) int {
	// write your code in Go 1.4
	positions := make([]int, X)
	var moves int
	for i, v := range A {
		if positions[v-1] != 1 {
			positions[v-1] = 1
			moves++
		}
		if moves == X {
			return i
		}
	}
	return -1
}
