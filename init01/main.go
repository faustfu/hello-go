// 1. Use multiple init() to initialize the package.
// 2. It is possible having multiple function:init. They will run one by one as loading the package.
package main

import "fmt"

func init() {
	fmt.Println("init 一")
}

func init() {
	fmt.Println("init 二")
}

func main() {
	fmt.Println("init 01...")
}
