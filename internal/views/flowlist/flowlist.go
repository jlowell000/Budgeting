package flowlist

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
	"jlowell000.github.io/budgeting/internal/views/flowform"
	"jlowell000.github.io/budgeting/internal/views/mainview"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type FlowListModel struct {
	Flows    []*periodicflow.PeriodicFlow // list of flows
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool

	GetFlowListFunc func() []*periodicflow.PeriodicFlow
	CreateFlowFunc  func(
		string,
		decimal.Decimal,
		period.Period,
	) *periodicflow.PeriodicFlow
}

type Model interface {
	tea.Model
	GetMain() *mainview.MainModel
	GetFlowList() *FlowListModel
	GetFlowForm() *flowform.FlowFormModel
}

func FlowListUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	flowList := m.GetFlowList()
	flowform := m.GetFlowForm()
	flowList.Flows = flowList.GetFlowListFunc()
	checkFormForNewData(flowList, m.GetFlowForm())

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "j", "down":
			flowList.Choice++
			if flowList.Choice > len(flowList.Flows)-1 {
				flowList.Choice = len(flowList.Flows) - 1
			}
		case "k", "up":
			flowList.Choice--
			if flowList.Choice < 0 {
				flowList.Choice = 0
			}

		case "n":
			flowform.LastScreen = 1
			flowform.Inputs = createFormInputs()
			main.Choice = 3
		case "b":
			main.Chosen = false
		case "enter":
			flowList.Chosen = true
		}
	}

	return m, nil
}

func FlowListView(m Model) string {
	flowList := m.GetFlowList()
	c := flowList.Choice
	// The header
	tpl := "Viewing Periodic Flows\n\n"
	tpl += "%s\n\n"
	tpl += util.Instructions()
	tpl += util.Dot + util.Subtle("n to create new") + util.Dot

	flows := ""
	for i, f := range flowList.Flows {
		flows += fmt.Sprintf("%s\n", util.Checkbox(f.String(), c == i))
	}

	return fmt.Sprintf(tpl, flows)
}

func createFormInputs() []textinput.Model {
	inputs := make([]textinput.Model, 3)
	var t textinput.Model
	for i := range inputs {
		t = textinput.New()
		t.Cursor.Style = util.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.Focus()
			t.PromptStyle = util.FocusedStyle
			t.TextStyle = util.FocusedStyle
		case 1:
			t.Placeholder = "Amount"
			t.Validate = isMoneyNumber
		case 2:
			t.Placeholder = "Period"
			t.SetSuggestions(period.PeriodStrings[1:])
			t.ShowSuggestions = true
		}

		inputs[i] = t
	}

	return inputs
}

func checkFormForNewData(flowList *FlowListModel, flowform *flowform.FlowFormModel) bool {
	if flowform.Submitted {
		d, _ := decimal.NewFromString(flowform.Inputs[1].Value())
		// flow :=
		flowList.CreateFlowFunc(
			flowform.Inputs[0].Value(),
			d,
			period.PeriodFromText(flowform.Inputs[2].Value()),
		)
		// flowList.Flows = append(flowList.Flows, flow)
		flowform.LastScreen = 0
		flowform.FocusIndex = 0
		flowform.Inputs = make([]textinput.Model, 0)
		flowform.Submitted = false
		return true
	}
	return false
}

func isMoneyNumber(input string) error {
	_, err := decimal.NewFromString(input)
	return err
}

func isPeriodString(input string) error {
	p := period.PeriodFromText(input)
	if p != period.Unknown {
		return nil
	}
	return errors.New("Unknown Period")
}
