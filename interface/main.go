package main

import (
	"fmt"
	"math"
	"reflect"
)

type Abser interface {
	Abs() float64
	Abs2() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func (f MyFloat) Abs2() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

// 参数的类型是否是指针，会直接影响到 interface 变量的使用
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Abs2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func tryVertex() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	// 如果希望这个赋值能够通过编译，那么 Vertex 的 Abs 方法必须是是普通类型，而不是指针类型
	// a = v
	fmt.Println(a.Abs())
	fmt.Println(a.Abs2())
}

type I interface {
	M()
}

type T struct {
	S string
}

/*
*
在 Go 语言中，方法的接收者类型决定了该方法是属于值类型还是指针类型。具体来说：

如果方法的接收者是值类型（例如 func (t T) M()），那么该方法可以被值类型的变量调用。
如果方法的接收者是指针类型（例如 func (t *T) M()），那么该方法只能被指针类型的变量调用。

所以，我们可以说： *T 类型实现了接口 I 的方法 M，而 T 类型没有实现接口 I 的方法 M。
*/
func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

// The interface type that specifies zero methods is known as the empty interface:
// An empty interface may hold values of any type.
func describe(any interface{}) {
	fmt.Printf("(%v, %T)\n", any, any)
}

func tryInterface() {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()
	i = &T{"Hello"}
	describe(i)
	i.M()

	t2 := T{"To be or not to be"}
	var i2 I
	i2 = &t2
	i2.M()
	describe(i2)
}

func tryTypeAssertion() {
	var i interface{} = "hello"
	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	// i holds a float64, then t will be the underlying float64 value, and ok will be true.
	// If i does not hold a float64, then t will be the ZERO value of type float64, and ok will be false.
	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}

type Number interface {
	int | int64 | float64 | float32
}

// 调用的时候必须要指定泛型的类型
func sum[T Number](nums interface{}) T {
	var total T
	v := reflect.ValueOf(nums)
	fmt.Printf("nums is %v, kind: %v, type: %v\n", v, v.Kind(), v.Type())
	fmt.Printf("nums[0] is %v\n", v.Index(0))
	for i := 0; i < v.Len(); i++ {
		total += v.Index(i).Interface().(T)
	}
	return total
}

// 这种泛型的写法只能支持 Slice 类型，不能支持 Array 类型
// func sum[T Number](nums []T) T {
// 	var total T
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

// 这种写法只能支持 int 类型的 Slice
// func sum(nums []int) int {
// 	total := 0
// 	for _, num := range nums {
// 		total += num
// 	}
// 	return total
// }

func trySwitches(any interface{}) {
	switch v := any.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		// %q is a quoted string, %v is the value in a default format (without quotes)
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	case []int:
		// fmt.Printf("The sum of %v is %v\n", v, sum(v))
		fmt.Printf("The sum of %v is %v\n", v, sum[int](v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func trySum() {
	nums := []int{1, 2, 3, 4, 5}
	fmt.Println(sum[int](nums))
	nums2 := []float64{1.1, 2.2, 3.3, 4.4, 5.5}
	fmt.Println(sum[float64](nums2))
}

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	// it is 'Sprintf', not 'Printf'
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (ip IPAddr) String() string {
	return fmt.Sprintf("%v.%v.%v.%v", ip[0], ip[1], ip[2], ip[3])
}

func tryStriner() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)

	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

func main() {
	// tryInterface()
	// tryTypeAssertion()
	trySwitches(21)
	trySwitches("maybe")
	trySwitches(true)
	trySwitches([]int{1, 3, 5, 2})
	trySum()
	tryStriner()
}
