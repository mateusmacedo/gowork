package specification

type NotSpecification[T Candidate] struct {
	spec Specification[T]
}

func NewNotSpecification[T Candidate](spec Specification[T]) *NotSpecification[T] {
	return &NotSpecification[T]{spec: spec}
}

func (s *NotSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return !s.spec.IsSatisfiedBy(candidate)
}
