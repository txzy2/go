package mathutil

import (
	"testing"
)

func TestNewCalculator(t *testing.T) {
	tests := []struct {
		name     string
		nums     []float64
		operator Operator
		wantErr  bool
	}{
		{
			name:     "valid calculator",
			nums:     []float64{1.0, 2.0, 3.0},
			operator: SUM,
			wantErr:  false,
		},
		{
			name:     "empty slice",
			nums:     []float64{},
			operator: SUM,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc, err := NewCalculator(tt.nums, tt.operator)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewCalculator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && calc == nil {
				t.Errorf("NewCalculator() returned nil calculator")
			}
		})
	}
}

func TestCalculator_Calculate(t *testing.T) {
	tests := []struct {
		name     string
		nums     []float64
		operator Operator
		expected float64
		wantErr  bool
	}{
		{
			name:     "addition",
			nums:     []float64{1.0, 2.0, 3.0},
			operator: SUM,
			expected: 6.0,
			wantErr:  false,
		},
		{
			name:     "subtraction",
			nums:     []float64{10.0, 2.0, 3.0},
			operator: MINUS,
			expected: 5.0,
			wantErr:  false,
		},
		{
			name:     "multiplication",
			nums:     []float64{2.0, 3.0, 4.0},
			operator: MULTIPLY,
			expected: 24.0,
			wantErr:  false,
		},
		{
			name:     "division",
			nums:     []float64{24.0, 2.0, 3.0},
			operator: DIVIDE,
			expected: 4.0,
			wantErr:  false,
		},
		{
			name:     "unknown operator",
			nums:     []float64{1.0, 2.0},
			operator: Operator("^"),
			expected: 0,
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc, err := NewCalculator(tt.nums, tt.operator)
			if err != nil {
				t.Errorf("NewCalculator() error = %v", err)
				return
			}
			
			result, err := calc.Calculate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Calculator.Calculate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Calculator.Calculate() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestMax(t *testing.T) {
	tests := []struct {
		name     string
		nums     []float64
		expected float64
		wantErr  bool
	}{
		{
			name:     "valid numbers",
			nums:     []float64{1.0, 5.0, 3.0, 9.0, 2.0},
			expected: 9.0,
			wantErr:  false,
		},
		{
			name:     "single number",
			nums:     []float64{42.0},
			expected: 42.0,
			wantErr:  false,
		},
		{
			name:     "negative numbers",
			nums:     []float64{-1.0, -5.0, -3.0},
			expected: -1.0,
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
			result, err := Max(tt.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("Max() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && result != tt.expected {
				t.Errorf("Max() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestPow(t *testing.T) {
	tests := []struct {
		name      string
		nums      []float64
		expectedS []float64
		expectedC []float64
		wantErr   bool
	}{
		{
			name:      "positive numbers",
			nums:      []float64{2.0, 3.0},
			expectedS: []float64{4.0, 9.0},
			expectedC: []float64{8.0, 27.0},
			wantErr:   false,
		},
		{
			name:      "negative numbers",
			nums:      []float64{-2.0, -3.0},
			expectedS: []float64{4.0, 9.0},
			expectedC: []float64{-8.0, -27.0},
			wantErr:   false,
		},
		{
			name:      "mixed numbers",
			nums:      []float64{1.0, -2.0, 0.0},
			expectedS: []float64{1.0, 4.0, 0.0},
			expectedC: []float64{1.0, -8.0, 0.0},
			wantErr:   false,
		},
		{
			name:      "empty slice",
			nums:      []float64{},
			expectedS: nil,
			expectedC: nil,
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			squares, cubes, err := Pow(tt.nums)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				if len(squares) != len(tt.expectedS) || len(cubes) != len(tt.expectedC) {
					t.Errorf("Pow() returned slices of wrong length")
					return
				}
				for i := range squares {
					if squares[i] != tt.expectedS[i] {
						t.Errorf("Pow() squares[%d] = %v, want %v", i, squares[i], tt.expectedS[i])
					}
					if cubes[i] != tt.expectedC[i] {
						t.Errorf("Pow() cubes[%d] = %v, want %v", i, cubes[i], tt.expectedC[i])
					}
				}
			}
		})
	}
}

