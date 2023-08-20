package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/jamesog/colournamer"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s <colour>...\n", os.Args[0])
		fmt.Fprintln(os.Stderr, `
<colour> can be an "HTML" hex value #0f0f0f or an RGB value rgb(0, 255, 128).

Multiple colour codes can be provided and you may mix hex and RGB.`)
		os.Exit(1)
	}

Loop:
	for _, colour := range os.Args[1:] {
		var r colournamer.Results
		var err error
		switch {
		case strings.HasPrefix(colour, "#"):
			r, err = hex(colour)
		case strings.HasPrefix(colour, "rgb("):
			r, err = rgb(colour)
		default:
			fmt.Fprintf(os.Stderr, "Unknown format for %q\n", colour)
			break Loop
		}
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			continue
		}
		name := r.ClosestName()
		t := cases.Title(language.English).String(name)
		fmt.Println(colour, t)
	}
}

func hex(c string) (colournamer.Results, error) {
	r, err := colournamer.FromHex(c)
	if err != nil {
		return colournamer.Results{}, fmt.Errorf("%s is not a valid hex value", c)
	}
	return r, nil
}

func rgb(c string) (colournamer.Results, error) {
	var r, g, b uint8
	c = strings.ReplaceAll(c, " ", "")
	fmt.Sscanf(c, "rgb(%d,%d,%d)", &r, &g, &b)
	res, err := colournamer.FromRGB(r, g, b)
	if err != nil {
		return colournamer.Results{}, fmt.Errorf("%s is not a valid RGB value: %v", c)
	}
	return res, nil
}
