package solution

func Solution(S string, P []int, Q []int) []int {
	genes := []string{"A", "C", "G", "T"}
	n := len(S)
	m := len(P)
	counters := make([][]int, 3)
	for i := range counters {
		counters[i] = make([]int, n+1)
	}

	chars := []rune(S)
	for i := 0; i < 3; i++ {
		for j := 0; j < n; j++ {
			if (string(chars[j])) == genes[i] {
				counters[i][j+1] = counters[i][j] + 1
			} else {
				counters[i][j+1] = counters[i][j]
			}
		}
	}

	result := make([]int, m)
	for i := 0; i < m; i++ {
		from := P[i]
		to := Q[i] + 1
		if counters[0][to]-counters[0][from] > 0 {
			result[i] = 1
		} else if counters[1][to]-counters[1][from] > 0 {
			result[i] = 2
		} else if counters[2][to]-counters[2][from] > 0 {
			result[i] = 3
		} else {
			result[i] = 4
		}
	}
	return result
}
