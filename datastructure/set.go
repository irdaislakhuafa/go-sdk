package datastructure

type Set[E comparable] struct {
	Values map[E]bool
}

type SetInterface[E comparable] interface {
	Add(value E)
	IsExists(value E) bool
	Del(value E)
	Slice() []E
}

func NewSet[E comparable](values ...E) SetInterface[E] {
	set := Set[E]{
		Values: map[E]bool{},
	}

	for _, v := range values {
		set.Values[v] = true
	}

	return &set
}

func (s *Set[E]) Add(value E) {
	s.Values[value] = true
}

func (s *Set[E]) IsExists(value E) bool {
	_, isExists := s.Values[value]
	return isExists
}

func (s *Set[E]) Del(value E) {
	delete(s.Values, value)
}

func (s *Set[E]) Slice() []E {
	results := []E{}
	for v := range s.Values {
		results = append(results, v)
	}
	return results
}
