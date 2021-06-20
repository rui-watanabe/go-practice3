package main

import (
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

func main() {
	// os.Exit(1)
	// fmt.Println("start")
	syncFunc()
	jsonFunc()
}
