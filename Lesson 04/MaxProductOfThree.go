package solution

import "sort"

func Solution(A []int) int {
	N := len(A)

	if N == 3 {
		return A[0] * A[1] * A[2]
	}
	sort.Ints(A)
	//largest number is 0
	if A[N-1] == 0 {
		return 0
	} else
	//all negative
	if A[N-1] < 0 {
		return A[N-1] * A[N-2] * A[N-3]
	} else {
		//max is either product of 2 biggest negatives and largest positive
		max := A[0] * A[1] * A[N-1]
		//or product of 3 largest numbers
		if A[N-1]*A[N-2]*A[N-3] > max {
			return A[N-1] * A[N-2] * A[N-3]
		} else {
			return max
		}
	}
}
