package main

import (
	"fmt"

	"lab/first/internal/arr"
	"lab/first/internal/mathutil"
)

func main() {
	math_struct := mathutil.Math{Nums: []float64{1, 2, 3}, Operator: mathutil.MULTIPLY}
	calc, err := math_struct.Calculate()
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Printf("Calculate res = %v\n", calc)
	}

	fmt.Printf("Max num = %v\n", math_struct.Max())

	pow, pow3 := math_struct.Pow()
	fmt.Printf("Pow = %v\n", pow)
	fmt.Printf("Pow3 = %v\n", pow3)

	arr := arr.Arr{Arr: arr.Gen_arr(10)}
	fmt.Printf("Arr = %v\n", arr.Arr)
	fmt.Printf("Len = %v\n", arr.Get_len())
}
