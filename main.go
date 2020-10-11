package main

import (
	"fmt"
	"time"

	"github.com/nkibbey/musicfella/read"
)

// Fella defines info to aid musicfella users to identify a real life album/ep/etc
type Fella struct {
	Name string
	// Artists created this music fella, primary artist should be the first element followed by the rest
	Artists []string
	// Type defines whether this music fella is an album, ep, or some other type
	Type string
	// Rdate defines release date, should only be used as a date and not exact time
	Rdate time.Time
	// Links define the links to find the music
	Links   []string
	Genre   []string
	Country string
}

func main() {
	fmt.Printf("Hello \n")
	x := read.Reader{}

	x.Read()
}
