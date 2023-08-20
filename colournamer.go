package colournamer

import (
	"errors"
	"image/color"
	"sort"

	"github.com/lucasb-eyer/go-colorful"
)

//go:generate go run ./cmd/colourgen -out colours.go

type Colour struct {
	Name     string
	Colour   colorful.Color
	Distance float64
}

func (c Colour) String() string { return c.Name }

// Hex is the HTML hex representation of the colour.
func (c Colour) Hex() string { return c.Colour.Hex() }

// byDistance implements sort.Interface for []Colour based on Distance.
type byDistance []Colour

func (d byDistance) Len() int           { return len(d) }
func (d byDistance) Swap(i, j int)      { d[i], d[j] = d[j], d[i] }
func (d byDistance) Less(i, j int) bool { return d[i].Distance < d[j].Distance }

// Results contains all colour matches, with each list containing an array of colours sorted by their similarity
// to the given colour.
type Results struct {
	Basic   []Colour
	HTML    []Colour
	NTC     []Colour
	Pantone []Colour
	ROYBGIV []Colour
}

// ClosestName returns the name of the colour whose RGB value is closest to the original value.
func (r Results) ClosestName() string {
	var possibleNames []Colour
	possibleNames = append(possibleNames, r.Basic[0])
	possibleNames = append(possibleNames, r.HTML[0])
	possibleNames = append(possibleNames, r.NTC[0])
	possibleNames = append(possibleNames, r.Pantone[0])
	possibleNames = append(possibleNames, r.ROYBGIV[0])
	sort.Sort(byDistance(possibleNames))
	return possibleNames[0].Name
}

func compute(source colorful.Color, list []Colour) []Colour {
	var r []Colour
	for _, colour := range list {
		colour.Distance = colour.Colour.DistanceRgb(source)
		r = append(r, colour)
	}
	sort.Sort(byDistance(r))
	return r
}

// FromHex takes a hex RGB string and returns the closest colour name.
func FromHex(hex string) (Results, error) {
	c, err := colorful.Hex(hex)
	if err != nil {
		return Results{}, err
	}

	results := Results{
		Basic:   compute(c, Basic),
		HTML:    compute(c, HTML),
		NTC:     compute(c, NTC),
		Pantone: compute(c, Pantone),
		ROYBGIV: compute(c, ROYGBIV),
	}

	return results, nil
}

// FromRGB takes RGB values and returns the closest colour name.
//
// The alpha value is set to 1. Use FromRGBA if you need to control the alpha.
func FromRGB(r, g, b uint8) (Results, error) {
	return FromRGBA(r, g, b, 1)
}

// FromRGBA takes RGBA values and returns the closest colour name.
func FromRGBA(r, g, b, a uint8) (Results, error) {
	c, ok := colorful.MakeColor(color.NRGBA{R: r, G: g, B: b, A: a})
	if !ok || !c.IsValid() {
		return Results{}, errors.New("invalid RGB")
	}

	results := Results{
		Basic:   compute(c, Basic),
		HTML:    compute(c, HTML),
		NTC:     compute(c, NTC),
		Pantone: compute(c, Pantone),
		ROYBGIV: compute(c, ROYGBIV),
	}

	return results, nil
}
