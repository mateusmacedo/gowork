package logging

import (
	"log"
	"reflect"
	"runtime"
	"time"
)

// LogDecorator é uma função genérica que aceita uma função `fn` e retorna uma nova função
// que loga os detalhes da execução de `fn`, incluindo argumentos, valor de retorno, e tempo de execução.
//
// Exemplo de uso:
// ```go
//
//	func soma(a, b int) int {
//	    return a + b
//	}
//
// soma = LogDecorator(soma)
// soma(1, 2) // Saída: "Iniciando a execução de main.soma com os argumentos: [1 2]" e "Função main.soma executada em 1.5µs com o retorno: [3]"
// ```
//
// Exemplo de uso com função anônima:
// ```go
//
//	fn := LogDecorator(func(a, b int) int {
//	    return a + b
//	})
//
// fn(1, 2) // Saída: "Iniciando a execução de main.func1 com os argumentos: [1 2]" e "Função main.func1 executada em 1.5µs com o retorno: [3]"
// ```
//
// Exemplo de uso com metodos de struct:
// ```go
// type Soma struct{}
//
//	func (s Soma) Soma(a, b int) int {
//	    return a + b
//	}
//
//	func SomaAdapter(s Soma) func(int, int) int {
//		return func(a, b int) int {
//			return s.Soma(a, b)
//		}
//	}
//
//	func main() {
//		soma := Soma{}
//
//		somaFunc := AddAdapter(soma) // Transforma o método `Add` em uma função adaptada.
//
//		decoratedSoma := LogDecorator(somaFunc) // Aplica o `LogDecorator` à função adaptada.
//
//		result := decoratedSoma(5, 3) // Usa a função decorada.
//	}
func LogDecorator[F any](fn F) F {
	fnType := reflect.TypeOf(fn)
	if fnType.Kind() != reflect.Func {
		log.Panic("LogDecorator: argumento fornecido não é uma função")
	}

	return reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
		start := time.Now()

		// Log dos argumentos da função
		argValues := make([]interface{}, len(args))
		for i, arg := range args {
			argValues[i] = arg.Interface()
		}
		log.Printf("Iniciando a execução de %s com os argumentos: %v\n", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), argValues)

		// Chamada da função original para cenarios variadicos e não variadicos
		fnAsValue := reflect.ValueOf(fn)
		if fnType.IsVariadic() {
			results = fnAsValue.CallSlice(args)
		} else {
			results = fnAsValue.Call(args)
		}

		// Log do valor de retorno e tempo de execução
		resultValues := make([]interface{}, len(results))
		for i, res := range results {
			resultValues[i] = res.Interface()
		}
		log.Printf("Função %s executada em %s com o retorno: %v\n", runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), time.Since(start), resultValues)

		return
	}).Interface().(F)
}
