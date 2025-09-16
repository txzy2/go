package main

import (
	"fmt"

	"lab/first/internal/arr"
	"lab/first/internal/mathutil"
)

func main() {
	math := mathutil.Math{A: 2, B: 4}
	fmt.Printf("Sum = %v\n", math.Sum())

	arr := arr.Arr{Arr: arr.Gen_arr(10)}
	fmt.Printf("Arr = %v\n", arr.Arr)
	fmt.Printf("Len = %v\n", arr.Get_len())
}
