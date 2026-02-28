package main

import (
	"fmt"

	"github.com/fominvic81/lntu-go/lab_03/calc"
)

func calculate(calculator calc.Calculator, a, b float64) {
	fmt.Printf("sum: %v\n", calculator.Sum(a, b))
	fmt.Printf("min: %v\n", calculator.Min(a, b))
	fmt.Printf("max: %v\n", calculator.Max(a, b))

	result, err := calculator.Divide(a, b)
	if err != nil {
		fmt.Printf("failed to divide: %v\n", err.Error())
	} else {
		fmt.Printf("div: %v\n", result)
	}
}

func main() {
	// task 1
	fmt.Printf("sum: %v\n", calc.Sum(12, 54))
	fmt.Printf("min: %v\n", calc.Min(34, 66))
	fmt.Printf("max: %v\n", calc.Max(42, 44))

	result, err := calc.Divide(43, 97)
	if err != nil {
		fmt.Printf("failed to divide: %v\n", err.Error())
	} else {
		fmt.Printf("div: %v\n", result)
	}

	// task 2
	calculator := calc.Calc{}
	calculate(calculator, 42, 13)
	calculate(calculator, 42, 0)
}
