package specification

// OrSpecification combina múltiplas especificações usando a lógica "OU".
type OrSpecification[T Candidate] struct {
	specs []Specification[T]
}

func NewOrSpecification[T Candidate](specs ...Specification[T]) *OrSpecification[T] {
	return &OrSpecification[T]{specs: specs}
}

// IsSatisfiedBy verifica se a especificação é satisfeita por um determinado candidato.
// Retorna true se pelo menos uma das especificações for satisfeita, caso contrário, retorna false.
func (s *OrSpecification[T]) IsSatisfiedBy(candidate T) bool {
	for _, spec := range s.specs {
		if spec.IsSatisfiedBy(candidate) {
			return true
		}
	}
	return false
}
