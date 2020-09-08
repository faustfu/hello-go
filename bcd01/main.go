package main

import (
	"fmt"

	"github.com/yerden/go-util/bcd"
)

func main() {
	a := []byte{0x12}

	dec := bcd.NewDecoder(bcd.Standard)
	dst := make([]byte, bcd.DecodedLen(len(a)))
	n, err := dec.Decode(dst, a)
	if err != nil {
		return
	}
	fmt.Println(string(dst[:n]))
}
