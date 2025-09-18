package mathutil

// Max возвращает максимальное значение из слайса чисел
func Max(nums []float64) (float64, error) {
	if check, err := checkLength(nums); !check {
		return 0, err
	}

	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max, nil
}

// Pow возвращает квадраты и кубы всех чисел (включая отрицательные)
func Pow(nums []float64) ([]float64, []float64, error) {
	if check, err := checkLength(nums); !check {
		return nil, nil, err
	}

	n := len(nums)
	squares := make([]float64, n)
	cubes := make([]float64, n)

	for i := 0; i < n; i++ {
		x := nums[i]
		squares[i] = x * x
		cubes[i] = squares[i] * x
	}

	return squares, cubes, nil
}
