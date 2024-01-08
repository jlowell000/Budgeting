package util

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/muesli/termenv"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/service"
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
		Subtle("s: save data") + Dot +
		Subtle("q, esc: quit")
}

func ProjectionString(accountService service.AccountServiceInterface, flowService service.PeriodicFlowServiceInterface) string {
	inflow := flowService.GetTotalWeeklyInflow()
	outflow := flowService.GetTotalWeeklyOutflow()
	totalflow := flowService.GetTotalWeeklyFlow()
	projection := flowService.GetProjectedTotalFlow(decimal.NewFromInt(6), period.Monthly)
	accountsTotal := accountService.GetTotal(true)

	return "\nTotal Inflows: " + inflow.String() + Dot +
		"Total Outflows: " + outflow.String() + Dot +
		"Total Flows: " + totalflow.String() + "\n" +
		"Accounts Total: " + accountsTotal.String() + "\n" +
		"6 Month Projected Change: " + projection.String() + Dot +
		"6 Month Projected Change: " + accountsTotal.Add(projection).String() + "\n\n"
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

func TimeFormat(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
