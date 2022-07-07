package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rating := 0
	cnt := 0
	var wg sync.WaitGroup

	for i := 0; i < 200; i++ {
		wg.Add(1)
		i := i
		go func() {
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(5)))
			fmt.Println(i)
			rating = (rating*cnt + rand.Intn(100)) / (cnt + 1)
			cnt++
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Print("avg rating: ")
	fmt.Println(rating)
	fmt.Print("total student: ")
	fmt.Println(cnt)
}
