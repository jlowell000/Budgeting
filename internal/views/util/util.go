package util

import (
	"fmt"

	"github.com/muesli/termenv"
)

var (
	Term    = termenv.EnvColorProfile()
	Keyword = makeFgStyle("211")
	Subtle  = makeFgStyle("241")
	Dot     = ColorFg(" • ", "236")
)

func Instructions() string {
	return Subtle("j/k, up/down: select") + Dot +
		Subtle("enter: choose") + Dot +
		Subtle("b, backspace: previous screen") + Dot +
		Subtle("q, esc: quit")
}

func Checkbox(label string, checked bool) string {
	if checked {
		return ColorFg("[x] "+label, "212")
	}
	return fmt.Sprintf("[ ] %s", label)
}

func ColorFg(val, color string) string {
	return termenv.String(val).Foreground(Term.Color(color)).String()
}

func makeFgStyle(color string) func(string) string {
	return termenv.Style{}.Foreground(Term.Color(color)).Styled
}
