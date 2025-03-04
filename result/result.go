package result

import (
	"fmt"
	"strconv"

	"github.com/charmbracelet/lipgloss"
)

var successfulStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#4CAF50")).
	PaddingTop(0).
	PaddingLeft(0).
	Width(40)

var failedStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#A52A2A")).
	PaddingTop(0).
	PaddingLeft(0).
	Width(40)

type SummaryParams struct {
	SuccessfulTests int
	FailedTests     int
}

type Result struct {
}

func (r *Result) Summary(s *SummaryParams) string {
	sTests := strconv.Itoa(s.SuccessfulTests)
	fTests := strconv.Itoa(s.FailedTests)

	successfulResult := successfulStyle.Render(fmt.Sprintf("successful tests count: %+v", sTests))
	failedResult := failedStyle.Render(fmt.Sprintf("failed tests count: %+v", fTests))

	return fmt.Sprintf("%s\n%s", successfulResult, failedResult)
}
