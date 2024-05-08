package datastructure

type Set[E comparable] struct {
	Values map[E]bool
}

type SetInterface[E comparable] interface {
	Add(values ...E)
	IsExists(value E) bool
	Del(values ...E)
	Slice() []E
	Clear()
	IsEmpty() bool
	Size() int
	DelIf(condition func(value E) bool) SetInterface[E]
}

// Set data type is like `slice` or `array` but without duplicate element and set element is unordered element
func NewSet[E comparable](values ...E) SetInterface[E] {
	set := Set[E]{
		Values: map[E]bool{},
	}

	for _, v := range values {
		set.Values[v] = true
	}

	return &set
}

func (s *Set[E]) Add(values ...E) {
	for _, v := range values {
		s.Values[v] = true
	}
}

func (s *Set[E]) IsExists(value E) bool {
	_, isExists := s.Values[value]
	return isExists
}

func (s *Set[E]) Del(values ...E) {
	for _, v := range values {
		delete(s.Values, v)
	}
}

func (s *Set[E]) Slice() []E {
	results := []E{}
	for v := range s.Values {
		results = append(results, v)
	}
	return results
}

func (s *Set[E]) Clear() {
	clear(s.Values)
}

func (s *Set[E]) IsEmpty() bool {
	return len(s.Values) == 0
}

func (s *Set[E]) Size() int {
	return len(s.Values)
}

func (s *Set[E]) DelIf(condition func(value E) bool) SetInterface[E] {
	for v := range s.Values {
		if condition(v) {
			s.Del(v)
		}
	}

	return s
}
