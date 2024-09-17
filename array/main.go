package main

import (
	"fmt"
)

func tryArrayAndSlice() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	a[0], a[1] = a[1], a[0]
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)
	s = append(s, 17, 19, 23)
	fmt.Println(s, len(s), s[len(s)-1])

	names := []string{"John", "Paul", "George", "Ringo"}
	nA := names[0:2]
	nB := names[1:3]
	fmt.Println(nA, nB)
	nB[0] = "XXX"
	fmt.Println(nA, nB)

	ss := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
	}

	fmt.Println(ss)

	emptyS := []int{}
	emptyS = append(emptyS, 1, 2, 3, 4)
	fmt.Println(emptyS)
	emptySCopy := make([]int, len(emptyS))
	copy(emptySCopy, emptyS)
	fmt.Println(emptySCopy)

	for i := range emptyS {
		fmt.Println(emptyS[i])
	}

	/**
	* 初始值：nullS 的初始值为 nil。
	* 内存分配：没有为 nullS 分配任何底层数组的内存。
	* 长度和容量：len(nullS) 和 cap(nullS) 都为 0。
	 */
	var nullS []int
	fmt.Println(nullS, len(nullS), cap(nullS))
	if nullS == nil {
		fmt.Println("nullS is nil!")
	}
	nullS2 := append(nullS, 1)
	fmt.Printf("nullS: %v, is null %v; nullS2: %v \n", nullS, nullS == nil, nullS2)

	sliceP := &nullS
	nullS = append(nullS, 23, 34, 76, 23)
	nullS = append(nullS, 42)
	fmt.Println(nullS, *sliceP)

	arr := [2]int{1, 2}
	slice2 := arr[:]
	slice2P := &slice2
	slice2 = append(slice2, 3)
	slice3 := append(slice2, 3, 4)
	fmt.Println(slice3, slice2, *slice2P, arr)

	/**
	* 初始值：notNullS2 是一个空切片，但不是 nil。
	* 内存分配：已为 notNullS2 分配了一个长度和容量都为 0 的底层数组。
	* 长度和容量：len(notNullS2) 和 cap(notNullS2) 都为 0。
	 */
	notNullS2 := []int{}
	if notNullS2 == nil {
		fmt.Println("notNullS2 is nil!")
	}

	/**
	* 初始值：notNullS3 是一个空切片，但不是 nil。
	* 内存分配：为 notNullS3 分配了一个长度和容量都为 0 的底层数组。
	* 长度和容量：len(notNullS3) 和 cap(notNullS3) 都为 0。
	 */
	notNullS3 := make([]int, 0)
	if notNullS3 == nil {
		fmt.Println("notNullS3 is nil!")
	}

	pow := []int{1, 2, 4, 8, 16, 32, 64, 128}
	for idx, value := range pow {
		fmt.Printf("2**%d = %d\n", idx, value)
	}

	pow2 := make([]int, 11)
	for idx := range pow2 {
		pow2[idx] = 1 << uint(idx)
	}
	fmt.Println(pow2)
}

type PixelCallback func(x, y uint8) uint8

func Pic(dx, dy int, cb PixelCallback) [][]uint8 {
	pic := [][]uint8{}
	for i := 0; i < dy; i++ {
		row := []uint8{}
		for j := 0; j < dx; j++ {
			row = append(row, uint8(cb(uint8(j), uint8(i))))
		}
		pic = append(pic, row)
	}
	return pic
}

func main() {
	tryArrayAndSlice()

	pic := Pic(10, 10, func(x, y uint8) uint8 {
		return x ^ y
	})
	fmt.Println(pic)
}
