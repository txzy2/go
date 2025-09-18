package arr

import "math/rand"

type Arr struct {
	Value []float64
}

func NewArr(Value []float64) *Arr {
	return &Arr{Value: Value}
}

func (arr Arr) Get_len() int {
	return len(arr.Value)
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
