package flowlist

import (
	"errors"
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
	"jlowell000.github.io/budgeting/internal/model/period"
	"jlowell000.github.io/budgeting/internal/model/periodicflow"
	"jlowell000.github.io/budgeting/internal/views/form"
	"jlowell000.github.io/budgeting/internal/views/mainview"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type FlowListModel struct {
	/* list of flows */
	Flows    []*periodicflow.PeriodicFlow
	Choice   int
	Cursor   int
	Selected map[int]struct{}
	Chosen   bool

	/* Tell the model how to Create a flows */
	CreateFlowFunc func(string, decimal.Decimal, period.Period) *periodicflow.PeriodicFlow
	/* Tell the model how to get list of flows */
	GetFlowListFunc func() []*periodicflow.PeriodicFlow
	/* Update FlowList */
	UpdateFlowFunc func(uuid.UUID, string, decimal.Decimal, period.Period) *periodicflow.PeriodicFlow
}

type Model interface {
	tea.Model
	GetMain() *mainview.MainModel
	GetFlowList() *FlowListModel
	GetForm() *form.FormModel
}

func FlowListUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	flowList := m.GetFlowList()
	form := m.GetForm()
	flowList.Flows = flowList.GetFlowListFunc()
	checkFormForNewData(flowList, form)

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
			form.LastScreen = 1
			form.Inputs = createFormInputs("", decimal.NewFromFloat(0.0), period.Weekly)
			main.Choice = 3
		case "b":
			main.Chosen = false
		case "enter":
			flowList.Chosen = true
			form.LastScreen = 1
			c := flowList.Choice
			form.Inputs = createFormInputs(
				flowList.Flows[c].Name,
				flowList.Flows[c].Amount,
				flowList.Flows[c].Period,
			)
			main.Choice = 3
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

func createFormInputs(
	n string,
	a decimal.Decimal,
	p period.Period,
) []textinput.Model {
	inputs := make([]textinput.Model, 3)
	var t textinput.Model
	for i := range inputs {
		t = textinput.New()
		t.Cursor.Style = util.CursorStyle
		t.CharLimit = 32

		switch i {
		case 0:
			t.Placeholder = "Name"
			t.SetValue(n)
			t.PromptStyle = util.FocusedStyle
			t.TextStyle = util.FocusedStyle
		case 1:
			t.Placeholder = "Amount"
			t.SetValue(a.String())
			t.Validate = isMoneyNumber
		case 2:
			t.Placeholder = "Period"
			t.SetValue(p.String())
			t.SetSuggestions(period.PeriodStrings[1:])
			t.ShowSuggestions = true
		}

		inputs[i] = t
	}

	return inputs
}

func checkFormForNewData(flowList *FlowListModel, form *form.FormModel) bool {
	if form.Submitted {
		d, _ := decimal.NewFromString(form.Inputs[1].Value())

		if !flowList.Chosen {
			flowList.CreateFlowFunc(
				form.Inputs[0].Value(),
				d,
				period.PeriodFromText(form.Inputs[2].Value()),
			)
		} else {
			flowList.Chosen = false
			flowList.UpdateFlowFunc(
				flowList.Flows[flowList.Choice].Id,
				form.Inputs[0].Value(),
				d,
				period.PeriodFromText(form.Inputs[2].Value()),
			)
		}
		form.ResetForm()
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
