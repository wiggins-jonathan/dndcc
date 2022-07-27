package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("#0099ff"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().MarginLeft(4).Background(lipgloss.Color("#ff3399")).Bold(true)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type state int

const (
	showRaces state = iota
)

type model struct {
	state state
	races raceModel
}

// Returns a model with the races view as default
func NewModel() model {
	return model{state: showRaces, races: newRaceModel()}
}

func (m model) Init() tea.Cmd {
	return nil
}

// Responds to msg & updates the model state accordingly
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg: // Key presses
		switch keypress := msg.String(); keypress {
		case "ctrl+c": // Exit program no matter the state
			return m, tea.Quit
		}
	}

	switch m.state {
	default:
		m.races, cmd = m.races.Update(msg)
		return m, cmd
	}

	return m, nil
}

// Prints the UI based on model state
func (m model) View() string {
	switch m.state {
	default:
		return m.races.View()
	}
}
