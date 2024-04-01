package logging

import (
	"testing"

	"github.com/mateusmacedo/gowork/pkg/logging/fixture"
)

func TestLogDecorator(t *testing.T) {
	// Test case 1: Function argument is not a function
	t.Run("PanicWhenArgsAreNotFunction", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("LogDecorator() did not panic")
			}
		}()

		LogDecorator(1)
	})

	// Test case 2: Test Add function
	t.Run("AddFunctionTableDrivenSuite", func(t *testing.T) {
		tests := []struct {
			name     string
			function func(...int) int
			args     []int
			want     int
		}{
			{"TestAdd", fixture.Add, []int{1, 2, 3, 4, 5}, 15},
			{"TestAddStruct", fixture.AddAdapter(fixture.AddStruct{}), []int{1, 2, 3, 4, 5}, 15},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				add := LogDecorator(tt.function)
				if got := add(tt.args...); got != tt.want {
					t.Errorf("LogDecorator() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test case 3: Test Add function as goroutine
	t.Run("AddFunctionAsGoroutineTableDrivenSuite", func(t *testing.T) {
		tests := []struct {
			name     string
			function func(...int) int
			args     []int
			want     int
		}{
			{"TestAdd", fixture.Add, []int{1, 2, 3, 4, 5}, 15},
			{"TestAddStruct", fixture.AddAdapter(fixture.AddStruct{}), []int{1, 2, 3, 4, 5}, 15},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				add := LogDecorator(tt.function)
				done := make(chan bool)
				go func() {
					defer close(done)
					if got := add(tt.args...); got != tt.want {
						t.Errorf("LogDecorator() = %v, want %v", got, tt.want)
					}
				}()
				<-done // Espera a goroutine terminar
			})
		}
	})

	// Test case 4: Test ReverseString function
	t.Run("ReverseStringFunctionTableDrivenSuite", func(t *testing.T) {
		tests := []struct {
			name     string
			function func(string) string
			args     string
			want     string
		}{
			{"TestReverseString", fixture.ReverseString, "abcde", "edcba"},
			{"TestReverseStringStruct", fixture.ReverseStringAdapter(fixture.ReverseStringStruct{}), "abcde", "edcba"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				reverse := LogDecorator(tt.function)
				if got := reverse(tt.args); got != tt.want {
					t.Errorf("LogDecorator() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test case 5: Test ReverseString function as goroutine
	t.Run("ReverseStringFunctionAsGoroutineTableDrivenSuite", func(t *testing.T) {
		tests := []struct {
			name     string
			function func(string) string
			args     string
			want     string
		}{
			{"TestReverseString", fixture.ReverseString, "abcde", "edcba"},
			{"TestReverseStringStruct", fixture.ReverseStringAdapter(fixture.ReverseStringStruct{}), "abcde", "edcba"},
		}

		for _, tt := range tests {
			tt := tt // Cria uma cópia local da variável para evitar capturar a variável de iteração.
			t.Run(tt.name, func(t *testing.T) {
				reverse := LogDecorator(tt.function)
				done := make(chan bool) // Canal para sincronização.
				go func() {
					defer close(done) // Garante que o canal seja fechado quando a goroutine terminar.
					if got := reverse(tt.args); got != tt.want {
						t.Errorf("LogDecorator() = %v, want %v", got, tt.want)
					}
				}()
				<-done // Espera a goroutine terminar.
			})
		}
	})

	// Test case 6: Test ProcessReverseAsGoroutineWithChan function
	t.Run("ProcessReverseAsGoroutineWithChanTableDrivenSuite", func(t *testing.T) {
		tests := []struct {
			name     string
			function func([]string) []string
			args     []string
			want     []string
		}{
			{"TestProcessReverseAsGoroutine", fixture.ProcessReverseAsGoroutineWithChan, []string{"abcde", "fghij", "klmno"}, []string{"edcba", "jihgf", "onmlk"}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				processReverse := LogDecorator(tt.function)
				if got := processReverse(tt.args); len(got) != len(tt.want) {
					t.Errorf("LogDecorator() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test case 7: Test ProcessReverseAsGoroutineWithMutex function
	t.Run("ProcessReverseAsGoroutineWithMutexTableDrivenSuite", func(t *testing.T) {
		tests := []struct {
			name     string
			function func([]string) []string
			args     []string
			want     []string
		}{
			{"TestProcessReverseAsGoroutine", fixture.ProcessReverseAsGoroutineWithMutex, []string{"abcde", "fghij", "klmno"}, []string{"edcba", "jihgf", "onmlk"}},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				processReverse := LogDecorator(tt.function)
				if got := processReverse(tt.args); len(got) != len(tt.want) {
					t.Errorf("LogDecorator() = %v, want %v", got, tt.want)
				}
			})
		}
	})
}

func BenchmarkLogDecoratorFixtureAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogDecorator(fixture.Add)(1, 2, 3, 4, 5)
	}
}

func BenchmarkLogDecoratorFixtureAddStruct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogDecorator(fixture.AddAdapter(fixture.AddStruct{}))(1, 2, 3, 4, 5)
	}
}

func BenchmarkLogDecoratorFixtureAddAsGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := make(chan bool)
		go func() {
			defer close(done)
			LogDecorator(fixture.Add)(1, 2, 3, 4, 5)
		}()
		<-done
	}
}

func BenchmarkLogDecoratorFixtureReverseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogDecorator(fixture.ReverseString)("abcde")
	}
}

func BenchmarkLogDecoratorFixtureReverseStringAsGoroutine(b *testing.B) {
	for i := 0; i < b.N; i++ {
		done := make(chan bool)
		go func() {
			defer close(done)
			LogDecorator(fixture.ReverseString)("abcde")
		}()
		<-done
	}
}

func BenchmarkLogDecoratorFixtureProcessReverseAsGoroutineWithChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogDecorator(fixture.ProcessReverseAsGoroutineWithChan)([]string{"abcde", "fghij", "klmno"})
	}
}

func BenchmarkLogDecoratorFixtureProcessReverseAsGoroutineWithMutex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		LogDecorator(fixture.ProcessReverseAsGoroutineWithMutex)([]string{"abcde", "fghij", "klmno"})
	}
}
