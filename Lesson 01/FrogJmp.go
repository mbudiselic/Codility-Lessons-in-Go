/*
A small frog wants to get to the other side of the road. The frog is currently located at position X and wants to get to a position greater than or equal to Y. The small frog always jumps a fixed distance, D.
Count the minimal number of jumps that the small frog must perform to reach its target.
Write a function:
func Solution(X int, Y int, D int) int
that, given three integers X, Y and D, returns the minimal number of jumps from position X to a position equal to or greater than Y.
For example, given:
  X = 10
  Y = 85
  D = 30
the function should return 3, because the frog will be positioned as follows:
after the first jump, at position 10 + 30 = 40
after the second jump, at position 10 + 30 + 30 = 70
after the third jump, at position 10 + 30 + 30 + 30 = 100
Assume that:
X, Y and D are integers within the range [1..1,000,000,000];
X ≤ Y.
Complexity:
expected worst-case time complexity is O(1);
expected worst-case space complexity is O(1).
*/

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
