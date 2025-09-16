package arr

import (
	"fmt"
	"testing"
)

// Проверяем что при передаче ненулевого размера возвращает не пустой слайс
func TestGen_arr(t *testing.T) {
	arr := Arr{Arr: Gen_arr(10)}
	fmt.Println(arr)

	if arr.Arr == nil {
		t.Errorf("Expected 10, got %v", arr.Get_len())
	}
}

// Проверяем что при передаче нулевого размера возвращает nil
func TestGenArrZeroLength(t *testing.T) {
	result := Gen_arr(0)

	if result != nil {
		t.Error("Expected empty slice, got nil")
	}
}
