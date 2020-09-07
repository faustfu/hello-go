package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s, sep := "", ""

	// case 1
	for _, arg := range os.Args[1:] { // 1. Declare a index-value pair to keep range output. 2. Use underscore to ignore unused variables.
		s += sep + arg
		sep = " "
	}

	fmt.Println(s)

	// case 2: Use strings.Join() to concate string array.
	fmt.Println(strings.Join(os.Args[1:], " "))
}
