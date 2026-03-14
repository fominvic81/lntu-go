package main

import "fmt"

func generate() <-chan int {
	out := make(chan int, 10)
	go func() {
		defer close(out)
		for i := range 100 {
			out <- i
		}
	}()
	return out
}

func filterEven(in <-chan int) <-chan int {
	out := make(chan int, 10)
	go func() {
		defer close(out)
		for i := range in {
			if i%2 == 0 {
				out <- i
			}
		}
	}()
	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int, 10)
	go func() {
		defer close(out)
		for i := range in {
			out <- i * i
		}
	}()
	return out
}

func sum(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		sum := int(0)
		for i := range in {
			sum += i
		}
		out <- sum
	}()
	return out
}

func main() {
	a := generate()
	b := filterEven(a)
	c := square(b)
	d := sum(c)

	sum := <-d
	fmt.Printf("Sum: %v\n", sum)
}
