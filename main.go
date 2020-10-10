package main

import (
	"fmt"
	"time"

	"github.com/nkibbey/musicfella/read"
)

// Album defines info to aid musicfella users to identify a real life album
type Album struct {
	Name    string
	Artist  string
	Rdate   time.Time
	Links   []string
	Genre   []string
	Country string
}

func main() {
	fmt.Printf("Hello \n")
	x := read.Reader{}

	x.Read()
}
