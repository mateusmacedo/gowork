package fixtures

import (
	specification "gowork/pkg/guards/specs"
)

type CallableCandidateAssertion func(candidate specification.Candidate) bool

type dummySpecification struct {
    Callable CallableCandidateAssertion
}

func (d dummySpecification) IsSatisfiedBy(candidate specification.Candidate) bool {
    // Use a função de chamada para avaliar o candidato
    return d.Callable(candidate)
}

func NewDummySpecification(callable CallableCandidateAssertion) specification.Specification[specification.Candidate] {
    return &dummySpecification{
        Callable: callable,
    }
}

