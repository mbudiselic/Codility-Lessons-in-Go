// Package solution package provides function to find leader of input array
package solution

// Solution finds leader of provided input array
func Solution(A []int) int {
	// find leader
	size := 0
	var value int
	for i := range A {
		if size == 0 {
			size++
			value = A[i]
		} else {
			if value != A[i] {
				size--
			} else {
				size++
			}
		}
	}
	candidate := -1
	if size > 0 {
		candidate = value
	}
	leader := -1
	count := 0
	for i := range A {
		if A[i] == candidate {
			count++
		}
	}
	if count > len(A)/2 {
		leader = candidate
	}
	// if leader found return its index
	if leader > -1 {
		for i := range A {
			if A[i] == leader {
				return i
			}
		}
	}
	// else return -1
	return -1
}
