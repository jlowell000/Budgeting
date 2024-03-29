package mainview

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
	"jlowell000.github.io/budgeting/internal/service"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type MainModel struct {
	Choice         int
	Cursor         int
	Selected       map[int]struct{}
	Chosen         bool
	AccountService service.AccountServiceInterface
	FlowService    service.PeriodicFlowServiceInterface
}

var mainChoises = []string{
	"View Periodic Flows",
	"View Accounts",
}

type Model interface {
	tea.Model
	GetMain() *MainModel
	SavaData()
}

func MainUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "down":
			main.Choice++
			if main.Choice > len(mainChoises) {
				main.Choice = len(mainChoises)
			}
		case "up":
			main.Choice--
			if main.Choice < 1 {
				main.Choice = 1
			}
		case "s":
			m.SavaData()

		case "enter":
			main.Chosen = true
			return m, nil
		}
	}

	return m, nil
}

func MainView(m Model) string {
	main := m.GetMain()
	c := main.Choice

	tpl := "What to do?\n\n"
	tpl += util.ProjectionString(main.AccountService, main.FlowService)
	tpl += "%s\n\n"
	tpl += util.Instructions()
	choices := ""
	for i, f := range mainChoises {
		choices += fmt.Sprintf("%s\n", util.Checkbox(f, c == i+1))
	}

	return fmt.Sprintf(tpl, choices)
}
