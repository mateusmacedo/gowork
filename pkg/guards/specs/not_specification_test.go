package specification_test

import (
	"reflect"
	"testing"

	"github.com/mateusmacedo/gowork/pkg/guards/fixtures"
	specification "github.com/mateusmacedo/gowork/pkg/guards/specs"
)

func TestNewNotSpecification(t *testing.T) {
	spec := fixtures.NewDummySpecification(func(candidate any) bool {
		return true
	})

	want := specification.NewNotSpecification(spec)

	tests := []struct {
		name string
		spec specification.Specification[any]
		want *specification.NotSpecification[any]
	}{
		{
			name: "Create a new not specification",
			spec: spec,
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specification.NewNotSpecification(tt.spec); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewNotSpecification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNotSpecification_IsSatisfiedBy(t *testing.T) {
	tests := []struct {
		name      string
		spec      specification.Specification[any]
		candidate any
		want      bool
	}{
		{
			name: "Satisfied by the not specification",
			spec: fixtures.NewDummySpecification(
				func(candidate any) bool {
					return false
				}),
			candidate: &struct{}{},
			want:      true,
		},
		{
			name: "Not satisfied by the not specification",
			spec: fixtures.NewDummySpecification(
				func(candidate any) bool {
					return true
				}),
			candidate: &struct{}{},
			want:      false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sut := specification.NewNotSpecification(tt.spec)
			if got := sut.IsSatisfiedBy(tt.candidate); got != tt.want {
				t.Errorf("NotSpecification.IsSatisfiedBy() = %v, want %v", got, tt.want)
			}
		})
	}
}
