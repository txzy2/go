package arr

import "math/rand"

type Arr struct {
	Arr []int
}

func (arr Arr) Get_len() int {
	return len(arr.Arr)
}

func Gen_arr(len uint8) []int {
	if len == 0 {
		return nil
	}

	arr := []int{}

	for i := uint8(0); i < len; i++ {
		arr = append(arr, rand.Intn(100))
	}

	return arr
}
