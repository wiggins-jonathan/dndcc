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
	showFeats
)

type model struct {
	state state

	// submodels
	races       raceModel
	classes     classModel
	backgrounds backgroundModel
	feats       featModel
}

// Returns a model with the races view as default
func NewModel() model {
	return model{
		state:       showRaces,
		races:       newRaceModel(),
		classes:     newClassModel(),
		backgrounds: newBackgroundModel(),
		feats:       newFeatModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// Responds to msg & updates the model state accordingly
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg: // Key presses
		// Don't match keys if filtering
		if m.races.list.FilterState() == list.Filtering ||
			m.classes.list.FilterState() == list.Filtering ||
			m.backgrounds.list.FilterState() == list.Filtering {
			break
		}

		switch keypress := msg.String(); keypress {
		case "enter", " ", "tab": // Save selection & switch state
			switch m.state {
			case showRaces:
				selected, ok := m.races.list.SelectedItem().(item)
				if ok { // Set selected item
					m.races.selected = string(selected)
				}

				m.classes, cmd = m.classes.Update(msg)
				m.state++

				return m, cmd
			case showClasses:
				selected, ok := m.classes.list.SelectedItem().(item)
				if ok {
					m.classes.selected = string(selected)
				}

				m.backgrounds, cmd = m.backgrounds.Update(msg)
				m.state++

				return m, cmd
			case showBackgrounds:
				selected, ok := m.backgrounds.list.SelectedItem().(item)
				if ok {
					m.backgrounds.selected = string(selected)
				}

				m.feats, cmd = m.feats.Update(msg)
				m.state++

				return m, cmd
			case showFeats:
				selected, ok := m.feats.list.SelectedItem().(item)
				if ok {
					m.feats.selected = string(selected)
				}
				return m, tea.Quit
			}
		case "shift+tab": // Go back to the previous selection
			switch m.state {
			case showClasses:
				m.races, cmd = m.races.Update(msg)
				m.state--
				return m, cmd
			case showBackgrounds:
				m.classes, cmd = m.classes.Update(msg)
				m.state--
				return m, cmd
			case showFeats:
				m.backgrounds, cmd = m.backgrounds.Update(msg)
				m.state--
				return m, cmd
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
	case showBackgrounds:
		m.backgrounds, cmd = m.backgrounds.Update(msg)
		return m, cmd
	case showFeats:
		m.feats, cmd = m.feats.Update(msg)
		return m, cmd
	}

	return m, cmd
}

// Prints the UI based on model state
func (m model) View() string {
	// Temporary exit message
	if m.races.selected != "" &&
		m.classes.selected != "" &&
		m.backgrounds.selected != "" &&
		m.feats.selected != "" {
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
	case showFeats:
		return m.feats.View()
	}

	return ""
}
