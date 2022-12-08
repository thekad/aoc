package stack

import "fmt"

// RuneStack is a naive stack implementation of type rune
type RuneStack struct {
	elements []rune
}

// NewRuneStack is the constructor
func NewRuneStack() *RuneStack {
	rs := RuneStack{
		elements: []rune{},
	}

	return &rs
}

// Push is a naive push implementation
func (s *RuneStack) Push(e rune) {
	s.elements = append(s.elements, e)
}

// IsEmpty returns true if the stack is empty
func (s *RuneStack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Pop is a naive pop implementation
func (s *RuneStack) Pop() *rune {
	if s.IsEmpty() {
		return nil
	}
	r := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return &r
}

// PopMulti returns a slice of how many items from the top of the stack
func (s *RuneStack) PopMulti(n int) ([]rune, error) {
	r := []rune{}

	if n > len(s.elements) {
		return r, fmt.Errorf("Can't pop %d when len is %d", n, len(s.elements))
	}

	idx := len(s.elements) - n
	r = append(r, s.elements[idx:]...)
	s.elements = s.elements[:idx]

	return r, nil
}

// Extend will append the given array to the existing array
func (s *RuneStack) Extend(add []rune) {
	s.elements = append(s.elements, add...)
}

func (s *RuneStack) String() string {
	var r []string
	for _, e := range s.elements {
		r = append(r, string(e))
	}
	return fmt.Sprintf("%v", r)
}

// StringStack is a naive stack implementation of type string
type StringStack struct {
	elements []string
}

// NewStringStack is the constructor
func NewStringStack() *StringStack {
	rs := StringStack{
		elements: []string{},
	}

	return &rs
}

// Push is a naive push implementation
func (s *StringStack) Push(e string) {
	s.elements = append(s.elements, e)
}

// IsEmpty returns true if the stack is empty
func (s *StringStack) IsEmpty() bool {
	return len(s.elements) == 0
}

// Pop is a naive pop implementation
func (s *StringStack) Pop() *string {
	if s.IsEmpty() {
		return nil
	}
	r := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]

	return &r
}

// PopMulti returns a slice of how many items from the top of the stack
func (s *StringStack) PopMulti(n int) ([]string, error) {
	r := []string{}

	if n > len(s.elements) {
		return r, fmt.Errorf("Can't pop %d when len is %d", n, len(s.elements))
	}

	idx := len(s.elements) - n
	r = append(r, s.elements[idx:]...)
	s.elements = s.elements[:idx]

	return r, nil
}

// Extend will append the given array to the existing array
func (s *StringStack) Extend(add []string) {
	s.elements = append(s.elements, add...)
}

func (s *StringStack) String() string {
	var r []string
	for _, e := range s.elements {
		r = append(r, string(e))
	}
	return fmt.Sprintf("%v", r)
}

// All returns all the elements in the stack
func (s *StringStack) All() []string {
	return s.elements
}
