package main

import "fmt"

func main() {
	// case 1: float number comparison
	a, b, c := 1.0, 2.0, 3.0
	fmt.Println("1.0 + 2.0 == 3.0 ?", a+b == c)

	// case 2: bit/set operations
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2
	fmt.Printf("x = %08b\n", x)
	fmt.Printf("y = %08b\n", y)
	fmt.Printf("x&y  = %08b\n", x&y)
	fmt.Printf("x|y  = %08b\n", x|y)
	fmt.Printf("x^y  = %08b\n", x^y)
	fmt.Printf("x&^y = %08b\n", x&^y)
}
