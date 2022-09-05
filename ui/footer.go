package ui

import (
	tea "github.com/charmbracelet/bubbletea"
)

type footer struct {
	RaceSelected       string
	ClassSelected      string
	BackgroundSelected string
}

func NewFooter() *footer {
	return &footer{"Race", "Class", "Background"}
}

func (f *footer) Init() tea.Cmd {
	return nil
}

func (f *footer) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return f, nil
}

func (f *footer) View() string {
	raceSelected := footerUnselected.Render(f.RaceSelected)
	if f.RaceSelected != "Race" {
		raceSelected = footerSelected.Render(f.RaceSelected)
	}

	classSelected := footerUnselected.Render(f.ClassSelected)
	if f.ClassSelected != "Class" {
		classSelected = footerSelected.Render(f.ClassSelected)
	}

	backgroundSelected := footerUnselected.Render(f.BackgroundSelected)
	if f.BackgroundSelected != "Background" {
		backgroundSelected = footerSelected.Render(f.BackgroundSelected)
	}

	return "\n" + raceSelected + classSelected + backgroundSelected
}
