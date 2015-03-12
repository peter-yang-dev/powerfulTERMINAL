package termui

import tm "github.com/nsf/termbox-go"
import rw "github.com/mattn/go-runewidth"

/* ---------------Port from termbox-go --------------------- */

type Attribute uint16

const (
	ColorDefault Attribute = iota
	ColorBlack
	ColorRed
	ColorGreen
	ColorYellow
	ColorBlue
	ColorMagenta
	ColorCyan
	ColorWhite
)

const (
	AttrBold Attribute = 1 << (iota + 9)
	AttrUnderline
	AttrReverse
)

/* ----------------------- End ----------------------------- */

func toTmAttr(x Attribute) tm.Attribute {
	return tm.Attribute(x)
}

func str2runes(s string) []rune {
	return []rune(s)
}

func trimStr2Runes(s string, w int) []rune {
	if w <= 0 {
		return []rune{}
	}
	return []rune(rw.Truncate(s, w, "…"))
}
