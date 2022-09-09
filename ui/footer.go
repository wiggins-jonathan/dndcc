package ui

import (
	"github.com/charmbracelet/bubbles/help"
	tea "github.com/charmbracelet/bubbletea"
)

type footer struct {
	help               help.Model
	keymap             *keyMap
	RaceSelected       string
	ClassSelected      string
	BackgroundSelected string
}

func NewFooter() *footer {
	h := help.New()
	k := DefaultKeyMap()
	return &footer{h, k, "Race", "Class", "Background"}
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

	help := helpStyle.Render("\n" + f.help.View(f.keymap))
	return "\n" + raceSelected + classSelected + backgroundSelected + help
}

type ToggleFooterMsg struct{}

// sends a msg to toggle the help in the footer
func (f *footer) ToggleHelp() ToggleFooterMsg {
	return ToggleFooterMsg{}
}
