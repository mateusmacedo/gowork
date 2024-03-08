package specification

type AndSpecification[T Candidate] struct {
	specs []Specification[T]
}

func NewAndSpecification[T Candidate](specs ...Specification[T]) *AndSpecification[T] {
	return &AndSpecification[T]{specs: specs}
}

func (s *AndSpecification[T]) IsSatisfiedBy(candidate T) bool {
	for _, spec := range s.specs {
		if !spec.IsSatisfiedBy(candidate) {
			return false
		}
	}
	return true
}
