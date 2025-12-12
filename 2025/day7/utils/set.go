package utils
// Generic Set using map[T]struct{}

/*
	var s Set[int]
	s.data = map[int]struct{}{
		10: {},
		20: {},
	}
*/	
type Set[T comparable] struct {
	data map[T]struct{} // represents zero bytes - good for set
}

// create object of Set
func NewSet[T comparable]() *Set[T] {
	return &Set[T]{data: make(map[T]struct{})}
}

func (s *Set[T]) Add(v T) {
	s.data[v] = struct{}{}
}

func (s *Set[T]) Remove(v T) {
	delete(s.data, v)
}

func (s *Set[T]) Has(v T) bool {
	_, ok := s.data[v]
	return ok
}

func (s *Set[T]) Size() int {
	return len(s.data)
}

func (s *Set[T]) Values() []T {
	res := make([]T, 0, len(s.data))
	for k := range s.data {
		res = append(res, k)
	}
	return res
}
