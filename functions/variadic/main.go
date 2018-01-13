package main

import (
	"fmt"
)

func main() {
	n := average(10, 20, 30, 40, 50)
	fmt.Println(n)
	fmt.Println("Alternative way of passing arguments")
	data := []float64{45, 64, 32, 75, 29, 90}
	fmt.Println(average(data...))
	fmt.Println("Alternative way of passing slice argument")
	sliceData := []float64{20, 60, 80, 140, 220, 360}
	fmt.Println(averageSlice(sliceData))
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

func averageSlice(sf []float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
