package mathutil

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

	if len(m.Nums) == 0 {
		return 0, "Invalid nums"
	}

	var result float64 = m.Nums[0]

	for _, num := range m.Nums {
		switch m.Operator {
		case MINUS:
			result -= num
		case PLUS:
			result += num
		case MULTIPLY:
			result *= num
		case DIV:
			if num == 0 {
				return 0, "Division by zero"
			}
			result /= num
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
	squares := make([]float64, n)
	cubes := make([]float64, n)

	for i := 0; i < n; i++ {
		x := m.Nums[i]
		if x > 0 {
			squares[i] = x * x
			cubes[i] = squares[i] * x
		}
	}

	return squares, cubes
}
