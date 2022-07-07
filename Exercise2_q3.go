package main

import (
	"fmt"
	"sync"
)

var mu sync.Mutex

func deposit(b *int, a int, c chan int) {
	mu.Lock()
	defer mu.Unlock()
	*b = *b + a
	c <- 1
}
func withdrawal(b *int, a int, c chan int) {
	//ime.Sleep(time.Millisecond * 5)
	mu.Lock()
	defer mu.Unlock()
	if *b < a {
		fmt.Println("not enough money")
	} else {
		*b = *b - a
	}
	c <- 1
}

func main() {
	bal := 500
	c := make(chan int, 2)
	go deposit(&bal, 200, c)
	go withdrawal(&bal, 600, c)
	<-c
	<-c
	fmt.Println(bal)

}
