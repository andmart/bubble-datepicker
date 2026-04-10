package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/ethanefung/bubble-datepicker"
)

type model struct {
	datepicker datepicker.Model
}

func initialModel() tea.Model {
	now := time.Now()
	holidaysStyles := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#FFac00"))
	vacationsStyles := lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("#00FFac"))

	holidays := []string{
		fmt.Sprintf("%d-%02d-01", now.Year(), int(now.Month())),
		fmt.Sprintf("%d-%02d-17", now.Year(), int(now.Month())),
		fmt.Sprintf("%d-%02d-22", now.Year(), int(now.Month())),
	}

	vacations := []string{
		fmt.Sprintf("%d-%02d-02", now.Year(), int(now.Month())),
		fmt.Sprintf("%d-%02d-03", now.Year(), int(now.Month())),
		fmt.Sprintf("%d-%02d-04", now.Year(), int(now.Month())),
	}

	daysToHighlight := map[string]lipgloss.Style{}

	for _, holiday := range holidays {
		daysToHighlight[holiday] = holidaysStyles
	}

	for _, vacation := range vacations {
		daysToHighlight[vacation] = vacationsStyles
	}

	highlight := datepicker.Highlight{Dates: daysToHighlight}

	dp := datepicker.NewWithHighlight(now, highlight)
	dp.SelectDate()

	return model{
		datepicker: dp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit
		default:
			datepicker, cmd := m.datepicker.Update(msg)
			m.datepicker = datepicker
			return m, cmd
		}
	}

	return m, nil
}

func (m model) View() string {

	// Send the UI for rendering
	return m.datepicker.View()
}

func main() {
	p := tea.NewProgram(initialModel())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
