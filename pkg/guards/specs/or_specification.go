package specification

type OrSpecification[T Candidate] struct {
	specs []Specification[T]
}

func NewOrSpecification[T Candidate](specs ...Specification[T]) *OrSpecification[T] {
	return &OrSpecification[T]{specs: specs}
}

func (s *OrSpecification[T]) IsSatisfiedBy(candidate T) bool {
	for _, spec := range s.specs {
		if spec.IsSatisfiedBy(candidate) {
			return true
		}
	}
	return false
}
