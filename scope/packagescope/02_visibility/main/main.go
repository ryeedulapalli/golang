package main

import (
	"fmt"
	"github.com/golang/scope/packagescope/02_visibility/vis"
)

func main() {
	fmt.Println(vis.MyName)
	vis.PrintVar()
}
