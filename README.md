# 🎨 Colour Namer

This is a Go port of the Node.js [color-namer package](https://github.com/colorjs/color-namer) and was inspired by the
[random color contrasts](https://botsin.space/@randomColorContrasts) bot on Mastodon.

From a given colour input &mdash; either an HTML hex code (`#00ff00`) or an RGB value (`0, 128, 255`) &mdash; it
computes the closest colour values from a variety of colour lists:

- Basic
- HTML - the HTML colour names
- NTC - a [collection](https://chir.ag/projects/ntc/) of named colours
- Pantone
- ROYBGIV - red, orange, yellow, blue, green, indigo, violet

## Usage

Fetch the module:

```
go get github.com/jamesog/colournamer
```

Add it as an import and use one of the `From` functions:

```go
import "github.com/jamesog/colournamer"

colournamer.FromHex("#008080")
colournamer.FromRGB(0, 128, 128)
```

Both functions return a `Results` struct containing each of the lists, with each list containing an array of colours
sorted by how close they are to the given colour value.

## CLI Tool

The `cmd/colournamer` tool lets you use the package as a CLI tool:

```
colournamer #008080 'rgb(0, 128, 128)'
```

The tool will accept either HTML hex values (preceded by `#`) or RGB values in the form `rgb(0, 0, 0)`. Spaces are
optional in the RGB form.

## FAQs

### Why?

I liked the [random color contrasts](https://botsin.space/@randomColorContrasts) bot and had been generating colours
from [Randoma11y](https://randoma11y.com/) and wanted to name them like the bot does.

I didn't want to have to install Node and NPM and have thousands of Node packages cluttering my filesystem.

 ### _Colours?_ It's spelled colors!

I'm English.