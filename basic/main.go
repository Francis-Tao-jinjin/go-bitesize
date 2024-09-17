package main

import (
	"fmt"
	"math/cmplx"
	"runtime"
	"strconv"
	"time"
	"unsafe"
)

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func Sqrt(x float64) float64 {
	z := 1.0
	for i := 0; i < 10; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z
}

func trySwitch() {
	defer fmt.Println("This should be printed last")
	defer fmt.Println("This should be printed right before the last")

	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", os)
	}

	fmt.Printf("When's Saturday?\n")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
	fmt.Printf("today %T, %v\n", today, today)
}

/**
 * 1. i 被声明并初始化为 0。
 * 2. defer 语句被注册，但不会立即执行。
 * 3. return 1 设置 i 为 1。
 * 4. defer 语句执行，将 i 自增为 2。
 * 5. 函数返回 i 的最终值 2。
**/
func tryDefer() (i int) {
	defer func() { i++ }()
	return 1
}

func tryPointer() {
	i, j := 42, 2701
	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer

	p = &j // point to j
	*p = *p / 37
	fmt.Println(j)

	fmt.Printf("Type of p: %T, Value is %v, pointed to value: %v\n", p, p, *p) // Type of p: *int, Value is 0xc0000140b0

	f := Sqrt(5)
	p = (*int)(unsafe.Pointer(&f))
	fmt.Printf("Type of p: %T, Value is %v, pointed to value: %v\n", p, p, *p) // Type of p: *int, Value is 0xc0000140b0

	type Vertex struct {
		X int
		Y int
	}
	v := Vertex{1, 2}
	sp := &v
	sp.X = 1e3
	(*sp).Y = 3e2
	fmt.Println(v)
}

func tryToString() {
	num := 123
	numStr := fmt.Sprintf("%d", num)
	fmt.Printf("Type of numStr: %T, Value is %v\n", numStr, numStr)
	numStr = fmt.Sprintf("%t", true)
	fmt.Printf("Type of numStr: %T, Value is %v\n", numStr, numStr)

	numStr = strconv.Itoa(num)
	boolStr := strconv.FormatBool(true)
	fmt.Printf("Type of numStr: %T, Value is %v\n", boolStr, boolStr)

	floatNum := 3.14159
	floatStr := strconv.FormatFloat(floatNum, 'f', 6, 64) // 'f'表示格式，6表示小数点后保留6位，64表示float64
	fmt.Println("Float as string using strconv:", floatStr)
	fmt.Println("Float as string using strconv:", strconv.FormatFloat(floatNum, 'E', -1, 32))
	fmt.Println("Float as string using strconv:", strconv.FormatFloat(floatNum, 'E', -1, 64))
	fmt.Println("Float as string using strconv:", strconv.FormatFloat(floatNum, 'g', -1, 64))
}

func main() {
	fmt.Printf("Type: %T , Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T , Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T , Value: %v\n", z, z)

	var i = 42
	f := float32(i) + 0.2
	u := uint(f)
	fmt.Println(i, f, u)

	// trySwitch()

	// fmt.Println(tryDefer())

	tryPointer()
	// tryToString()
}
