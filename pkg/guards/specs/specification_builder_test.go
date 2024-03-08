package specification_test

import (
	"gowork/pkg/guards/fixtures"
	specification "gowork/pkg/guards/specs"
	"reflect"
	"testing"
)

func TestNewSpecificationBuilder(t *testing.T) {
	tests := []struct {
		name string
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testando a criação de um SpecificationBuilder",
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
	// Definir uma função de chamada para a especificação dummy
	callable := func(candidate any) bool {
		// Implemente sua lógica de avaliação aqui
		// Por exemplo, sempre retornar true para fins de teste
		return true
	}

	// Criar uma especificação dummy usando pacote fixtures
	dummySpec := fixtures.NewDummySpecification(callable)

	// Criar uma SpecificationBuilder
	builder := specification.NewSpecificationBuilder[any]()

	// Criar a expectativa de uma instância de SpecificationBuilder
	want := builder.WithSpecification(dummySpec)

	tests := []struct {
		name string
		args struct {
			spec specification.Specification[any]
		}
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testando a criação de um SpecificationBuilder com uma especificação dummy",
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
	// Definir uma função de chamada para a especificação dummy
	callable := func(candidate any) bool {
		// Implemente sua lógica de avaliação aqui
		// Por exemplo, sempre retornar true para fins de teste
		return true
	}

	// Criar uma especificação dummy usando pacote fixtures
	dummySpec := fixtures.NewDummySpecification(callable)

	// Criar uma SpecificationBuilder
	builder := specification.NewSpecificationBuilder[any]()

	// Criar a expectativa de uma instância de SpecificationBuilder
	want := builder.And(dummySpec)

	tests := []struct {
		name string
		args struct {
			spec specification.Specification[any]
		}
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testando a criação de um SpecificationBuilder com uma especificação dummy",
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
	// Definir uma função de chamada para a especificação dummy
	callable := func(candidate any) bool {
		// Implemente sua lógica de avaliação aqui
		// Por exemplo, sempre retornar true para fins de teste
		return true
	}

	// Criar uma especificação dummy usando pacote fixtures
	dummySpec := fixtures.NewDummySpecification(callable)

	// Criar uma SpecificationBuilder
	builder := specification.NewSpecificationBuilder[any]()

	// Criar a expectativa de uma instância de SpecificationBuilder
	want := builder.Or(dummySpec)

	tests := []struct {
		name string
		args struct {
			spec specification.Specification[any]
		}
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testando a criação de um SpecificationBuilder com uma especificação dummy",
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
	// Definir uma função de chamada para a especificação dummy
	callable := func(candidate any) bool {
		// Implemente sua lógica de avaliação aqui
		// Por exemplo, sempre retornar true para fins de teste
		return true
	}

	// Criar uma especificação dummy usando pacote fixtures
	dummySpec := fixtures.NewDummySpecification(callable)

	// Criar uma SpecificationBuilder
	builder := specification.NewSpecificationBuilder[any]()

	// Criar a expectativa de uma instância de SpecificationBuilder
	want := builder.WithSpecification(dummySpec).Not()

	tests := []struct {
		name string
		want specification.SpecificationBuilder[any]
	}{
		{
			name: "Testando a criação de um SpecificationBuilder com uma especificação dummy",
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
	// Definir uma função de chamada para a especificação dummy
	callable := func(candidate any) bool {
		// Implemente sua lógica de avaliação aqui
		// Por exemplo, sempre retornar true para fins de teste
		return true
	}

	// Criar uma especificação dummy usando pacote fixtures
	dummySpec := fixtures.NewDummySpecification(callable)

	// Criar uma SpecificationBuilder
	builder := specification.NewSpecificationBuilder[any]()

	// Criar a expectativa de uma instância de SpecificationBuilder
	want := builder.WithSpecification(dummySpec).Build()

	tests := []struct {
		name string
		want specification.Specification[any]
	}{
		{
			name: "Testando a criação de um SpecificationBuilder com uma especificação dummy",
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