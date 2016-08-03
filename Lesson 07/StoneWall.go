package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(H []int) int {
	// write your code in Go 1.4
	var wall wallStack
	bricks := 0
	for _, v := range H {
		if wall.empty() {
			wall.push(v)
		} else {
			for !wall.empty() && v < wall.peek() {
				bricks++
				wall.pop()
			}
			if wall.empty() {
				wall.push(v)
			}
			if !wall.empty() && v > wall.peek() {
				wall.push(v)
			}
		}
	}
	return bricks + wall.len()
}

type wallStack struct{ vec []int }

func (s wallStack) empty() bool { return len(s.vec) == 0 }
func (s wallStack) peek() int   { return s.vec[len(s.vec)-1] }
func (s wallStack) len() int    { return len(s.vec) }
func (s *wallStack) push(i int) { s.vec = append(s.vec, i) }
func (s *wallStack) pop() int {
	d := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return d
}
