package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func withMutex() int {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go func() {
		defer close(evenCh)
		defer close(oddCh)
		for i := 1; i <= 1000; i++ {
			if i%2 == 0 {
				evenCh <- i
			} else {
				oddCh <- i
			}
		}
	}()

	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			case v, ok := <-evenCh:
				if !ok {
					break loop
				}
				if v%3 == 0 {
					mu.Lock()
					counter++
					mu.Unlock()
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			case v, ok := <-oddCh:
				if !ok {
					break loop
				}
				if v%33 == 0 {
					mu.Lock()
					counter--
					mu.Unlock()
				}
			}
		}
	}()

	wg.Wait()
	return counter
}

func withAtomic() int64 {
	evenCh := make(chan int)
	oddCh := make(chan int)

	go func() {
		defer close(evenCh)
		defer close(oddCh)
		for i := 1; i <= 1000; i++ {
			if i%2 == 0 {
				evenCh <- i
			} else {
				oddCh <- i
			}
		}
	}()

	var counter int64
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			case v, ok := <-evenCh:
				if !ok {
					break loop
				}
				if v%3 == 0 {
					atomic.AddInt64(&counter, 1)
				}
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
	loop:
		for {
			select {
			case v, ok := <-oddCh:
				if !ok {
					break loop
				}
				if v%33 == 0 {
					atomic.AddInt64(&counter, -1)
				}
			}
		}
	}()

	wg.Wait()
	return counter
}

func main() {
	fmt.Println("With mutex:", withMutex())
	fmt.Println("With atomic:", withAtomic())
}
