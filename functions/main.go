package main

import (
	"fmt"
	"math"
)

func comput(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	fmt.Println(comput(hypot))
	fmt.Println(comput(math.Pow))

	// pos, neg := adder(), adder()
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(
	// 		pos(i),
	// 		neg(-2*i),
	// 	)
	// }
	fmt.Println("Fibonacci:")
	fib := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(fib())
	}
}
