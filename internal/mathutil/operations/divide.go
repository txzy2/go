package operations

import (
	"lab/first/internal/types"
)

// Реализация калькулятора для деления
type Divide struct{}

func (d Divide) Process(nums []float64) (float64, error) {
	if check, err := checkBinaryOperation(nums); !check {
		return 0, err
	}

	result := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0, types.ErrDivisionByZero
		}
		result /= nums[i]
	}
	return result, nil
}
