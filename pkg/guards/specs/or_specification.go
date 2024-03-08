package specification

// OrSpecification combina múltiplas especificações usando a lógica "OU".
type OrSpecification struct {
	specs []Specification[Candidate]
}

func NewOrSpecification(specs ...Specification[Candidate]) *OrSpecification {
	return &OrSpecification{specs: specs}
}

// IsSatisfiedBy verifica se a especificação é satisfeita por um determinado candidato.
// Retorna true se pelo menos uma das especificações for satisfeita, caso contrário, retorna false.
func (s *OrSpecification) IsSatisfiedBy(candidate Candidate) bool {
	for _, spec := range s.specs {
		if spec.IsSatisfiedBy(candidate) {
			return true
		}
	}
	return false
}
