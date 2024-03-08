package specification

// NotSpecification nega uma especificação existente.
type NotSpecification[T Candidate] struct {
	spec Specification[T]
}

func NewNotSpecification[T Candidate](spec Specification[T]) *NotSpecification[T] {
	return &NotSpecification[T]{spec: spec}
}

// IsSatisfiedBy verifica se a especificação não é satisfeita por um determinado candidato.
// Retorna true se a especificação não for satisfeita, caso contrário, retorna false.
func (s *NotSpecification[T]) IsSatisfiedBy(candidate T) bool {
	return !s.spec.IsSatisfiedBy(candidate)
}
