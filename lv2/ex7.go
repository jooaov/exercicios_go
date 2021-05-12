package main

import "fmt"

func main() {
	const (
		_ = 1994 + iota
		b
		c
		d
		e
	)

	fmt.Println(b, c, d, e)

}
