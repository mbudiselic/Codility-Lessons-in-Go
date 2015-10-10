package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(X int, Y int, D int) int {
	// write your code in Go 1.4
	diff := Y - X
	switch {
	case diff == 0:
		return 0
	case diff < D:
		return 1
	}
	r := diff % D
	switch {
	case r == 0:
		return diff / D
	default:
		return (diff / D) + 1
	}
}
