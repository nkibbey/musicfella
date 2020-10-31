package main

import (
	"fmt"
	"strings"
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

// printArtists prints artists in form of a1, a2, a3 or simply a1
// this works for any str slice but I wish to keep the artist formatting specific
// in order to change when desired and not effect others
func printArtists(artists []string) string {
	np := ""
	if len(artists) > 0 {
		// first element is being assigned directly while others are sprinted w/format operator
		np = artists[0]
	}
	for i, a := range artists {
		if i > 0 {
			np += fmt.Sprintf(", %s", a)
		}
	}

	return np
}

// printBoldArtists prints artists in form of printArtists func except
// wraps string in bold (**) indicators, trimming start/tail whitespaces
func printBoldArtists(artists []string) string {
	a := printArtists(artists)
	if a != "" {
		bold := "**"
		// string concat $bold + $a was producing MISSING tags when $a had values like %# together
		// have yet to find if %s is completely safe
		a = fmt.Sprintf("%s%s%s", bold, strings.TrimSpace(a), bold)
	}
	return a
}

func (f Fella) String() string {
	if f.Type == "album" {
		return fmt.Sprintf("%v", printBoldArtists(f.Artists))
	}
	return fmt.Sprintf("Name:")
}

func main() {
	fmt.Printf("Hello \n")
	x := read.Reader{}
	a := [...]string{"hi"}
	y := fmt.Sprintf("%v", a)
	fmt.Printf("%v\n", y)
	x.Read()
}
