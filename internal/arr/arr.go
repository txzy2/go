package arr

import "math/rand"

type Arr struct {
	Arr []float64
}

func (arr Arr) Get_len() int {
	return len(arr.Arr)
}

func Gen_arr(len uint8) []float64 {
	if len == 0 {
		return nil
	}

	arr := []float64{}

	for i := uint8(0); i < len; i++ {
		arr = append(arr, float64(rand.Intn(100)))
	}

	return arr
}
