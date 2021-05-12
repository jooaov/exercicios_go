package main

import "fmt"

var x hotdog

type hotdog int

func main() {

	fmt.Println("%V", x)
	fmt.Println("%T", x)
	x = 24

}
