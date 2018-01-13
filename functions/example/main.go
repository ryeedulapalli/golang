package main

import (
	"fmt"
)

func main() {
	fmt.Println(half(1))
	fmt.Println(half(2))
	halfFunc := func(hf int) (int, bool) {
		return hf / 2, hf%2 == 0
	}
	fmt.Println("Using anonymous function expression")
	fmt.Println(halfFunc(1))
	fmt.Println(halfFunc(2))
	fmt.Println("Variadic function")
	aSlice := []int{2, 5, 3, 10}
	fmt.Println(variadicGreatest(aSlice))
	foo(aSlice...)
	fmt.Println(max(250, 150, 50, 350, 125))
	fmt.Println("Expression value")
	evalExp := (true && false) || (false && true) || !(false && false)
	fmt.Println(evalExp)
}

func half(n int) (int, bool) {
	return n / 2, n%2 == 0
}

func variadicGreatest(sf []int) int {
	fmt.Printf("%T \t", sf)
	var largest int
	for _, v := range sf {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func max(numbers ...int) int {
	fmt.Printf("%T \t", numbers)
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}

func foo(numbers ...int) {
	fmt.Println(numbers)
}
