package operations

import "lab/first/pkg/validators"

// Реализация калькулятора для сложения
type Sum struct{}

func (s Sum) Process(nums []float64) (float64, error) {
	if check, err := validators.CheckUnaryOperation(nums); !check {
		return 0, err
	}

	result := 0.0
	for _, num := range nums {
		result += num
	}
	return result, nil
}
