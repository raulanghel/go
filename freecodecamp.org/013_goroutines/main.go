package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	// go sayHello()
	// time.Sleep(100 * time.Millisecond)

	// bad practice because of race condition
	// var msg = "Hello"
	// go func() {
	// 	fmt.Println(msg)
	// }()
	// msg = "Goodbye"
	// time.Sleep(100 * time.Millisecond)

	// good practice
	// var msg = "Hello"
	// wg.Add(1)
	// go func(msg string) {
	// 	fmt.Println(msg)
	// 	wg.Done()
	// }(msg)
	// msg = "Goodbye"
	// wg.Wait()

	fmt.Printf("Threads: %v\n", runtime.GOMAXPROCS(-1))
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock()
		go sayHello()
		m.Lock()
		go increment()
	}
	wg.Wait()
}

func sayHello() {
	fmt.Printf("Hello #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
