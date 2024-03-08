package specification_test

import (
	"gowork/pkg/guards/fixtures"
	specification "gowork/pkg/guards/specs"
	"reflect"
	"testing"
)

func TestNewNotSpecification(t *testing.T) {
	spec := fixtures.NewDummySpecification(func(candidate specification.Candidate) bool {
		return true
	})

	want := specification.NewNotSpecification(spec)

	tests := []struct {
		name string
		spec specification.Specification[specification.Candidate]
		want *specification.NotSpecification
	}{
		{
			name: "Testando a criação da NotSpecification com uma especificação dummy",
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
