/*
A non-empty zero-indexed array A consisting of N integers is given. A pair of integers (P, Q), such that 0 ≤ P < Q < N, is called a slice of array A (notice that the slice contains at least two elements). The average of a slice (P, Q) is the sum of A[P] + A[P + 1] + ... + A[Q] divided by the length of the slice. To be precise, the average equals (A[P] + A[P + 1] + ... + A[Q]) / (Q − P + 1).

For example, array A such that:
    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8

contains the following example slices:

        slice (1, 2), whose average is (2 + 2) / 2 = 2;
        slice (3, 4), whose average is (5 + 1) / 2 = 3;
        slice (1, 4), whose average is (2 + 2 + 5 + 1) / 4 = 2.5.

The goal is to find the starting position of a slice whose average is minimal.

Write a function:

    func Solution(A []int) int

that, given a non-empty zero-indexed array A consisting of N integers, returns the starting position of the slice with the minimal average. If there is more than one slice with a minimal average, you should return the smallest starting position of such a slice.

For example, given array A such that:
    A[0] = 4
    A[1] = 2
    A[2] = 2
    A[3] = 5
    A[4] = 1
    A[5] = 5
    A[6] = 8

the function should return 1, as explained above.

Assume that:

        N is an integer within the range [2..100,000];
        each element of array A is an integer within the range [−10,000..10,000].

Complexity:

        expected worst-case time complexity is O(N);
        expected worst-case space complexity is O(N), beyond input storage (not counting the storage required for input arguments).

Elements of input arrays can be modified.
*/

package solution

// you can also use imports, for example:
//import "fmt"
import "math"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int) int {
	// write your code in Go 1.4
	pref := prefSum(A)
	N := len(A)
	maxIndex := N - 1
	var min float32 = math.MaxFloat32
	var minIndex int
	for i := 0; i < N-1; i++ {
		if i+1 <= maxIndex {
			calcAvg(&pref, i, i+1, &minIndex, &min)
		}
		if i+2 <= maxIndex {
			calcAvg(&pref, i, i+2, &minIndex, &min)
		}
	}
	return minIndex
}

func calcAvg(P *[]int, x, y int, minIndex *int, min *float32) {
	avg := sliceAvg(*P, x, y)
	if avg < *min {
		*min = avg
		*minIndex = x
	}
	if avg == *min && x < *minIndex {
		*minIndex = x
	}
}

func prefSum(A []int) []int {
	n := len(A)
	P := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		P[i] = P[i-1] + A[i-1]
	}
	return P
}

func sliceSum(P []int, x, y int) int {
	return P[y+1] - P[x]
}

func sliceAvg(P []int, x, y int) float32 {
	return float32(sliceSum(P, x, y)) / float32(y-x+1)
}
