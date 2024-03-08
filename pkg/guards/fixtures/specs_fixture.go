package fixtures

import (
	specification "gowork/pkg/guards/specs"
)

type CallableCandidateAssertion func(candidate any) bool

type dummySpecification struct {
    Callable CallableCandidateAssertion
}

func (d dummySpecification) IsSatisfiedBy(candidate any) bool {
    // Use a função de chamada para avaliar o candidato
    return d.Callable(candidate)
}

func NewDummySpecification(callable CallableCandidateAssertion) specification.Specification[any] {
    return &dummySpecification{
        Callable: callable,
    }
}

