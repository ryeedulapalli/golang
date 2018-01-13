package main

import (
	"fmt"
)

func filter(numbers []int, callback func(int)) {
	for _, v := range numbers {
		callback(v)
	}
}

func filterCallBack(cbNumbers []int, callback2 func(int) bool) []int {
	var xs []int
	for _, n := range cbNumbers {
		if callback2(n) {
			xs = append(xs, n)
		}
	}
	return xs
}

func main() {
	filter([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})
	fmt.Println("Call back example sepration")
	xs := filterCallBack([]int{1, 2, 3, 4}, func(n int) bool {
		return n > 1
	})

	fmt.Println(xs)
}
