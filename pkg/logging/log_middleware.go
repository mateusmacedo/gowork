package logging

import (
	"context"
	"reflect"
	"runtime"
	"time"

	"go.uber.org/zap"
)

type Loggable interface {
	Invoke(args []reflect.Value) []reflect.Value
}

type FuncWrapper struct {
	fn reflect.Value
}

func NewFuncWrapper(fn interface{}) *FuncWrapper {
	if reflect.TypeOf(fn).Kind() != reflect.Func {
		panic("NewFuncWrapper: argument must be a function")
	}
	return &FuncWrapper{fn: reflect.ValueOf(fn)}
}

func (f *FuncWrapper) Invoke(args ...interface{}) []reflect.Value {
	reflectedArgs := make([]reflect.Value, len(args))
	for i, arg := range args {
		reflectedArgs[i] = reflect.ValueOf(arg)
	}
	return f.fn.Call(reflectedArgs)
}

func LogMiddleware(ctx context.Context, monitoring bool, f interface{}) interface{} {
	fnVal := reflect.ValueOf(f)
	fnType := fnVal.Type()

	// Verifica se a função é válida.
	if fnType.Kind() != reflect.Func {
		panic("LogMiddleware: o argumento fornecido não é uma função")
	}

	wrappedFn := reflect.MakeFunc(fnType, func(args []reflect.Value) (results []reflect.Value) {
		// Prepara o logger.
		logger := Console(ctx, monitoring) // Assume-se que Console é uma função que retorna um *zap.Logger.
		funcName := runtime.FuncForPC(fnVal.Pointer()).Name()

		// Log dos argumentos da função
		argValues := make([]interface{}, len(args))
		for i, arg := range args {
			argValues[i] = arg.Interface()
		}
		logger.Info("Iniciando a execução da função", zap.String("nome_da_funcao", funcName), zap.Any("argumentos", argValues))

		start := time.Now()

		// Tratamento especial para funções variádicas.
		if fnType.IsVariadic() {
			results = fnVal.CallSlice(args)
		} else {
			results = fnVal.Call(args)
		}

		// Log do valor de retorno e tempo de execução
		resultValues := make([]interface{}, len(results))
		for i, res := range results {
			resultValues[i] = res.Interface()
		}
		logger.Info("Função executada", zap.String("nome_da_funcao", funcName), zap.Duration("tempo_de_execucao", time.Since(start)), zap.Any("retorno", resultValues))

		return
	})

	return wrappedFn.Interface()
}
