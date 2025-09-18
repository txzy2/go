package main

import (
	"fmt"
	"time"

	"lab/first/internal/arr"
	"lab/first/internal/mathutil"
)

func main() {
	start := time.Now()

	fmt.Println("")
	fmt.Println("[== Math ==]")

	math_struct := mathutil.NewMath(arr.Gen_arr(10), mathutil.MINUS)
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

	arr := arr.NewArr(arr.Gen_arr(10))
	fmt.Printf("Arr = %v\n", arr.Value)
	fmt.Printf("Len = %v\n", arr.Get_len())

	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}
