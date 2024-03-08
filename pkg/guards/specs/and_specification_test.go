package specification_test

import (
	"gowork/pkg/guards/fixtures"
	specification "gowork/pkg/guards/specs"
	"reflect"
	"testing"
)

func TestNewAndSpecification(t *testing.T) {
	// Definir uma função de chamada para a especificação dummy
	callable := func(candidate specification.Candidate) bool {
		// Implemente sua lógica de avaliação aqui
		// Por exemplo, sempre retornar true para fins de teste
		return true
	}

	// Criar uma especificação dummy usando pacote fixtures
	dummySpec := fixtures.NewDummySpecification(callable)

	// Criar uma lista de especificações para passar para NewAndSpecification
	specs := []specification.Specification[specification.Candidate]{dummySpec}

	// Criar a expectativa de uma instância de AndSpecification
	want := specification.NewAndSpecification(specs...)

	tests := []struct {
		name string
		args struct {
			specs []specification.Specification[specification.Candidate]
		}
		want *specification.AndSpecification
	}{
		{
			name: "Testando a criação da AndSpecification com uma especificação dummy",
			args: struct {
				specs []specification.Specification[specification.Candidate]
			}{
				specs: specs,
			},
			want: want,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := specification.NewAndSpecification(tt.args.specs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewAndSpecification() = %v, want %v", got, tt.want)
			}
		})
	}
}