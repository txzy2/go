package main

import (
	"fmt"

	"lab/first/internal/arr"
	"lab/first/internal/mathutil"
)

func main() {
	fmt.Println("")
	fmt.Println("[== Math ==]")

	math_struct := mathutil.Math{Nums: arr.Gen_arr(10), Operator: mathutil.MINUS}
	fmt.Printf("Nums = %v\n", math_struct.Nums)
	calc, err := math_struct.Calculate()
	if err != "" {
		fmt.Println(err)
	} else {
		fmt.Printf("Calculate res for MINUS = %v\n", calc)
	}

	fmt.Printf("Max num = %v\n", math_struct.Max())

	squares, cubes := math_struct.Pow()
	fmt.Printf("squares = %v\n", squares)
	fmt.Printf("cubes = %v\n", cubes)

	fmt.Println("")
	fmt.Println("[== Arrs ==]")

	arr := arr.Arr{Arr: arr.Gen_arr(10)}
	fmt.Printf("Arr = %v\n", arr.Arr)
	fmt.Printf("Len = %v\n", arr.Get_len())
}
