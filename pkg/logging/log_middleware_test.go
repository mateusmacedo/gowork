package logging

import (
	"reflect"
	"testing"

	"github.com/mateusmacedo/gowork/pkg/logging/fixture"
)

func TestLogMiddleware(t *testing.T) {
	// Test case 1: Function argument is not a function
	t.Run("PanicWhenArgsAreNotFunction", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Errorf("LogMiddleware() did not panic")
			}
		}()

		LogMiddleware(nil, false, 1)
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
				add := LogMiddleware(nil, false, tt.function).(func(...int) int)
				if got := add(tt.args...); got != tt.want {
					t.Errorf("LogMiddleware() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test case 3: Test Add function as goroutine race
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
			tt := tt // Cria uma cópia local da variável para evitar capturar a variável de iteração.
			t.Run(tt.name, func(t *testing.T) {
				add := LogMiddleware(nil, false, tt.function).(func(...int) int)
				done := make(chan bool) // Canal para sincronização.
				go func() {
					defer close(done) // Garante que o canal seja fechado quando a goroutine terminar.
					if got := add(tt.args...); got != tt.want {
						t.Errorf("LogMiddleware() = %v, want %v", got, tt.want)
					}
				}()
				<-done // Espera a goroutine terminar.
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
			{"TestReverseString", fixture.ReverseString, "Hello, World!", "!dlroW ,olleH"},
			{"TestReverseStringStruct", fixture.ReverseStringAdapter(fixture.ReverseStringStruct{}), "Hello, World!", "!dlroW ,olleH"},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				reverse := LogMiddleware(nil, false, tt.function).(func(string) string)
				if got := reverse(tt.args); got != tt.want {
					t.Errorf("LogMiddleware() = %v, want %v", got, tt.want)
				}
			})
		}
	})

	// Test case 5: Test ReverseString function as goroutine race condition
	t.Run("ReverseStringFunctionAsGoroutineTableDrivenSuite", func(t *testing.T) {
		tests := []struct {
			name     string
			function func(string) string
			args     string
			want     string
		}{
			{"TestReverseString", fixture.ReverseString, "Hello, World!", "!dlroW ,olleH"},
			{"TestReverseStringStruct", fixture.ReverseStringAdapter(fixture.ReverseStringStruct{}), "Hello, World!", "!dlroW ,olleH"},
		}

		for _, tt := range tests {
			tt := tt // Cria uma cópia local da variável para evitar capturar a variável de iteração.
			t.Run(tt.name, func(t *testing.T) {
				reverse := LogMiddleware(nil, false, tt.function).(func(string) string)
				done := make(chan bool) // Canal para sincronização.
				go func() {
					defer close(done) // Garante que o canal seja fechado quando a goroutine terminar.
					if got := reverse(tt.args); got != tt.want {
						t.Errorf("LogMiddleware() = %v, want %v", got, tt.want)
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
				processReverse := LogMiddleware(nil, false, tt.function).(func([]string) []string)
				done := make(chan bool) // Canal para sincronização.
				go func() {
					defer close(done) // Ensure the channel is closed when the goroutine finishes
					got := processReverse(tt.args)
					if !reflect.DeepEqual(fixture.SortStringsSlice(got), fixture.SortStringsSlice(tt.want)) {
						t.Errorf("LogMiddleware() = %v, want %v", got, tt.want)
					}
				}()
				<-done // Espera a goroutine terminar.
			})
		}
	})
}

func BenchmarkLogMiddlewareFixtureAdd(b *testing.B) {
	add := LogMiddleware(nil, false, fixture.Add).(func(...int) int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(1, 2, 3, 4, 5)

	}
}

func BenchmarkLogMiddlewareFixtureAddStruct(b *testing.B) {
	add := LogMiddleware(nil, false, fixture.AddAdapter(fixture.AddStruct{})).(func(...int) int)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		add(1, 2, 3, 4, 5)
	}
}

func BenchmarkLogMiddlewareFixtureReverseString(b *testing.B) {
	reverse := LogMiddleware(nil, false, fixture.ReverseString).(func(string) string)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse("Hello, World!")
	}
}

func BenchmarkLogMiddlewareFixtureReverseStringStruct(b *testing.B) {
	reverse := LogMiddleware(nil, false, fixture.ReverseStringAdapter(fixture.ReverseStringStruct{})).(func(string) string)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reverse("Hello, World!")
	}
}

func BenchmarkLogMiddlewareFixtureProcessReverseAsGoroutineWithChan(b *testing.B) {
	processReverse := LogMiddleware(nil, false, fixture.ProcessReverseAsGoroutineWithChan).(func([]string) []string)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		processReverse([]string{"abcde", "fghij", "klmno"})
	}
}

func BenchmarkLogMiddlewareFixtureProcessReverseAsGoroutineWithMutex(b *testing.B) {
	processReverse := LogMiddleware(nil, false, fixture.ProcessReverseAsGoroutineWithMutex).(func([]string) []string)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		processReverse([]string{"abcde", "fghij", "klmno"})
	}
}
