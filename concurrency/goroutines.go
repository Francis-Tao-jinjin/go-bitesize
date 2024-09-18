package main

import (
	"fmt"
	"sync"
	"time"
)

func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Printf("%v: %q\n", i, s)
	}
}

func TryGoroutine() {
	wg := sync.WaitGroup{}
	wg.Add(1) // wg.Add(1) 需要在启动 goroutine 之前调用
	go say("world", &wg)
	// say("hello")
	func() {
		wg.Wait()
		fmt.Println("WaitGroup finished")
	}()
	fmt.Println("Main function finished")
}
