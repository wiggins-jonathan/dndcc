package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"gitlab.com/wiggins.jonathan/dndcc/data"

	tea "github.com/charmbracelet/bubbletea"
)

type raceModel struct {
	data     []data.Races
	list     list.Model
	selected string
}

// Instantiates raceModel with a list of races
func newRaceModel() raceModel {
	d, err := data.NewRaces()
	if err != nil || len(d) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	// Inject races into list
	races := data.ListRaceNames(d)
	items := stringToItem(races)
	l := newListModel()
	l.SetItems(items)

	l.Title = "Choose a race:"

	return raceModel{list: l, data: d}
}

func (r raceModel) Init() tea.Cmd {
	return nil
}

func (r raceModel) Update(msg tea.Msg) (raceModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.list.SetWidth(msg.Width)
		return r, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc": // Reset selection
			r.list.ResetFilter()
			r.list.Select(0)
			return r, nil
		case "s":
			if r.list.ShowStatusBar() {
				r.list.SetShowStatusBar(false)
				return r, nil
			}
			r.list.SetShowStatusBar(true)
			return r, nil
		}
	}

	var cmd tea.Cmd
	r.list, cmd = r.list.Update(msg)
	return r, cmd
}

func (r raceModel) View() string {
	return "\n" + r.list.View()
}
