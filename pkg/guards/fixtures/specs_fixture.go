package fixtures

import (
	specification "gowork/pkg/guards/specs"
)

type CallableCandidateAssertion func(candidate any) bool

type dummySpecification struct {
    Callable CallableCandidateAssertion
}

func (d dummySpecification) IsSatisfiedBy(candidate any) bool {
    return d.Callable(candidate)
}

func NewDummySpecification(callable CallableCandidateAssertion) specification.Specification[any] {
    return &dummySpecification{
        Callable: callable,
    }
}

