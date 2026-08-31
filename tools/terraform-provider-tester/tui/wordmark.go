package tui

import (
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// The "ignition" wordmark: a dot-matrix "GH/TF" that resolves left-to-right
// from a scatter of dithered dots into solid blocks, in the GitHub Primer
// palette. It echoes the dissolve technique seen in modern terminal splashes
// (block letters dithering into a dot field) while using OUR OWN block glyphs
// for the letters - none of another product's letterforms are reproduced.
//
// renderWordmark is a PURE function of (progress, ascii): no clock, no
// randomness, no I/O. The scatter pattern is derived from a deterministic hash
// of each pixel's coordinates, so a given progress always renders identically.
// That keeps the animation golden-testable; the Update loop advances progress
// via tick messages (see introTick* in update wiring), never inside View.

// Pixel glyphs. Every pixel is exactly two cells wide so the block letters read
// as squares on a terminal's ~2:1 tall cells and every row measures equally.
const (
	pixelSolid  = "██" // fully resolved stroke
	ditherHeavy = "▓▓" // mid-resolve
	ditherLight = "░░" // early-resolve
	scatterDot  = "· " // stray dot scattered just ahead of the leading edge
	pixelBlank  = "  " // unlit / not yet reached

	pixelSolidASCII  = "##"
	ditherHeavyASCII = "::"
	ditherLightASCII = ".."
	scatterDotASCII  = ". "
)

const (
	wordmarkRows = 5   // every glyph is 5 cells tall
	glyphCols    = 3   // every glyph is 3 pixels wide
	gapCols      = 1   // one blank pixel column between glyphs
	feather      = 5.0 // width (in pixel columns) of the dither transition band
)

// wordmarkCellWidth is the rendered display width of every row: each glyph is
// three pixels wide with a one-pixel gap, and each pixel occupies two cells.
const wordmarkCellWidth = (len(wordmarkText)*glyphCols + (len(wordmarkText)-1)*gapCols) * 2

// blockFont is our own 5x3 dot-matrix font, limited to the glyphs in "GH/TF".
// '#' marks a lit pixel; ' ' an unlit one.
var blockFont = map[rune][]string{
	'G': {
		"###",
		"#  ",
		"# #",
		"# #",
		"###",
	},
	'H': {
		"# #",
		"# #",
		"###",
		"# #",
		"# #",
	},
	'/': {
		"  #",
		"  #",
		" # ",
		"#  ",
		"#  ",
	},
	'T': {
		"###",
		" # ",
		" # ",
		" # ",
		" # ",
	},
	'F': {
		"###",
		"#  ",
		"## ",
		"#  ",
		"#  ",
	},
}

const wordmarkText = "GH/TF"

// litGrid builds the combined [rows][cols] lit/unlit matrix for the whole word,
// inserting one blank gap column between letters. cols is the pixel width.
func litGrid() ([][]bool, int) {
	letters := []rune(wordmarkText)
	cols := len(letters)*glyphCols + (len(letters)-1)*gapCols
	grid := make([][]bool, wordmarkRows)
	for y := range grid {
		grid[y] = make([]bool, cols)
	}
	x := 0
	for i, r := range letters {
		g := blockFont[r]
		for y := 0; y < wordmarkRows; y++ {
			for c := 0; c < glyphCols; c++ {
				if g[y][c] == '#' {
					grid[y][x+c] = true
				}
			}
		}
		x += glyphCols
		if i < len(letters)-1 {
			x += gapCols // leave the gap column unlit
		}
	}
	return grid, cols
}

// hash01 maps a pixel coordinate to a stable value in [0,1). Integer-only so it
// is identical across platforms and runs (golden-stable).
func hash01(x, y int) float64 {
	h := uint32(x)*73856093 ^ uint32(y)*19349663
	h ^= h >> 13
	h *= 0x5bd1e995
	h ^= h >> 15
	return float64(h&0xffff) / 65536.0
}

// pixelGlyphs returns the (solid, heavy, light, dot) glyph set for the profile.
func pixelGlyphs(ascii bool) (solid, heavy, light, dot string) {
	if ascii {
		return pixelSolidASCII, ditherHeavyASCII, ditherLightASCII, scatterDotASCII
	}
	return pixelSolid, ditherHeavy, ditherLight, scatterDot
}

// styleFor wraps a pixel glyph in its semantic color for the resolve stage and
// letter. Color is stripped under the Ascii profile used by golden tests.
func styleFor(stage int, letter rune) lipgloss.Style {
	switch stage {
	case stageSolid:
		switch letter {
		case 'T', 'F':
			return lipgloss.NewStyle().Foreground(TerraformPurple)
		case '/':
			return lipgloss.NewStyle().Foreground(Muted)
		default:
			return lipgloss.NewStyle().Foreground(Accent)
		}
	case stageHeavy:
		return lipgloss.NewStyle().Foreground(TerraformPurple)
	default: // stageLight / stageDot
		return lipgloss.NewStyle().Foreground(Muted)
	}
}

func wordmarkLetterAtColumn(x int) rune {
	stride := glyphCols + gapCols
	index := x / stride
	if index < 0 || index >= len([]rune(wordmarkText)) || x%stride >= glyphCols {
		return 0
	}
	return []rune(wordmarkText)[index]
}

const (
	stageBlank = iota
	stageDot
	stageLight
	stageHeavy
	stageSolid
)

// pixelStage decides how a lit pixel renders given how far the leading edge has
// swept past it (d = edge - x). Unlit pixels are handled by the caller.
func pixelStage(d, h float64) int {
	switch {
	case d >= feather:
		return stageSolid
	case d <= 0:
		// Ahead of the edge: a sparse scatter of stray dots just before it.
		if d > -2 && h < 0.18 {
			return stageDot
		}
		return stageBlank
	default:
		// Transition band: pixels lock to solid as the edge passes; the rest
		// shimmer as heavy/light dither.
		fill := d / feather
		switch {
		case h < fill:
			return stageSolid
		case h < fill+0.4:
			return stageHeavy
		default:
			return stageLight
		}
	}
}

// renderWordmark renders the "GH/TF" ignition wordmark at the given progress
// in [0,1]. progress 0 is a bare scatter; progress 1 is the fully resolved
// solid wordmark. ascii selects the no-block fallback glyph set.
func renderWordmark(progress float64, ascii bool) string {
	if progress < 0 {
		progress = 0
	}
	if progress > 1 {
		progress = 1
	}
	grid, cols := litGrid()
	solid, heavy, light, dot := pixelGlyphs(ascii)

	// The edge sweeps from column 0 to cols+feather so the rightmost pixels
	// fully resolve at progress 1.
	edge := progress * (float64(cols) + feather)

	var b strings.Builder
	for y := 0; y < wordmarkRows; y++ {
		for x := 0; x < cols; x++ {
			if !grid[y][x] {
				b.WriteString(pixelBlank)
				continue
			}
			d := edge - float64(x)
			stage := pixelStage(d, hash01(x, y))
			var glyph string
			switch stage {
			case stageSolid:
				glyph = solid
			case stageHeavy:
				glyph = heavy
			case stageLight:
				glyph = light
			case stageDot:
				glyph = dot
			default:
				b.WriteString(pixelBlank)
				continue
			}
			if ascii {
				b.WriteString(glyph)
			} else {
				b.WriteString(styleFor(stage, wordmarkLetterAtColumn(x)).Render(glyph))
			}
		}
		if y < wordmarkRows-1 {
			b.WriteByte('\n')
		}
	}
	return b.String()
}
