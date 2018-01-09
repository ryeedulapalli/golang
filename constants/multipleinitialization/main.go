package main

import (
	"fmt"
)

const (
	//Pi will be exposed as it is capitalized
	Pi = 3.14
	//Language will be exposed as it is capitalized
	Language = "Go"
)

func main() {
	fmt.Println(Pi)
	fmt.Println(Language)
}
