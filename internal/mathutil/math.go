package mathutil

import (
	"fmt"
)

type Operator string

const (
	MINUS    Operator = "-"
	PLUS     Operator = "+"
	MULTIPLY Operator = "*"
	DIV      Operator = "/"
)

type Math struct {
	Nums     []float64
	Operator Operator
}

func (o Operator) IsValid() bool {
	switch o {
	case MINUS, PLUS, MULTIPLY, DIV:
		return true
	default:
		return false
	}
}

func (m Math) Calculate() (float64, string) {
	if !m.Operator.IsValid() {
		return 0, "Invalid operator"
	}
	fmt.Println(m.Operator)

	var result float64

	switch m.Operator {
	case MULTIPLY:
		result = 1.0
	case PLUS, MINUS, DIV:
		result = 0.0
	}

	for _, num := range m.Nums {
		switch m.Operator {
		case MINUS:
			result -= float64(num)
		case PLUS:
			result += float64(num)
		case MULTIPLY:
			result *= float64(num)
		case DIV:
			if num == 0 {
				return 0, "Division by zero"
			}
			result /= float64(num)
		}
	}

	return result, ""
}

func (m Math) Max() float64 {
	var max float64 = m.Nums[0]

	for i := 0; i < len(m.Nums); i++ {
		if max < m.Nums[i] {
			max = m.Nums[i]
		}
	}

	return max

}

func (m Math) Pow() ([]float64, []float64) {
	n := len(m.Nums)
	pow := make([]float64, n)
	pow3 := make([]float64, n)

	for i := 0; i < n; i++ {
		x := m.Nums[i]
		if x > 0 {
			pow[i] = x * x
			pow3[i] = pow[i] * x
		}
	}

	return pow, pow3
}
