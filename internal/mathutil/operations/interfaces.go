package operations

// Интерфейс калькулятора для операций
type Calculator interface {
	Process(nums []float64) (float64, error)
}
