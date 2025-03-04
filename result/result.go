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

	successfulSummary := fmt.Sprintf("successful compatible tests count: %+v", sTests)
	failedSummary := fmt.Sprintf("failed compatible tests count: %+v", fTests)

	successfulResult := successfulStyle.Render(successfulSummary)
	failedResult := failedStyle.Render(failedSummary)

	return fmt.Sprintf("%s\n%s", successfulResult, failedResult)
}
