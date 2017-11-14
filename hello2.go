package main

import "fmt"

func main() {
	fmt.Println("This is just a test of this")
	z := 5
	zp := junk(z)
	fmt.Println(zp)
}

func junk(x int) int {
	y := x + 1
	return y
}
