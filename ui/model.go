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
	showBackgrounds
)

type model struct {
	common      *commonModel
	state       state
	races       raceModel
	classes     classModel
	backgrounds backgroundModel
}

// Returns a model with the races view as default
func NewModel() model {
	c := newCommonModel()
	return model{
		common: c,
		state:  showRaces,
		races:  newRaceModel(c),
	}
}

// This type & the FilterValue() method allow filtering lists
type item string

func (i item) FilterValue() string { return string(i) }

func (m model) Init() tea.Cmd {
	return nil
}

// Responds to msg & updates the model state accordingly
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg: // Key presses
		// Don't match keys if filtering
		if m.common.list.FilterState() == list.Filtering {
			break
		}

		switch keypress := msg.String(); keypress {
		case "s":
			if m.common.list.ShowStatusBar() {
				m.common.list.SetShowStatusBar(false)
				return m, nil
			}
			m.common.list.SetShowStatusBar(true)
			return m, nil
		case "esc": // Reset selections & go back to classes
			m.common.list.ResetFilter()
			m.common.list.Select(0)
			m.races = newRaceModel(m.common)
			m.state = showRaces
			return m, nil
		case "enter", " ", "tab": // Save selection & switch state
			switch m.state {
			case showRaces:
				selected, ok := m.common.list.SelectedItem().(item)
				if ok {
					m.races.selected = string(selected)
				}

				m.common.list.ResetFilter()
				m.common.list.Select(0) // Reset cursor to beginning of list

				m.classes = newClassModel(m.common)
				m.state = showClasses

				return m, nil
			case showClasses:
				selected, ok := m.common.list.SelectedItem().(item)
				if ok {
					m.classes.selected = string(selected)

					m.common.list.ResetFilter()
					m.common.list.Select(0) // Reset cursor to beginning of list

					m.backgrounds = newBackgroundModel(m.common)
					m.state = showBackgrounds
				}

				return m, nil
			case showBackgrounds:
				selected, ok := m.common.list.SelectedItem().(item)
				if ok {
					m.backgrounds.selected = string(selected)
				}
				return m, tea.Quit
			}
		}
	}

	var cmd tea.Cmd
	switch m.state {
	case showRaces:
		m.races, cmd = m.races.Update(msg)
		return m, cmd
	case showClasses:
		m.classes, cmd = m.classes.Update(msg)
		return m, cmd
	case showBackgrounds:
		m.backgrounds, cmd = m.backgrounds.Update(msg)
		return m, cmd
	default:
		return m, nil
	}
}

// Prints the UI based on model state
func (m model) View() string {
	if m.races.selected != "" && m.classes.selected != "" && m.backgrounds.selected != "" {
		return quitTextStyle.Render(fmt.Sprintf("Human sacrifice! %s & %s living together! Mass hysteria!", m.races.selected, m.classes.selected))
	}
	switch m.state {
	case showClasses:
		return m.classes.View()
	case showBackgrounds:
		return m.backgrounds.View()
	default:
		return m.races.View()
	}
}
