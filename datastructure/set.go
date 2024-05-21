package datastructure

type Set[E comparable] struct {
	values map[E]bool
}

type SetInterface[E comparable] interface {
	// Added list element to set.
	Add(values ...E)

	// To check is element is exists on set.
	IsExists(value E) bool

	// Delete list element of set.
	Del(values ...E)

	// Convert set to slice. Maybe you need it for interation.
	Slice() []E

	// Clear list element from set.
	Clear()

	// Check is set is empty. Return true is empty and false is not empty.
	IsEmpty() bool

	// Get size of element on set.
	Size() int

	// Delete element from set if condition is matched.
	DelIf(condition func(value E) bool) SetInterface[E]

	// Filter elements of set. Will return new set with filtered elements.
	Filter(condition func(value E) bool) SetInterface[E]
}

// Set data type is like `slice` or `array` but without duplicate element and set element is unordered element
func NewSet[E comparable](values ...E) SetInterface[E] {
	set := Set[E]{
		values: map[E]bool{},
	}

	for _, v := range values {
		set.values[v] = true
	}

	return &set
}

func (s *Set[E]) Add(values ...E) {
	for _, v := range values {
		s.values[v] = true
	}
}

func (s *Set[E]) IsExists(value E) bool {
	_, isExists := s.values[value]
	return isExists
}

func (s *Set[E]) Del(values ...E) {
	for _, v := range values {
		delete(s.values, v)
	}
}

func (s *Set[E]) Slice() []E {
	results := []E{}
	for v := range s.values {
		results = append(results, v)
	}
	return results
}

func (s *Set[E]) Clear() {
	s.values = nil
}

func (s *Set[E]) IsEmpty() bool {
	return len(s.values) == 0
}

func (s *Set[E]) Size() int {
	return len(s.values)
}

func (s *Set[E]) DelIf(condition func(value E) bool) SetInterface[E] {
	for v := range s.values {
		if condition(v) {
			s.Del(v)
		}
	}

	return s
}

func (s *Set[E]) Filter(condition func(value E) bool) SetInterface[E] {
	var s1 SetInterface[E] = &Set[E]{
		values: map[E]bool{},
	}

	for v := range s.values {
		if condition(v) {
			s1.Add(v)
		}
	}

	return s1
}
