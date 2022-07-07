package main

import (
	"encoding/json"
	"fmt"
	"sync"
)

var mu sync.Mutex

func main() {
	arr := make([]string, 5)
	arr = []string{"quick", "brown", "fox", "dog", "quick"}
	mp := make(map[string]int)

	c := make(chan int, 5)
	for i, s := range arr {
		i := i
		s := s
		go func() {
			mu.Lock()
			defer mu.Unlock()
			for _, val := range s {
				if mp[string(val)] == 0 {
					mp[string(val)] = 1
				} else {
					mp[string(val)] = mp[string(val)] + 1
				}
			}
			c <- i
		}()
	}
	for i := 0; i < len(arr); i++ {
		fmt.Println(<-c)
	}

	p, _ := json.MarshalIndent(mp, " ", " ")
	fmt.Printf(string(p))

}
