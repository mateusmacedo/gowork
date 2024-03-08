package specification

type Candidate interface{}

type Specification[T Candidate] interface {
    IsSatisfiedBy(T) bool
}
