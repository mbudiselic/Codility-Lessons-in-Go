package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

//int set implementation, found online
type IntSet struct {
	set map[int]bool
}

func NewIntSet() *IntSet {
	return &IntSet{make(map[int]bool)}
}

func (set *IntSet) Add(i int) bool {
	_, found := set.set[i]
	set.set[i] = true
	return !found //False if it existed already
}

func (set *IntSet) Contains(i int) bool {
	_, found := set.set[i]
	return found //true if it existed already
}

func (set *IntSet) Remove(i int) {
	delete(set.set, i)
}

func (set *IntSet) Size() int {
	return len(set.set)
}

//end of set implementation

func Solution(A []int) int {
	// write your code in Go 1.4
	numSet := NewIntSet()
	for _, v := range A {
		numSet.Add(v)
	}
	return numSet.Size()
}
