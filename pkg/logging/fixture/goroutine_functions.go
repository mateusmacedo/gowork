package fixture

import (
	"sort"
	"sync"
)

func SortStringsSlice(slice []string) []string {
	sorted := make([]string, len(slice))
	copy(sorted, slice)
	sort.Strings(sorted)
	return sorted
}

func ReverseString(s string) string {
	var result string
	for _, r := range s {
		result = string(r) + result
	}
	return result
}

type ReverseStringStruct struct{}

func (rs ReverseStringStruct) ReverseString(s string) string {
	return ReverseString(s)
}

func ReverseStringAdapter(rs ReverseStringStruct) func(string) string {
	return func(s string) string {
		return rs.ReverseString(s)
	}
}

func ProcessReverseAsGoroutineWithChan(data []string) []string {
	resultChan := make(chan string, len(data))
	var wg sync.WaitGroup

	for _, d := range data {
		wg.Add(1)
		go func(d string) {
			defer wg.Done()
			resultChan <- ReverseString(d)
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

func ProcessReverseAsGoroutineWithMutex(data []string) []string {
    var result []string
    var wg sync.WaitGroup
    var mutex sync.Mutex

    for _, d := range data {
        wg.Add(1)
        go func(d string) {
            defer wg.Done()
            res := ReverseString(d)
            mutex.Lock()
            result = append(result, res)
            mutex.Unlock()
        }(d)
    }

    wg.Wait()

    return result
}