package main

import "fmt"

type hotdog int

var b hotdog

func main() {
	x := 10
	fmt.Printf("%d\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%b\n", x)
}
