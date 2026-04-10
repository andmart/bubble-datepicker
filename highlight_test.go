package datepicker

import (
	"testing"
	"time"

	"github.com/charmbracelet/lipgloss"
)

func TestHighlight_IsHighlighted_Found(t *testing.T) {
	h := Highlight{
		Dates: map[string]lipgloss.Style{
			"2026-04-10": lipgloss.NewStyle().Bold(true),
		},
	}

	date := time.Date(2026, 4, 10, 0, 0, 0, 0, time.UTC)

	style, ok := h.IsHighlighted(date)

	if !ok {
		t.Fatalf("expected date to be highlighted")
	}

	if style.GetBold() != true {
		t.Fatalf("expected style to be bold")
	}
}

func TestHighlight_IsHighlighted_NotFound(t *testing.T) {
	h := Highlight{
		Dates: map[string]lipgloss.Style{
			"2026-04-10": lipgloss.NewStyle(),
		},
	}

	date := time.Date(2026, 4, 11, 0, 0, 0, 0, time.UTC)

	_, ok := h.IsHighlighted(date)

	if ok {
		t.Fatalf("expected date NOT to be highlighted")
	}
}

func TestHighlight_IsHighlighted_NilMap(t *testing.T) {
	h := Highlight{
		Dates: nil,
	}

	date := time.Now()

	_, ok := h.IsHighlighted(date)

	if ok {
		t.Fatalf("expected false for nil map")
	}
}