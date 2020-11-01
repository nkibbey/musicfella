package main

import (
	"fmt"
	"sort"
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
	// Artists featured on this music fella
	FArtists []string
	// RDate defines release date, should only be used as a date and not exact time
	RDate time.Time
	// Links define the links to find the music
	Links   []string
	Genre   []string
	Country string
}

// String prints fella in form of, **artist(s)** - projectName + optional (feat. artists)
func (f Fella) String() string {
	aa := fmt.Sprintf("%s - %s", printBoldArtists(f.Artists), f.Name)
	feat := strings.TrimSpace(printArtists(f.FArtists))
	if feat != "" {
		aa = fmt.Sprintf("%s (feat. %s)", aa, feat)
	}
	return aa
}

// fellasOfType provides new slice composed of only the fellas that match Type provided
func fellasOfType(fellas []Fella, t string) (tFellas []Fella) {
	for _, v := range fellas {
		if v.Type == t {
			tFellas = append(tFellas, v)
		}
	}
	return
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

// printGrouping string of Title and each Fella with two new lines of spacing between (none at end)
func printGrouping(f []Fella, title string) string {
	var g strings.Builder
	g.WriteString(title)

	for _, v := range f {
		g.WriteString(fmt.Sprintf("\n\n%v", v))
	}

	return g.String()
}

// printWeeklyFellas string of albums/eps/songs separated out into sections
func printWeeklyFellas(f []Fella) string {
	a, eps, songs := fellasOfType(f, "album"), fellasOfType(f, "ep"), fellasOfType(f, "song")
	sort.Slice(a, func(i, j int) bool {
		return strings.ToLower(printArtists(a[i].Artists)) < strings.ToLower(printArtists(a[j].Artists))
	})
	sort.Slice(eps, func(i, j int) bool {
		return strings.ToLower(printArtists(eps[i].Artists)) < strings.ToLower(printArtists(eps[j].Artists))
	})
	sort.Slice(songs, func(i, j int) bool {
		return strings.ToLower(printArtists(songs[i].Artists)) < strings.ToLower(printArtists(songs[j].Artists))
	})

	var summary strings.Builder
	if len(a) > 0 {
		summary.WriteString(printGrouping(a, "## Albums"))
	}
	if len(eps) > 0 {
		if summary.Len() != 0 {
			summary.WriteString("\n\n")
		}
		summary.WriteString(printGrouping(eps, "## Eps"))
	}
	if len(songs) > 0 {
		if summary.Len() != 0 {
			summary.WriteString("\n\n")
		}
		summary.WriteString(printGrouping(songs, "## Songs"))
	}

	return summary.String()
}

func main() {
	fmt.Printf("Hello \n")
	x := read.Reader{}
	x.Read()

	format := "2020-01-01"

	vbbbDate, _ := time.Parse(format, "2018-12-12")
	vbbb := Fella{
		Artists:  []string{"clipping."},
		Name:     "Visions of Bodies Being Burned",
		Type:     "album",
		RDate:    vbbbDate,
		FArtists: []string{"Sickness", "Michael Esposito", "Cam & China", "Greg Stuart..."},
		Country:  "usa",
		Genre:    []string{"rap", "alternative", "harsh noise"},
		Links:    []string{"https://clppng.bandcamp.com/album/visions-of-bodies-being-burned"},
	}
	fmt.Println(vbbb)

	fmt.Println(printWeeklyFellas([]Fella{
		vbbb,
		vbbb,
		Fella{
			Artists: []string{"loona"},
			Name:    "[12:00]",
			Type:    "album",
		},
	}))
}
