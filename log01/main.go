// 1. Use package:log to log messages.
// 2. Package:log uses stderr as default writer.
package main

import (
	"log"
	"os"
)

func main() {
	// case 1: Use stderr as default writer.
	log.Println("Hello world!")

	// case 2: use file as writer.
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("Hello world!")
}
