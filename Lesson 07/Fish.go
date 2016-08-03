package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, B []int) int {
	// write your code in Go 1.4
	var s fishStack
	for i := 0; i < len(A); i++ {
		fishSize := A[i]
		fishDirection := B[i]
		if s.empty() {
			s.push(i)
		} else {
			for !s.empty() && fishDirection-B[s.peek()] == -1 && A[s.peek()] < fishSize {
				s.pop()
			}
			if !s.empty() {
				if fishDirection-B[s.peek()] != -1 {
					s.push(i)
				}
			} else {
				s.push(i)
			}
		}
	}
	return s.len()
}

type fishStack struct{ vec []int }

func (s fishStack) empty() bool { return len(s.vec) == 0 }
func (s fishStack) peek() int   { return s.vec[len(s.vec)-1] }
func (s fishStack) len() int    { return len(s.vec) }
func (s *fishStack) push(i int) { s.vec = append(s.vec, i) }
func (s *fishStack) pop() int {
	d := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return d
}
