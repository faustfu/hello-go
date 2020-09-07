// 1. Get arguments from os.Args.
// 2. for-loop is the only one loop way in golang.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string // declare two string variables in function scope.

	for i := 1; i < len(os.Args); i++ { // declare a integer variable:i in this "for" block scope.
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
