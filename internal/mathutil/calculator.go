package mathutil

import "lab/first/internal/mathutil/operations"

// Фабрика калькуляторов
var calculatorFactory = map[Operator]operations.Calculator{
	SUM:      operations.Sum{},
	MINUS:    operations.Minus{},
	MULTIPLY: operations.Multiply{},
	DIVIDE:   operations.Divide{},
}
