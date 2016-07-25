package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(N int) int {
	// write your code in Go 1.4
	var gapsize, current int
	var previousBit, currentBit int

	for i := 0; i < 32; i++ {
		bitmask := 1 << uint(i)
		if bitmask >= N {
			return gapsize
		}

		if i > 0 {
			previousBit = currentBit
		}
		currentBit = N & bitmask

		if previousBit != 0 && currentBit == 0 {
			current = 1
			continue
		}
		if currentBit == 0 && current != 0 {
			current++
			continue
		}
		if previousBit == 0 && currentBit != 0 {
			if current >= gapsize {
				gapsize = current
			}
			continue
		}
	}
	return gapsize
}
