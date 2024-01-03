package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type MainModel struct {
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool
}

var mainChoises = []string{
	"View Periodic Flows",
	"View Accounts",
}

func mainUpdate(msg tea.Msg, m AppModel) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			m.Main.Choice++
			if m.Main.Choice > len(mainChoises) {
				m.Main.Choice = len(mainChoises)
			}
		case "k", "up":
			m.Main.Choice--
			if m.Main.Choice < 1 {
				m.Main.Choice = 1
			}
		case "enter":
			m.Main.Chosen = true
			return m, nil
		}
	}

	return m, nil
}

func mainView(m AppModel) string {
	c := m.Main.Choice

	tpl := "What to do?\n\n"
	tpl += "%s\n\n"
	tpl += Subtle("j/k, up/down: select") + Dot + Subtle("enter: choose") + Dot + Subtle("q, esc: quit")

	choices := ""
	for i, f := range mainChoises {
		choices += fmt.Sprintf("%s\n", Checkbox(f, c == i+1))
	}

	return fmt.Sprintf(tpl, choices)
}
