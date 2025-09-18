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

	nums := arr.Gen_arr(10)
	calculator, err := mathutil.NewCalculator(nums, mathutil.MINUS)
	if err != nil {
		fmt.Printf("Error creating calculator: %v\n", err)
		return
	}
	
	fmt.Printf("Nums = %v\n", calculator.Nums)
	calc, err := calculator.Calculate()
	if err != nil {
		fmt.Printf("Calculate error: %v\n", err)
	} else {
		fmt.Printf("Calculate res for MINUS = %v\n", calc)
	}

	max, err := mathutil.Max(calculator.Nums)
	if err != nil {
		fmt.Printf("Max error: %v\n", err)
	} else {
		fmt.Printf("Max num = %v\n", max)
	}

	squares, cubes, err := mathutil.Pow(calculator.Nums)
	if err != nil {
		fmt.Printf("Pow error: %v\n", err)
	} else {
		fmt.Printf("squares = %v\n", squares)
		fmt.Printf("cubes = %v\n", cubes)
	}

	fmt.Println("")
	fmt.Println("[== Arrs ==]")

	arr := arr.NewArr(arr.Gen_arr(10))
	fmt.Printf("Arr = %v\n", arr.Value)
	fmt.Printf("Len = %v\n", arr.Get_len())

	elapsed := time.Since(start)
	fmt.Printf("Runtime: %s\n", elapsed)
}
