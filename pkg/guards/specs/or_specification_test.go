package specification_test

import (
	"reflect"
	"testing"

	"github.com/mateusmacedo/gowork/pkg/guards/fixtures"
	specification "github.com/mateusmacedo/gowork/pkg/guards/specs"
)

func TestNewOrSpecification(t *testing.T) {
	callable := func(candidate any) bool {
		return true
	}
	dummySpec := fixtures.NewDummySpecification(callable)
	specs := []specification.Specification[any]{dummySpec}
	want := specification.NewOrSpecification(specs...)

	tests := []struct {
		name string
		args struct {
			specs []specification.Specification[any]
		}
		want *specification.OrSpecification[any]
	}{
		{
			name: "Create a new or specification",
			args: struct {
				specs []specification.Specification[any]
			}{
				specs: specs,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specification.NewOrSpecification(tt.args.specs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrSpecification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrSpecification_IsSatisfiedBy(t *testing.T) {
	tests := []struct {
		name      string
		specs     []specification.Specification[any]
		candidate any
		want      bool
	}{
		{
			name: "Satisfied by almost one specifications",
			specs: []specification.Specification[any]{
				fixtures.NewDummySpecification(
					func(candidate any) bool {
						return true
					},
				),
				fixtures.NewDummySpecification(
					func(candidate any) bool {
						return false
					},
				),
			},
			candidate: &struct{}{},
			want:      true,
		},
		{
			name: "Not satisfied by all specifications",
			specs: []specification.Specification[any]{
				fixtures.NewDummySpecification(
					func(candidate any) bool {
						return false
					},
				),
				fixtures.NewDummySpecification(
					func(candidate any) bool {
						return false
					},
				),
			},
			candidate: &struct{}{},
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := specification.NewOrSpecification(tt.specs...)
			if got := sut.IsSatisfiedBy(tt.candidate); got != tt.want {
				t.Errorf("OrSpecification.IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
