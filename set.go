package goset

// Set is a custom type that uses a map to store unique elements of any type.
type Set struct {
	elements map[interface{}]struct{}
}

// New creates and returns a new Set.
func New() *Set {
	return &Set{
		elements: make(map[interface{}]struct{}),
	}
}

// Add adds an element to the set.
func (s *Set) Add(element interface{}) {
	s.elements[element] = struct{}{}
}

// AddAll adds all elements from a slice to the set.
func (s *Set) AddAll(elements []interface{}) {
	for _, element := range elements {
		s.Add(element)
	}
}

// AddStrings adds all strings to the set.
func (s *Set) AddStrings(strs ...string) {
	for _, str := range strs {
		s.Add(str)
	}
}

// Remove removes an element from the set.
func (s *Set) Remove(element interface{}) {
	delete(s.elements, element)
}

// Contains checks if an element is in the set.
func (s *Set) Contains(element interface{}) bool {
	_, exists := s.elements[element]
	return exists
}

// Size returns the number of elements in the set.
func (s *Set) Size() int {
	return len(s.elements)
}
