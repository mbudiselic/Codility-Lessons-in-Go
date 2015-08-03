/*
A non-empty zero-indexed array A consisting of N integers is given. Array A represents numbers on a tape.
Any integer P, such that 0 < P < N, splits this tape into two non-empty parts: A[0], A[1], ..., A[P − 1] and A[P], A[P + 1], ..., A[N − 1].
The difference between the two parts is the value of: |(A[0] + A[1] + ... + A[P − 1]) − (A[P] + A[P + 1] + ... + A[N − 1])|
In other words, it is the absolute difference between the sum of the first part and the sum of the second part.
For example, consider array A such that:
  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3
We can split this tape in four places:
P = 1, difference = |3 − 10| = 7
P = 2, difference = |4 − 9| = 5
P = 3, difference = |6 − 7| = 1
P = 4, difference = |10 − 3| = 7
Write a function:
func Solution(A []int) int
that, given a non-empty zero-indexed array A of N integers, returns the minimal difference that can be achieved.
For example, given:
  A[0] = 3
  A[1] = 1
  A[2] = 2
  A[3] = 4
  A[4] = 3
the function should return 1, as explained above.
Assume that:
N is an integer within the range [2..100,000];
each element of array A is an integer within the range [−1,000..1,000].
Complexity:
expected worst-case time complexity is O(N);
expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).
Elements of input arrays can be modified.
*/

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
