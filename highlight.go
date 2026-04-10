package datepicker

import (
	"time"

	"github.com/charmbracelet/lipgloss"
)

type Highlight struct{
	// Dates is a map of date strings to lipgloss styles. The date strings should be in the format "2006-01-02" (YYYY-MM-DD).
	Dates map[string] lipgloss.Style
}

func (h Highlight) IsHighlighted(t time.Time) (lipgloss.Style, bool) {	
	dateStr := t.Format("2006-01-02")
	style, ok := h.Dates[dateStr]
	return style, ok
}