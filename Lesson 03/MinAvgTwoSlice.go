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
