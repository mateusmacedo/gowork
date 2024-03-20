package logging

import (
	"context"
	"sync"
	"testing"

	"go.uber.org/zap"
)

func add(numbers ...int) int {
	Console(context.Background(), false).Info("Iniciando a execução da função", zap.String("nome_da_funcao", "add"), zap.Any("argumentos", numbers))
	result := 0
	for _, number := range numbers {
		result += number
	}
	Console(context.Background(), false).Info("Finalizando a execução da função", zap.String("nome_da_funcao", "add"), zap.Any("resultado", result))
	return result
}

type addStruct struct{}

func (s addStruct) Add(numbers ...int) int {
	Console(context.Background(), false).Info("Iniciando a execução da função", zap.String("nome_da_funcao", "add"), zap.Any("argumentos", numbers))
	result := 0
	for _, number := range numbers {
		result += number
	}
	Console(context.Background(), false).Info("Finalizando a execução da função", zap.String("nome_da_funcao", "add"), zap.Any("resultado", result))
	return result
}

func reverseString(s string) string {
	Console(context.Background(), false).Info("Iniciando a execução da função", zap.String("nome_da_funcao", "reverseString"), zap.String("argumentos", s))
	var result string
	for _, r := range s {
		result = string(r) + result
	}
	Console(context.Background(), false).Info("Finalizando a execução da função", zap.String("nome_da_funcao", "reverseString"), zap.String("resultado", result))
	return result
}

type reverseStringStruct struct{}

func (rs reverseStringStruct) ReverseString(s string) string {
	return reverseString(s)
}

func reverseStringAdapter(rs reverseStringStruct) func(string) string {
	return func(s string) string {
		return rs.ReverseString(s)
	}
}

func processReverseAsGoroutineWithChan(data []string) []string {
	resultChan := make(chan string, len(data))
	var wg sync.WaitGroup

	for _, d := range data {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			resultChan <- reverseString(d)
		}(d)
	}

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	var result []string
	for res := range resultChan {
		result = append(result, res)
	}

	return result
}

func processReverseAsGoroutineWithMutex(data []string) []string {
    var result []string
    var wg sync.WaitGroup
    var mutex sync.Mutex

    for _, d := range data {
        wg.Add(1)
        go func(d string) {
            defer wg.Done()
            res := reverseString(d)
            mutex.Lock()
            result = append(result, res)
            mutex.Unlock()
        }(d)
    }

    wg.Wait()

    return result
}

func Benchmark_Console(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Console(context.Background(), false)
	}
}

func Benchmark_ConsoleWithMonitoring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Console(context.Background(), true)
	}
}

func Benchmark_ConsoleLog(b *testing.B) {
	logger := Console(context.Background(), false)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("test")
	}
}

func Benchmark_ConsoleLogWithMonitoring(b *testing.B) {
	logger := Console(context.Background(), true)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		logger.Info("test")
	}
}

func Benchmark_Add(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2, 3, 4, 5)
	}
}

func Benchmark_AddStruct(b *testing.B) {
	s := addStruct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		s.Add(1, 2, 3, 4, 5)
	}
}

func Benchmark_ReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reverseString("Hello, World!")
	}
}

func Benchmark_ReverseStringStruct(b *testing.B) {
	rs := reverseStringStruct{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		rs.ReverseString("Hello, World!")
	}
}

func Benchmark_ProcessReverseAsGoroutineWithChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processReverseAsGoroutineWithChan([]string{"abcde", "fghij", "klmno"})
	}
}

func Benchmark_ProcessReverseAsGoroutineWithMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		processReverseAsGoroutineWithMutex([]string{"abcde", "fghij", "klmno"})
	}
}

