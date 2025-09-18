package types

import "errors"

// Предопределенные ошибки
var (
	ErrInvalidInput    = errors.New("invalid input: empty slice")
	ErrDivisionByZero  = errors.New("division by zero")
	ErrUnknownOperator = errors.New("unknown operator")
)
