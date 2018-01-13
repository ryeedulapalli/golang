package main

import (
	"fmt"
)

func main() {
	greeting()
	greetingWithName("Dad")
	greetingStranger("Jane", "Doe")
	fmt.Println(greetingWithReturn("Andy ", "America"))
	fmt.Println(greetingWithMultipleReturns("Joe ", "Biden "))
}

func greeting() {
	fmt.Println("Hello")
}

func greetingWithName(name string) {
	fmt.Println("Hello ", name)
}

func greetingStranger(first, last string) {
	fmt.Println("Hello ", first, last)
}

func greetingWithReturn(first, last string) string {
	return fmt.Sprint("Hello ", first, last)
}

func greetingWithMultipleReturns(fname, lname string) (string, string) {
	return fmt.Sprint("Hello ", fname, lname), fmt.Sprint("Hello ", lname, fname)
}
