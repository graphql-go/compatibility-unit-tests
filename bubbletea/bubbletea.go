package bubbletea

import (
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type BubbleTea struct {
	cursor  int
	choice  string
	choices []string
	ui      UI
}

type UI struct {
	header string
}

func (b BubbleTea) Init() tea.Cmd {
	return nil
}

func (b BubbleTea) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return b, tea.Quit

		case "enter":
			b.choice = b.choices[b.cursor]
			return b, tea.Quit

		case "down", "j":
			b.cursor++
			if b.cursor >= len(b.choices) {
				b.cursor = 0
			}

		case "up", "k":
			b.cursor--
			if b.cursor < 0 {
				b.cursor = len(b.choices) - 1
			}
		}
	}

	return b, nil
}

func (b BubbleTea) View() string {
	s := strings.Builder{}
	s.WriteString(b.ui.header)
	s.WriteString("")

	for i := 0; i < len(b.choices); i++ {
		if b.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(b.choices[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

type RunResult struct {
	Choice string
}

type BubbleTeaParams struct {
	Choices []string
	UI      BubbleTeaUIParams
}

type BubbleTeaUIParams struct {
	Header string
}

func NewBubbleTea(p *BubbleTeaParams) *BubbleTea {
	return &BubbleTea{
		choices: p.Choices,
		ui: UI{
			header: p.UI.Header,
		},
	}
}

func (b *BubbleTea) Run() (*RunResult, error) {
	teaProgram := tea.NewProgram(b)

	m, err := teaProgram.Run()
	if err != nil {
		return nil, err
	}

	result := &RunResult{}

	if m, ok := m.(BubbleTea); ok && m.choice != "" {
		result.Choice = m.choice
	}

	return result, nil
}
