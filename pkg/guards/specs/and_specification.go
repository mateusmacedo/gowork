package specification

// AndSpecification combina múltiplas especificações usando a lógica "E".
type AndSpecification struct {
    specs []Specification[Candidate]
}

func NewAndSpecification(specs ...Specification[Candidate]) *AndSpecification {
    return &AndSpecification{specs: specs}
}

// IsSatisfiedBy verifica se a especificação é satisfeita por um determinado candidato.
// Retorna true se todas as especificações forem satisfeitas, caso contrário, retorna false.
func (s *AndSpecification) IsSatisfiedBy(candidate Candidate) bool {
    for _, spec := range s.specs {
        if !spec.IsSatisfiedBy(candidate) {
            return false
        }
    }
    return true
}
