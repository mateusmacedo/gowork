package specification

// NotSpecification nega uma especificação existente.
type NotSpecification struct {
	spec Specification[Candidate]
}

func NewNotSpecification(spec Specification[Candidate]) *NotSpecification {
	return &NotSpecification{spec: spec}
}

// IsSatisfiedBy verifica se a especificação não é satisfeita por um determinado candidato.
// Retorna true se a especificação não for satisfeita, caso contrário, retorna false.
func (s *NotSpecification) IsSatisfiedBy(candidate Candidate) bool {
	return !s.spec.IsSatisfiedBy(candidate)
}
