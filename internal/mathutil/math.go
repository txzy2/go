package mathutil

import (
	"lab/first/internal/types"
)

type Operator string

const (
	MINUS    Operator = "-"
	SUM      Operator = "+"
	MULTIPLY Operator = "*"
	DIVIDE   Operator = "/"
)

type Calculator struct {
	Nums     []float64
	Operator Operator
}

func NewCalculator(nums []float64, operator Operator) (*Calculator, error) {
	if check, err := checkLength(nums); !check {
		return nil, err
	}
	return &Calculator{Nums: nums, Operator: operator}, nil
}

func (c Calculator) Calculate() (float64, error) {
	calculator, exists := calculatorFactory[c.Operator]
	if !exists {
		return 0, types.ErrUnknownOperator
	}

	return calculator.Process(c.Nums)
}

func checkLength(nums []float64) (bool, error) {
	if len(nums) == 0 {
		return false, types.ErrInvalidInput
	}

	return true, nil
}

