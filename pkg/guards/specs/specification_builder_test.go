package specification_test

import (
	"reflect"
	"testing"

	"github.com/mateusmacedo/gowork/pkg/guards/fixtures"
	specification "github.com/mateusmacedo/gowork/pkg/guards/specs"
)

func TestNewSpecificationBuilder(t *testing.T) {
	tests := []struct {
		name string
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testing the creation of a new SpecificationBuilder",
			want: &specification.BaseSpecificationBuilder[any]{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specification.NewSpecificationBuilder[any](); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSpecificationBuilder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseSpecificationBuilder_BuildWithSpecification(t *testing.T) {
	callable := func(candidate any) bool {
		return true
	}

	dummySpec := fixtures.NewDummySpecification(callable)
	builder := specification.NewSpecificationBuilder[any]()
	want := builder.WithSpecification(dummySpec)

	tests := []struct {
		name string
		args struct {
			spec specification.Specification[any]
		}
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testing the WithSpecification composition",
			args: struct {
				spec specification.Specification[any]
			}{
				spec: dummySpec,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.WithSpecification(tt.args.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseSpecificationBuilder.WithSpecification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseSpecificationBuilder_BuildAnd(t *testing.T) {
	callable := func(candidate any) bool {
		return true
	}

	dummySpec := fixtures.NewDummySpecification(callable)
	builder := specification.NewSpecificationBuilder[any]()
	want := builder.And(dummySpec)

	tests := []struct {
		name string
		args struct {
			spec specification.Specification[any]
		}
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testing the And composition",
			args: struct {
				spec specification.Specification[any]
			}{
				spec: dummySpec,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.And(tt.args.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseSpecificationBuilder.And() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseSpecificationBuilder_BuildOr(t *testing.T) {
	callable := func(candidate any) bool {
		return true
	}

	dummySpec := fixtures.NewDummySpecification(callable)

	builder := specification.NewSpecificationBuilder[any]()

	want := builder.Or(dummySpec)

	tests := []struct {
		name string
		args struct {
			spec specification.Specification[any]
		}
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testing the Or composition",
			args: struct {
				spec specification.Specification[any]
			}{
				spec: dummySpec,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.Or(tt.args.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseSpecificationBuilder.Or() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseSpecificationBuilder_BuildNot(t *testing.T) {
	callable := func(candidate any) bool {
		return true
	}
	dummySpec := fixtures.NewDummySpecification(callable)
	builder := specification.NewSpecificationBuilder[any]()
	want := builder.WithSpecification(dummySpec).Not()

	tests := []struct {
		name string
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testing the Not composition",
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.Not(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseSpecificationBuilder.Not() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBaseSpecificationBuilder_Build(t *testing.T) {
	callable := func(candidate any) bool {
		return true
	}
	dummySpec := fixtures.NewDummySpecification(callable)
	builder := specification.NewSpecificationBuilder[any]()
	want := builder.WithSpecification(dummySpec).Build()

	tests := []struct {
		name string
		want specification.Specification[any]
	}{
		{
			name: "Testing the build",
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := builder.Build(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BaseSpecificationBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
