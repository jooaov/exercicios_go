package main

import "fmt"

func main() {
	var x int = 12
	fmt.Printf("%d\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%b\n", x)

	x = 12 >> 1
	fmt.Printf("%d\n", x)
	fmt.Printf("%#x\n", x)
	fmt.Printf("%b\n", x)
}
