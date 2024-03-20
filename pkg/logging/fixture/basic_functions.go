package fixture

func Add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

type AddStruct struct{}

func (s AddStruct) Add(numbers ...int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}

func AddAdapter(s AddStruct) func(...int) int {
	return func(numbers ...int) int {
		return s.Add(numbers...)
	}
}