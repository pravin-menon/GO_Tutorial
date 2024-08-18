package main

import (
	"fmt"
	"math"
	"runtime"
)

// func main(){
// 	for i := 0; i < 50; i++ {
// 		go printNumber(i)
// 	}
// }

// func printNumber(num int) {
// 	fmt.Println(num)
// }

func Sqrt(x float64) float64 {
	z := float64(x/2)
	y := float64(0)
	// i := float64(1)
	for {
		z -= (z*z - x) / (2*z)
		// fmt.Printf("Iteration %g, value %g\n", i, z)
		// fmt.Println(y)
		if y - z == 0 {
			break
		}
		// i +=1
		y = z
	}
	return y
}

func main() {
	fmt.Printf("Output from custom Sqrt function is %g\n", Sqrt(10984300008909))
	fmt.Printf("Output from custom Math function is %g\n", math.Sqrt(10984300008909))
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Println("Windows.")
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
	// Sqrt(2)
}