package operations

// Реализация калькулятора для вычитания
type Minus struct{}

func (m Minus) Process(nums []float64) (float64, error) {
	if check, err := checkBinaryOperation(nums); !check {
		return 0, err
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result -= nums[i]
	}
	return result, nil
}
