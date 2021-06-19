package main

import (
	"fmt"
	"sync"
)

func syncFunc() {
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("first goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("second goroutine")
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 5; i++ {
			fmt.Println("third goroutine")
		}
		wg.Done()
	}()

	wg.Wait()
}

func main() {
	// os.Exit(1)
	// fmt.Println("start")
	syncFunc()
}
