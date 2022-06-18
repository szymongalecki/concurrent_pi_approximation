package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var ws sync.WaitGroup
	var add = make(chan float64)
	var sub = make(chan float64)
	var quit = make(chan bool)
	var pairs int = 5
	var pi float64 = 0.0

	for i := 0; i < pairs; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			add <- 4 / ((4 * float64(i)) + 1)
		}(i)

		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sub <- 4 / ((4 * float64(i)) - 1)
		}(i + 1)
	}

	ws.Add(1)
	go func() {
		defer ws.Done()
		for {
			select {
			case a := <-add:
				pi += a
				fmt.Printf("+ %.6f \n", a)
			case s := <-sub:
				pi -= s
				fmt.Printf("- %.6f \n", s)
			case <-quit:
				fmt.Printf("= %.6f \n", pi)
				return
			}
		}
	}()

	wg.Wait()
	quit <- true
	ws.Wait()
}
