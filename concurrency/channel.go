package main

import (
	"fmt"
	"sync"
	"time"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func TryChannel() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

func TryBufferedChannel() {
	c := make(chan int, 2)
	c <- 1
	c <- 2
	fmt.Println(<-c)
	fmt.Println(<-c)
}

func TryBlockingChannel() {
	startTime := time.Now()
	c := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		fmt.Println("Goroutine 1: Sending data to channel...") // 1
		c <- 42
		fmt.Println("Goroutine 1: Data sent to channel") // 4
		wg.Done()
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Main: Start receiving data from channel...") // 2
	value := <-c
	fmt.Println("Main: Data received:", value) // 3
	wg.Wait()
	fmt.Println(time.Since(startTime))
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TryFibonacci() {
	// c := make(chan int, 10)
	// go fibonacci(cap(c), c)
	c := make(chan int)
	go fibonacci(20, c)
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TrySelect() {
	c := make(chan int)
	quit := make(chan int)

	/**
	主 goroutine 从 channel c 接收 10 个斐波那契数并打印。
	接收完 10 个数后，向 channel quit 发送一个信号，表示完成。
	*/
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// fibonacci2(c, quit)
	// fmt.Println("Main function finished")
	/**
	启动 fibonacci2 函数作为 goroutine：
	fibonacci2 函数在一个新的 goroutine 中运行，生成斐波那契数并发送到 channel c。
	主 goroutine 从 channel c 接收数据：
	*/
	go fibonacci2(c, quit)
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	quit <- 0
	fmt.Println("Main function finished")
	/**
	上面两个实现的区别是：
	1. 执行顺序：
	在第一种写法中，fibonacci2 函数在主 goroutine 中运行，而接收操作在另一个 goroutine 中进行。
	在第二种写法中，fibonacci2 函数在一个新的 goroutine 中运行，而接收操作在主 goroutine 中进行。

	2. 阻塞行为：
	在第一种写法中，主 goroutine 会被 fibonacci2 函数阻塞，直到接收到 quit 信号。
	在第二种写法中，主 goroutine 会被接收操作阻塞，直到接收到 10 个斐波那契数。

	3. 并发性：
	第一种写法中，生成斐波那契数和接收操作是并发进行的。
	第二种写法中，生成斐波那契数和接收操作也是并发进行的，但主 goroutine 负责接收操作。
	*/
}

func TrySelect3() {

	// time.Tick 在 Go 中的行为类似于 JavaScript 中的 setInterval。它会返回一个 channel，每隔指定的时间间隔向该 channel 发送当前时间的值。
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)

	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
