package main

import (
	"fmt"
)

const (
	a = iota //0
	b = iota //1
	c = iota //2
)

const (
	d = iota
	e
	f
)

const (
	_ = iota
	g = iota * 10
	h = iota * 10
)

const (
	_  = iota             // 0
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb = 1 << (iota * 10) // 1 << (2 * 10)
	gb = 1 << (iota * 10) // 1 << (3 * 10)
	tb = 1 << (iota * 10) // 1 << (4 * 10)
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Printf("%b \t %d \n", kb, kb)
	fmt.Printf("%b \t %d \n", mb, mb)
	fmt.Printf("%b \t %d \n", gb, gb)
	fmt.Printf("%b \t %d \n", tb, tb)
}
