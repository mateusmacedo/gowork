package specification_test

import (
	"gowork/pkg/guards/fixtures"
	specification "gowork/pkg/guards/specs"
	"reflect"
	"testing"
)

func TestNewAndSpecification(t *testing.T) {
	specs := []specification.Specification[specification.Candidate]{
		fixtures.NewDummySpecification(func(candidate specification.Candidate) bool {
			return true
		}),
		fixtures.NewDummySpecification(func(candidate specification.Candidate) bool {
			return true
		}),
	}

	want := specification.NewAndSpecification(specs...)

	tests := []struct {
		name  string
		specs []specification.Specification[specification.Candidate]
		want  *specification.AndSpecification[specification.Candidate]
	}{
		{
			name:  "Testando a criação da AndSpecification com uma especificação dummy",
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
		specs     []specification.Specification[specification.Candidate]
		candidate specification.Candidate
		want      bool
	}{
		{
			name: "satisfied by all specifications",
			specs: []specification.Specification[specification.Candidate]{
				fixtures.NewDummySpecification(
					func(candidate specification.Candidate) bool {
						return true
					},
				),
				fixtures.NewDummySpecification(
					func(candidate specification.Candidate) bool {
						return true
					},
				),
			},
			candidate: &struct{}{},
			want:      true,
		},
		{
			name: "not satisfied by all specifications",
			specs: []specification.Specification[specification.Candidate]{
				fixtures.NewDummySpecification(
					func(candidate specification.Candidate) bool {
						return false
					},
				),
				fixtures.NewDummySpecification(
					func(candidate specification.Candidate) bool {
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
