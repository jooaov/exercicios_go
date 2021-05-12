package main

import "fmt"

var x int = 42
var y string = "james"
var z bool = true

func main() {
	s := fmt.Sprintf("%v\t%v\t%v", x, y, z)
	// fmt.Printf("%T\n", x)
	// fmt.Printf("%T\n", y)
	// fmt.Printf("%T\n", z)
	fmt.Printf("%v\n", s)
}
