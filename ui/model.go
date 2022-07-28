package ui

import (
	"fmt"
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
	showClasses
)

type model struct {
	state   state
	races   raceModel
	classes classModel
}

// Returns a model with the races view as default
func NewModel() model {
	return model{state: showRaces, races: newRaceModel()}
}

// This type & the FilterValue() method allow filtering lists
type item string

func (i item) FilterValue() string { return string(i) }

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
		case "enter", " ": // Save selection & switch state
			switch m.state {
			case showRaces:
				selected, ok := m.races.list.SelectedItem().(item)
				if ok {
					m.races.selected = string(selected)
				}
				m.classes = newClassModel()
				m.state = showClasses
				return m, cmd
			case showClasses:
				selected, ok := m.classes.list.SelectedItem().(item)
				if ok {
					m.classes.selected = string(selected)
				}
				return m, tea.Quit
			}
		}
	}

	switch m.state {
	case showRaces:
		m.races, cmd = m.races.Update(msg)
		return m, cmd
	case showClasses:
		m.classes, cmd = m.classes.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

// Prints the UI based on model state
func (m model) View() string {
	if m.races.selected != "" && m.classes.selected != "" {
		return quitTextStyle.Render(fmt.Sprintf("Human sacrifice! %s & %s living together! Mass hysteria!", m.races.selected, m.classes.selected))
	}
	switch m.state {
	case showClasses:
		return m.classes.View()
	default:
		return m.races.View()
	}
}
