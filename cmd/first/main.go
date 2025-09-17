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

	squares, cubes := math_struct.Pow()
	fmt.Printf("squares = %v\n", squares)
	fmt.Printf("cubes = %v\n", cubes)

	arr := arr.Arr{Arr: arr.Gen_arr(10)}
	fmt.Printf("Arr = %v\n", arr.Arr)
	fmt.Printf("Len = %v\n", arr.Get_len())
}
