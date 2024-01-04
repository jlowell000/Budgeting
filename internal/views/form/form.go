package form

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/cursor"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"jlowell000.github.io/budgeting/internal/views/mainview"
	"jlowell000.github.io/budgeting/internal/views/util"
)

type FormModel struct {
	LastScreen int
	Submitted  bool
	FocusIndex int
	Inputs     []textinput.Model
	CursorMode cursor.Mode
}

type Model interface {
	tea.Model
	GetMain() *mainview.MainModel
	GetForm() *FormModel
	UpdateInputs(msg tea.Msg) tea.Cmd
}

var ()

func FormUpdate(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	main := m.GetMain()
	form := m.GetForm()
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "b":
			main.Chosen = true
			main.Choice = form.LastScreen
			form.Submitted = false
		// Change cursor mode
		case "ctrl+r":
			form.CursorMode++
			if form.CursorMode > cursor.CursorHide {
				form.CursorMode = cursor.CursorBlink
			}
			cmds := make([]tea.Cmd, len(form.Inputs))
			for i := range form.Inputs {
				cmds[i] = form.Inputs[i].Cursor.SetMode(form.CursorMode)
			}
			return m, tea.Batch(cmds...)

		// Set focus to next input
		case "enter", "up", "down":
			s := msg.String()

			// Did the user press enter while the submit button was focused?
			// If so, exit.
			if s == "enter" && form.FocusIndex == len(form.Inputs) {
				main.Chosen = true
				main.Choice = form.LastScreen
				form.Submitted = true
				return m, nil
			}

			// Cycle indexes
			if s == "up" {
				form.FocusIndex--
			} else {
				form.FocusIndex++
			}

			if form.FocusIndex > len(form.Inputs) {
				form.FocusIndex = 0
			} else if form.FocusIndex < 0 {
				form.FocusIndex = len(form.Inputs)
			}

			cmds := make([]tea.Cmd, len(form.Inputs))
			for i := 0; i <= len(form.Inputs)-1; i++ {
				if i == form.FocusIndex {
					// Set focused state
					cmds[i] = form.Inputs[i].Focus()
					form.Inputs[i].PromptStyle = util.FocusedStyle
					form.Inputs[i].TextStyle = util.FocusedStyle
					continue
				}
				// Remove focused state
				form.Inputs[i].Blur()
				form.Inputs[i].PromptStyle = util.NoStyle
				form.Inputs[i].TextStyle = util.NoStyle
			}

			return m, tea.Batch(cmds...)
		}
	}

	// Handle character input and blinking
	cmd := m.UpdateInputs(msg)

	return m, cmd
}

func FormView(m Model) string {
	form := m.GetForm()
	var b strings.Builder

	for i := range form.Inputs {
		b.WriteString(form.Inputs[i].View())
		if i < len(form.Inputs)-1 {
			b.WriteRune('\n')
		}
	}

	button := &util.BlurredButton
	if form.FocusIndex == len(form.Inputs) {
		button = &util.FocusedButton
	}
	fmt.Fprintf(&b, "\n\n%s\n\n", *button)

	b.WriteString(util.HelpStyle.Render("cursor mode is "))
	b.WriteString(util.CursorModeHelpStyle.Render(form.CursorMode.String()))
	b.WriteString(util.HelpStyle.Render(" (ctrl+r to change style)"))

	return b.String()
}

func (f *FormModel) ResetForm() {
	f.LastScreen = 0
	f.FocusIndex = 0
	f.Inputs = make([]textinput.Model, 1)
	f.Submitted = false
}
