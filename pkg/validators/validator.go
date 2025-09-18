package validators

import "lab/first/pkg/types"

func CheckLength(nums []float64) (bool, error) {
	if len(nums) == 0 {
		return false, types.ErrInvalidInput
	}

	return true, nil
}

// checkBinaryOperation проверяет, что для бинарных операций (деление, вычитание) есть минимум 2 числа
func CheckBinaryOperation(nums []float64) (bool, error) {
	if len(nums) < 2 {
		return false, types.ErrInvalidInput
	}
	return true, nil
}

// checkUnaryOperation проверяет, что для унарных операций (сложение, умножение) есть минимум 1 число
func CheckUnaryOperation(nums []float64) (bool, error) {
	if len(nums) == 0 {
		return false, types.ErrInvalidInput
	}
	return true, nil
}
