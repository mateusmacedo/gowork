package specification

// AndSpecification combina múltiplas especificações usando a lógica "E".
type AndSpecification[T Candidate] struct {
    specs []Specification[T]
}

func NewAndSpecification[T Candidate](specs ...Specification[T]) *AndSpecification[T] {
    return &AndSpecification[T]{specs: specs}
}
// IsSatisfiedBy verifica se a especificação é satisfeita por um determinado candidato.
// Retorna true se todas as especificações forem satisfeitas, caso contrário, retorna false.
func (s *AndSpecification[T]) IsSatisfiedBy(candidate T) bool {
    for _, spec := range s.specs {
        if !spec.IsSatisfiedBy(candidate) {
            return false
        }
    }
    return true
}

