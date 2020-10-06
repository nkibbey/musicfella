package main

import (
	"fmt"

	"github.com/nkibbey/musicfella/read"
)

func main() {
	fmt.Printf("Hello \n")
	x := read.Reader{}

	x.Read()
}
