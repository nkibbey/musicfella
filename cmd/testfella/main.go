package main

import (
	"fmt"
	"time"

	"github.com/nkibbey/musicfella/internal/fella"
	"github.com/nkibbey/musicfella/internal/read"
)

func main() {
	fmt.Printf("Hello \n")
	x := read.Reader{}
	x.Read()

	format := "2006-01-02"

	vbbbDate, _ := time.Parse(format, "2018-12-12")
	vbbb := fella.Fella{
		Artists:  []string{"clipping."},
		Name:     "Visions of Bodies Being Burned",
		Type:     "album",
		RDate:    vbbbDate,
		FArtists: []string{"Sickness", "Michael Esposito", "Cam & China", "Greg Stuart..."},
		Country:  "usa",
		Genre:    []string{"rap", "alternative", "harsh noise"},
	}
	fmt.Println(vbbb)

	fmt.Println(fella.PrintWeeklyFellas([]fella.Fella{
		vbbb,
		vbbb,
		{
			Artists: []string{"loona"},
			Name:    "[12:00]",
			Type:    "album",
		},
	}))
}
