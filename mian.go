package main

import (
	"fmt"
	"sync"
)

func main() {
	bowl := 10
	defer fmt.Println(bowl)
	var wg sync.WaitGroup
	defer wg.Wait()
	for i, v := range []int{1, 2, 3, 4, 5, 6} {
		wg.Add(1)
		go func(i, v int) {
			defer wg.Done()
			bowl *= v
			fmt.Println(i, v, bowl)
		}(i, v)
	}
}
