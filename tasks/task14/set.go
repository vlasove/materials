// Package task14 Имеется последовательность
// строк - (cat, cat, dog, cat, tree)
// создать для нее собственное множество.
package task14

import "fmt"

var (
	words = []string{"cat", "cat", "dog", "cat", "tree"}
)

// Set ...
type Set struct {
	elems []string
}

// New ...
func New(elements ...string) *Set {
	s := &Set{}
	s.elems = append(s.elems, elements...)
	s.sanitize()
	return s
}

// sanitize ...
func (s *Set) sanitize() {
	nodups := []string{}
	encountered := make(map[string]bool)
	for _, element := range s.elems {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	s.elems = nodups
}

// String ...
func (s *Set) String() string {
	return fmt.Sprintf("Set(%v)", s.elems)
}

// Add ...
func (s *Set) Add(word string) {
	s.elems = append(s.elems, word)
	s.sanitize()
}

// AddMany ...
func (s *Set) AddMany(words ...string) {
	s.elems = append(s.elems, words...)
	s.sanitize()
}

// Size ...
func (s *Set) Size() int {
	return len(s.elems)
}

// Start ...
func Start() {
	s := New(words[:]...)
	s.Add("elephant")
	s.AddMany("one", "two", "three", "three", "five")
	fmt.Println(s)
}
