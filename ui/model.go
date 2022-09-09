package ui

import (
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	minHeight = 4
	minWidth  = 30
)

type state int

const (
	showRaces state = iota
	showClasses
	showBackgrounds
	showFeats
)

type model struct {
	state  state
	height int
	width  int
	footer *footer

	// submodels
	races       raceModel
	classes     classModel
	backgrounds backgroundModel
	feats       featModel
}

// Returns a model with the races view as default
func NewModel() model {
	f := NewFooter()
	return model{
		state:       showRaces,
		footer:      f,
		races:       NewRaceModel(f),
		classes:     NewClassModel(f),
		backgrounds: NewBackgroundModel(f),
		feats:       NewFeatModel(),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

// Responds to msg & updates the model state accordingly
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
	case tea.KeyMsg: // Key presses
		if m.isFiltering() {
			break // Don't match keys if filtering
		}

		switch keypress := msg.String(); keypress {
		case "enter", " ", "tab": // Save selection & switch state
			switch m.state {
			case showRaces:
				m.races, cmd = m.races.Update(msg)
				m.state++
				return m, cmd
			case showClasses:
				m.classes, cmd = m.classes.Update(msg)
				m.state++
				return m, cmd
			case showBackgrounds:
				m.backgrounds, cmd = m.backgrounds.Update(msg)
				m.state++
				return m, cmd
			case showFeats:
				m.feats, cmd = m.feats.Update(msg)
				return m, cmd
			}
		case "shift+tab": // Go back to the previous selection
			switch m.state {
			case showClasses:
				m.classes, cmd = m.classes.Update(msg)
				m.state--
				return m, cmd
			case showBackgrounds:
				m.backgrounds, cmd = m.backgrounds.Update(msg)
				m.state--
				return m, cmd
			case showFeats:
				m.feats, cmd = m.feats.Update(msg)
				m.state--
				return m, cmd
			}
		case "ctrl+c":
			return m, tea.Quit
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
	if m.height < minHeight || m.width < minWidth {
		return "Window too small.\nEnlarge viewing area."
	}

	var view string
	switch m.state {
	case showRaces:
		view = m.races.View()
	case showClasses:
		view = m.classes.View()
	case showBackgrounds:
		view = m.backgrounds.View()
	case showFeats:
		view = m.feats.View()
	}

	return lipgloss.JoinVertical(lipgloss.Top, view, m.footer.View())
}

// clean this up with reflection
func (m model) isFiltering() bool {
	if m.races.list.FilterState() == list.Filtering ||
		m.classes.list.FilterState() == list.Filtering ||
		m.backgrounds.list.FilterState() == list.Filtering ||
		m.feats.list.FilterState() == list.Filtering {
		return true
	}
	return false
}
