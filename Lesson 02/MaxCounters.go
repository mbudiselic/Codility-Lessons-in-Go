package solution

// you can also use imports, for example:
// import "fmt"

// you can use fmt.Println for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int, A []int) []int {
	// write your code in Go 1.4
	counters := make([]int, N)
	var max, lastMax int
	for _, v := range A {
		if v <= N {
			if counters[v-1] < lastMax {
				counters[v-1] = lastMax
			}
			counters[v-1]++
			if counters[v-1] > max {
				max = counters[v-1]
			}
		} else {
			lastMax = max
		}
	}

	for i, v := range counters {
		if v < lastMax {
			counters[i] = lastMax
		}
	}

	return counters
}
