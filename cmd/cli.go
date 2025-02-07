package cmd

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type CLI struct {
}

type model struct {
	cursor  int
	choice  string
	choices map[string]string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			m.choice = m.choices[m.cursor]
			return m, tea.Quit

		case "down", "j":
			m.cursor++
			if m.cursor >= len(m.choices) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(m.choices) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	s := strings.Builder{}
	s.WriteString("")

	for i := 0; i < len(m.choices); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(m.choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

type RunResult struct {
	Choice string
}

type RunParams struct {
	Choices map[string]string
}

func (c *CLI) Run(p *RunParams) (*RunResult, error) {
	teaProgram := tea.NewProgram(model{
		choices: p.Choices,
	})

	m, err := teaProgram.Run()
	if err != nil {
		return nil, err
	}

	result := &RunResult{}

	if m, ok := m.(model); ok && m.choice != "" {
		result.Choice = m.choice
	}

	return result, nil
}
