package calc

import (
	"fmt"
	"math"
)

// task 1

func Sum(nums ...float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func Max(nums ...float64) float64 {
	maxNum := math.Inf(-1)
	for _, num := range nums {
		maxNum = math.Max(maxNum, num)
	}
	return maxNum
}

func Min(nums ...float64) float64 {
	minNum := math.Inf(1)
	for _, num := range nums {
		minNum = math.Min(minNum, num)
	}
	return minNum
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("can not divide by zero")
	}
	return a / b, nil
}

func init() {
	fmt.Printf("init: %v\n", Sum(5, 2))
}

// task 2

type Calculator interface {
	Sum(nums ...float64) float64
	Max(nums ...float64) float64
	Min(nums ...float64) float64
	Divide(a, b float64) (float64, error)
}

type Calc struct{}

func (c Calc) Sum(nums ...float64) float64 {
	return Sum(nums...)
}
func (c Calc) Max(nums ...float64) float64 {
	return Max(nums...)
}
func (c Calc) Min(nums ...float64) float64 {
	return Min(nums...)
}
func (c Calc) Divide(a, b float64) (float64, error) {
	return Divide(a, b)
}
