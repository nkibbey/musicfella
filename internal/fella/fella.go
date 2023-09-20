package fella

import (
	"fmt"
	"sort"
	"strings"
	"time"
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
	RDate   time.Time
	Genre   []string
	Country string
}

// String prints fella in form of, **artist(s)** - projectName + optional (feat. artists)
func (f Fella) String() string {
	aa := fmt.Sprintf("%s - %s", PrintBoldArtists(f.Artists), f.Name)
	feat := strings.TrimSpace(PrintArtists(f.FArtists))
	if feat != "" {
		aa = fmt.Sprintf("%s (feat. %s)", aa, feat)
	}
	return aa
}

// FellasOfType provides new slice composed of only the fellas that match Type provided
func FellasOfType(fellas []Fella, t string) (tFellas []Fella) {
	for _, v := range fellas {
		if v.Type == t {
			tFellas = append(tFellas, v)
		}
	}
	return
}

// PrintArtists prints artists in form of a1, a2, a3 or simply a1
// this works for any str slice but I wish to keep the artist formatting specific
// in order to change when desired and not effect others
func PrintArtists(artists []string) string {
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

// PrintBoldArtists prints artists in form of printArtists func except
// wraps string in bold (**) indicators, trimming start/tail whitespaces
func PrintBoldArtists(artists []string) string {
	a := PrintArtists(artists)
	if a != "" {
		bold := "**"
		// string concat $bold + $a was producing MISSING tags when $a had values like %# together
		// have yet to find if %s is completely safe
		a = fmt.Sprintf("%s%s%s", bold, strings.TrimSpace(a), bold)
	}
	return a
}

// PrintGrouping string of Title and each Fella with two new lines of spacing between (none at end)
func PrintGrouping(f []Fella, title string) string {
	var g strings.Builder
	g.WriteString(title)

	for _, v := range f {
		g.WriteString(fmt.Sprintf("\n\n%v", v))
	}

	return g.String()
}

// PrintWeeklyFellas string of albums/eps/songs separated out into sections
func PrintWeeklyFellas(f []Fella) string {
	a, eps, songs := FellasOfType(f, "album"), FellasOfType(f, "ep"), FellasOfType(f, "song")
	sort.Slice(a, func(i, j int) bool {
		return strings.ToLower(PrintArtists(a[i].Artists)) < strings.ToLower(PrintArtists(a[j].Artists))
	})
	sort.Slice(eps, func(i, j int) bool {
		return strings.ToLower(PrintArtists(eps[i].Artists)) < strings.ToLower(PrintArtists(eps[j].Artists))
	})
	sort.Slice(songs, func(i, j int) bool {
		return strings.ToLower(PrintArtists(songs[i].Artists)) < strings.ToLower(PrintArtists(songs[j].Artists))
	})

	var summary strings.Builder
	if len(a) > 0 {
		summary.WriteString(PrintGrouping(a, "## Albums"))
	}
	if len(eps) > 0 {
		if summary.Len() != 0 {
			summary.WriteString("\n\n")
		}
		summary.WriteString(PrintGrouping(eps, "## Eps"))
	}
	if len(songs) > 0 {
		if summary.Len() != 0 {
			summary.WriteString("\n\n")
		}
		summary.WriteString(PrintGrouping(songs, "## Songs"))
	}

	return summary.String()
}
