package solution

import "math"

func Solution(A []int) int {
	N := len(A)
	max := math.MinInt32
	for P := 0; P < N-2; P++ {
		for Q := P + 1; Q < N-1; Q++ {
			for R := Q + 1; R < N; R++ {
				if A[P]*A[Q]*A[R] > max {
					max = A[P] * A[Q] * A[R]
				}
			}
		}
	}
	return max
}
