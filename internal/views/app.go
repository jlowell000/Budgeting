package views

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/muesli/reflow/indent"
	"github.com/muesli/termenv"
)

type AppModel struct {
	Main        MainModel
	FlowList    FlowListModel
	AccountList AccountListModel
	Quitting    bool
}

var (
	Term    = termenv.EnvColorProfile()
	Keyword = makeFgStyle("211")
	Subtle  = makeFgStyle("241")
	Dot     = ColorFg(" • ", "236")
)

func (m AppModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m AppModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			m.Quitting = true
			return m, tea.Quit
		}
	}

	// Hand off the message and model to the appropriate update function for the
	// appropriate view based on the current state.
	if m.Main.Chosen == true {
		if m.Main.Choice == 1 {
			return flowListUpdate(msg, m)
		} else if m.Main.Choice == 2 {
			return accountListUpdate(msg, m)
		}
	}
	return mainUpdate(msg, m)
}

func (m AppModel) View() string {
	var s string
	if m.Quitting {
		return "\n  See you later!\n\n"
	}
	if m.Main.Chosen == true {
		if m.Main.Choice == 1 {
			s = flowListView(m)
		} else if m.Main.Choice == 2 {
			s = accountListView(m)
		}
	} else {
		s = mainView(m)
	}
	return indent.String("\n"+s+"\n\n", 2)
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
