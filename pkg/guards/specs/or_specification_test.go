package specification_test

import (
	"gowork/pkg/guards/fixtures"
	specification "gowork/pkg/guards/specs"
	"reflect"
	"testing"
)

func TestNewOrSpecification(t *testing.T) {
	// Definir uma função de chamada para a especificação dummy
	callable := func(candidate specification.Candidate) bool {
		// Implemente sua lógica de avaliação aqui
		// Por exemplo, sempre retornar true para fins de teste
		return true
	}

	// Criar uma especificação dummy usando pacote fixtures
	dummySpec := fixtures.NewDummySpecification(callable)

	// Criar uma lista de especificações para passar para NewOrSpecification
	specs := []specification.Specification[specification.Candidate]{dummySpec}

	// Criar a expectativa de uma instância de OrSpecification
	want := specification.NewOrSpecification(specs...)

	tests := []struct {
		name string
		args struct {
			specs []specification.Specification[specification.Candidate]
		}
		want *specification.OrSpecification[specification.Candidate]
	}{
		{
			name: "Testando a criação da OrSpecification com uma especificação dummy",
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
			if got := specification.NewOrSpecification(tt.args.specs...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewOrSpecification() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestOrSpecification_IsSatisfiedBy(t *testing.T) {
	tests := []struct {
		name   string
		specs []specification.Specification[specification.Candidate]
		candidate   specification.Candidate
		want   bool
	}{
		{
			name:      "satisfied by almost one specifications",
			specs:     []specification.Specification[specification.Candidate]{
				fixtures.NewDummySpecification(
					func(candidate specification.Candidate) bool {
						// Implemente sua lógica de avaliação aqui
						// Por exemplo, sempre retornar true para fins de teste
						return true
					},
				),
				fixtures.NewDummySpecification(
					func(candidate specification.Candidate) bool {
						// Implemente sua lógica de avaliação aqui
						// Por exemplo, sempre retornar true para fins de teste
						return false
					},
				),
			},
			candidate: &struct{}{}, // Substitua pelo candidato real, se necessário
			want:      true,
		},
		{
			name:      "not satisfied by all specifications",
			specs:     []specification.Specification[specification.Candidate]{
				fixtures.NewDummySpecification(
					func(candidate specification.Candidate) bool {
						// Implemente sua lógica de avaliação aqui
						// Por exemplo, sempre retornar true para fins de teste
						return false
					},
				),
				fixtures.NewDummySpecification(
					func(candidate specification.Candidate) bool {
						// Implemente sua lógica de avaliação aqui
						// Por exemplo, sempre retornar true para fins de teste
						return false
					},
				),
			},
			candidate: &struct{}{}, // Substitua pelo candidato real, se necessário
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