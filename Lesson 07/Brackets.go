package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(S string) int {
	// write your code in Go 1.4
	var s bracketStack
	for _, c := range S {
		if s.empty() {
			s.push(c)
		} else {
			if (c == ')' && s.peek() == '(') || (c == '}' && s.peek() == '{') || (c == ']' && s.peek() == '[') {
				s.pop()
			} else if c == '(' || c == '{' || c == '[' {
				s.push(c)
			} else {
				return 0
			}
		}
	}

	if s.empty() {
		return 1
	}
	return 0
}

type bracketStack struct{ vec []rune }

func (s bracketStack) empty() bool  { return len(s.vec) == 0 }
func (s bracketStack) peek() rune   { return s.vec[len(s.vec)-1] }
func (s bracketStack) len() int     { return len(s.vec) }
func (s *bracketStack) push(i rune) { s.vec = append(s.vec, i) }
func (s *bracketStack) pop() rune {
	d := s.vec[len(s.vec)-1]
	s.vec = s.vec[:len(s.vec)-1]
	return d
}
