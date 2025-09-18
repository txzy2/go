package operations

import (
	"lab/first/internal/types"
	"testing"
)

func TestSum_Process(t *testing.T) {
	tests := []struct {
		name     string
		nums     []float64
		expected float64
		wantErr  bool
	}{
		{
			name:     "valid single number",
			nums:     []float64{5.0},
			expected: 5.0,
			wantErr:  false,
		},
		{
			name:     "valid multiple numbers",
			nums:     []float64{1.0, 2.0, 3.0, 4.0},
			expected: 10.0,
			wantErr:  false,
		},
		{
			name:     "empty slice",
			nums:     []float64{},
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "negative numbers",
			nums:     []float64{-1.0, -2.0, 3.0},
			expected: 0.0,
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := Sum{}
			result, err := s.Process(tt.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("Sum.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Sum.Process() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMinus_Process(t *testing.T) {
	tests := []struct {
		name     string
		nums     []float64
		expected float64
		wantErr  bool
	}{
		{
			name:     "valid two numbers",
			nums:     []float64{10.0, 3.0},
			expected: 7.0,
			wantErr:  false,
		},
		{
			name:     "valid multiple numbers",
			nums:     []float64{10.0, 2.0, 3.0},
			expected: 5.0,
			wantErr:  false,
		},
		{
			name:     "single number - should error",
			nums:     []float64{5.0},
			expected: 0,
			wantErr:  true,
		},
		{
			name:     "empty slice",
			nums:     []float64{},
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Minus{}
			result, err := m.Process(tt.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("Minus.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Minus.Process() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMultiply_Process(t *testing.T) {
	tests := []struct {
		name     string
		nums     []float64
		expected float64
		wantErr  bool
	}{
		{
			name:     "valid single number",
			nums:     []float64{5.0},
			expected: 5.0,
			wantErr:  false,
		},
		{
			name:     "valid multiple numbers",
			nums:     []float64{2.0, 3.0, 4.0},
			expected: 24.0,
			wantErr:  false,
		},
		{
			name:     "with zero",
			nums:     []float64{1.0, 0.0, 3.0},
			expected: 0.0,
			wantErr:  false,
		},
		{
			name:     "empty slice",
			nums:     []float64{},
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := Multiply{}
			result, err := m.Process(tt.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("Multiply.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Multiply.Process() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestDivide_Process(t *testing.T) {
	tests := []struct {
		name     string
		nums     []float64
		expected float64
		wantErr  bool
		errType  error
	}{
		{
			name:     "valid division",
			nums:     []float64{10.0, 2.0},
			expected: 5.0,
			wantErr:  false,
		},
		{
			name:     "valid multiple divisions",
			nums:     []float64{24.0, 2.0, 3.0},
			expected: 4.0,
			wantErr:  false,
		},
		{
			name:     "division by zero",
			nums:     []float64{10.0, 0.0},
			expected: 0,
			wantErr:  true,
			errType:  types.ErrDivisionByZero,
		},
		{
			name:     "single number - should error",
			nums:     []float64{5.0},
			expected: 0,
			wantErr:  true,
			errType:  types.ErrInvalidInput,
		},
		{
			name:     "empty slice",
			nums:     []float64{},
			expected: 0,
			wantErr:  true,
			errType:  types.ErrInvalidInput,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := Divide{}
			result, err := d.Process(tt.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("Divide.Process() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Divide.Process() = %v, want %v", result, tt.expected)
			}
			if tt.wantErr && tt.errType != nil && err != tt.errType {
				t.Errorf("Divide.Process() error = %v, want %v", err, tt.errType)
			}
		})
	}
}
