package mathutil

import (
	"lab/first/pkg/types"
	"lab/first/pkg/validators"
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
	if check, err := validators.CheckLength(nums); !check {
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
