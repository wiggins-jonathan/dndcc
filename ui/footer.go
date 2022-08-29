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
	raceSelected := itemStyle.Render(f.RaceSelected)
	classSelected := itemStyle.Render(f.ClassSelected)
	backgroundSelected := itemStyle.Render(f.BackgroundSelected)

	return raceSelected + classSelected + backgroundSelected
}
