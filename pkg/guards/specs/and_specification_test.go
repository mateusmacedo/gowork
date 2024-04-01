package specification_test

import (
	"reflect"
	"testing"

	"github.com/mateusmacedo/gowork/pkg/guards/fixtures"
	specification "github.com/mateusmacedo/gowork/pkg/guards/specs"
)

func TestNewAndSpecification(t *testing.T) {
	specs := []specification.Specification[any]{
		fixtures.NewDummySpecification(func(candidate any) bool {
			return true
		}),
		fixtures.NewDummySpecification(func(candidate any) bool {
			return true
		}),
	}

	want := specification.NewAndSpecification(specs...)

	tests := []struct {
		name  string
		specs []specification.Specification[any]
		want  *specification.AndSpecification[any]
	}{
		{
			name:  "Create a new and specification",
			specs: specs,
			want:  want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specification.NewAndSpecification(tt.specs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAndSpecification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAndSpecification_IsSatisfiedBy(t *testing.T) {
	tests := []struct {
		name      string
		specs     []specification.Specification[any]
		candidate any
		want      bool
	}{
		{
			name: "Satisfied by all specifications",
			specs: []specification.Specification[any]{
				fixtures.NewDummySpecification(
					func(candidate any) bool {
						return true
					},
				),
				fixtures.NewDummySpecification(
					func(candidate any) bool {
						return true
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
						return true
					},
				),
			},
			candidate: &struct{}{},
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := specification.NewAndSpecification(tt.specs...)
			if got := sut.IsSatisfiedBy(tt.candidate); got != tt.want {
				t.Errorf("AndSpecification.IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
