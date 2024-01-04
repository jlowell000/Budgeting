package util

import (
	"fmt"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/shopspring/decimal"
)

var (
	Term    = termenv.EnvColorProfile()
	Keyword = makeFgStyle("211")
	Subtle  = makeFgStyle("241")
	Dot     = ColorFg(" • ", "236")

	FocusedStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("205"))
	BlurredStyle        = lipgloss.NewStyle().Foreground(lipgloss.Color("240"))
	CursorStyle         = FocusedStyle.Copy()
	NoStyle             = lipgloss.NewStyle()
	HelpStyle           = BlurredStyle.Copy()
	CursorModeHelpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("244"))

	FocusedButton = FocusedStyle.Copy().Render("[ Submit ]")
	BlurredButton = fmt.Sprintf("[ %s ]", BlurredStyle.Render("Submit"))
)

func Instructions() string {
	return Subtle("up/down: select") + Dot +
		Subtle("enter: choose") + Dot +
		Subtle("b: previous screen") + Dot +
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

func IsMoneyNumber(input string) error {
	_, err := decimal.NewFromString(input)
	return err
}
