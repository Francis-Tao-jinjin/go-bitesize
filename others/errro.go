package main

import (
	"fmt"
	"strconv"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

// Error() 方法是 error 接口的方法, 只要实现了 Error() 方法的类型都可以作为 error 类型使用
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func TryError() {
	if err := run(); err != nil {
		fmt.Println(err)
	}

	i, err := strconv.Atoi("42.34")
	if err != nil {
		fmt.Printf("couldn't convert number: %v\n", err)
	} else {
		fmt.Println(i)
	}

	fmt.Println(Sqrt(2.2))
	fmt.Println(Sqrt(-2.2))
}
