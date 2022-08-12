package ui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	lg                = lipgloss.NewStyle()
	titleStyle        = lg.MarginLeft(2).Foreground(lipgloss.Color("#0099ff"))
	itemStyle         = lg.PaddingLeft(4)
	selectedItemStyle = lg.MarginLeft(4).Background(lipgloss.Color("#ff3399")).Bold(true)
	quitTextStyle     = lg.Margin(1, 0, 2, 4)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
)

type state int

const (
	showRaces state = iota
	showClasses
	showBackgrounds
)

type model struct {
	common *commonModel
	state  state

	// submodels
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
		case "s": // Toggle the status bar
			if m.common.list.ShowStatusBar() {
				m.common.list.SetShowStatusBar(false)
				return m, nil
			}
			m.common.list.SetShowStatusBar(true)
			return m, nil
		case "esc": // Reset selections & go back to races
			m.common.list.ResetFilter()
			m.common.list.Select(0)
			m.races = newRaceModel(m.common)
			m.state = showRaces
			return m, nil
		case "enter", " ", "tab": // Save selection & switch state
			switch m.state {
			case showRaces:
				selected, ok := m.common.list.SelectedItem().(item)
				if ok { // Get selected item & its position in the list
					m.races.selected = string(selected)
					m.races.selectedIndex = m.common.list.Index()
				}

				m.common.list.ResetFilter()
				m.common.list.Select(0) // Reset cursor to beginning of list

				m.classes = newClassModel(m.common)
				m.state++

				return m, nil
			case showClasses:
				selected, ok := m.common.list.SelectedItem().(item)
				if ok {
					m.classes.selected = string(selected)
					m.classes.selectedIndex = m.common.list.Index()
				}

				m.common.list.ResetFilter()
				m.common.list.Select(0) // Reset cursor to beginning of list

				m.backgrounds = newBackgroundModel(m.common)
				m.state++

				return m, nil
			case showBackgrounds:
				selected, ok := m.common.list.SelectedItem().(item)
				if ok {
					m.backgrounds.selected = string(selected)
					m.classes.selectedIndex = m.common.list.Index()
				}
				return m, tea.Quit
			}
		case "shift+tab": // Go back to the previous selection
			switch m.state {
			case showClasses:
				m.common.list.ResetFilter()
				m.common.list.Select(m.races.selectedIndex) // we need to return to races.selected
				m.races = newRaceModel(m.common)
				m.state--
				return m, nil
			case showBackgrounds:
				m.common.list.ResetFilter()
				m.common.list.Select(m.classes.selectedIndex) // we need to return to races.selected
				m.classes = newClassModel(m.common)
				m.state--
				return m, nil
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
	}

	return m, nil
}

// Prints the UI based on model state
func (m model) View() string {
	// Temporary exit message
	if m.races.selected != "" &&
		m.classes.selected != "" &&
		m.backgrounds.selected != "" {
		return quitTextStyle.Render(fmt.Sprintf(
			"%s! %s & %s living together! Mass hysteria!",
			m.races.selected,
			m.classes.selected,
			m.backgrounds.selected,
		))
	}
	switch m.state {
	case showRaces:
		return m.races.View()
	case showClasses:
		return m.classes.View()
	case showBackgrounds:
		return m.backgrounds.View()
	}

	return ""
}
