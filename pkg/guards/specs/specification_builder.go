package specification

type SpecificationBuilder[T Candidate] interface {
	WithSpecification(spec Specification[T]) SpecificationBuilder[T]
	And(spec Specification[T]) SpecificationBuilder[T]
	Or(spec Specification[T]) SpecificationBuilder[T]
	Not() SpecificationBuilder[T]
	Build() Specification[T]
}

type BaseSpecificationBuilder[T Candidate] struct {
	spec Specification[T]
}

func NewSpecificationBuilder[T Candidate]() SpecificationBuilder[T] {
	return &BaseSpecificationBuilder[T]{}
}

func (b *BaseSpecificationBuilder[T]) WithSpecification(spec Specification[T]) SpecificationBuilder[T] {
	b.spec = spec
	return b
}

func (b *BaseSpecificationBuilder[T]) And(spec Specification[T]) SpecificationBuilder[T] {
	b.spec = NewAndSpecification(b.spec, spec)
	return b
}

func (b *BaseSpecificationBuilder[T]) Or(spec Specification[T]) SpecificationBuilder[T] {
	b.spec = NewOrSpecification(b.spec, spec)
	return b
}

func (b *BaseSpecificationBuilder[T]) Not() SpecificationBuilder[T] {
	b.spec = NewNotSpecification(b.spec)
	return b
}

func (b *BaseSpecificationBuilder[T]) Build() Specification[T] {
	return b.spec
}
