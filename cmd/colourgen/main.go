package main

import (
	"bytes"
	"flag"
	"fmt"
	"github.com/lucasb-eyer/go-colorful"
	"go/format"
	"io"
	"os"
)

func main() {
	outFile := flag.String("out", "", "Output file (default: stdout)")
	flag.Parse()
	out := os.Stdout
	if *outFile != "" {
		f, err := os.Create(*outFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Couldn't open %s: %v\n", *outFile, err)
			os.Exit(1)
		}
		out = f
	}

	var buf bytes.Buffer
	buf.WriteString("package colournamer\n\n")
	buf.WriteString("import \"github.com/lucasb-eyer/go-colorful\"\n\n")
	buf.WriteString("var (\n")
	generateColors(&buf, "Basic", basic)
	generateColors(&buf, "HTML", html)
	generateColors(&buf, "NTC", ntc)
	generateColors(&buf, "Pantone", pantone)
	generateColors(&buf, "ROYGBIV", roygbiv)
	buf.WriteString(")\n")
	src, err := format.Source(buf.Bytes())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error formatting generated code: %v\n", err)
		os.Exit(1)
	}
	out.Write(src)
}

func generateColors(buf io.Writer, name string, colours []colour) {
	fmt.Fprintf(buf, "\t%s = []Colour{\n", name)
	for _, c := range colours {
		cc, _ := colorful.Hex(c.hex)
		fmt.Fprintf(buf, "\t\t{Name: %q, Colour: %#v},\n", c.name, cc)
	}
	fmt.Fprintln(buf, "\t}")
}
