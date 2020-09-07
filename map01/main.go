// 1. Use make() to initial a map/channel.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // Declare and initial a string-int map by make()

	input := bufio.NewScanner(os.Stdin) // Declare stdin as the input
	for input.Scan() {                  // Read lines from stdin.
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()

	fmt.Println()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line) // %d -> int, %s -> string
		}
	}
}
