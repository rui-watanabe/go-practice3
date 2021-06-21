package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"sync"
	"time"
)

type A struct{}

type User struct {
	ID      int       `json:"id"`
	Name    string    `json:"name"`
	Email   string    `json:"email"`
	Created time.Time `json:"created"`
	A       *A        `json:"A,omitempty"`
}

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

func jsonFunc() {
	u := new(User)
	u.ID = 1
	u.Name = "Max"
	u.Email = "test@gmail.com"
	u.Created = time.Now()
	fmt.Println(u)

	bs, err := json.Marshal(u) //byte slice
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(bs))

	fmt.Printf("%T\n", bs)

	u2 := new(User)

	err2 := json.Unmarshal(bs, *u2)

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(u2)

}

func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("start")
	time.Sleep(2 * time.Second)
	fmt.Println("end")
	ch <- "process result"
}

func contextFunc() {
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)

	defer cancel()

	go longProcess(ctx, ch)

L:
	for {
		select {
		case <-ctx.Done():
			fmt.Println("err")
			fmt.Println(ctx.Err())
			break L
		case s := <-ch:
			fmt.Println(s)
			fmt.Println("success")
			break L
		}
	}
}

func main() {
	// os.Exit(1)
	// fmt.Println("start")
	syncFunc()
	jsonFunc()
	contextFunc()
}
