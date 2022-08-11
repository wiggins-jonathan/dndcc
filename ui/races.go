package ui

import (
	"fmt"
	"os"

	"gitlab.com/wiggins.jonathan/dndcc/data"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type raceModel struct {
	common        *commonModel
	selected      string
	selectedIndex int
}

// Instantiates raceModel with a list of races
func newRaceModel(common *commonModel) raceModel {
	races, err := data.ListRaceNames()
	if err != nil || len(races) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	// Inject races into common list
	items := make([]list.Item, len(races))
	for i, race := range races {
		items[i] = item(race)
	}
	common.list.SetItems(items)

	common.list.Title = "Choose a race:"

	return raceModel{common: common}
}

func (r raceModel) Init() tea.Cmd {
	return nil
}

func (r raceModel) Update(msg tea.Msg) (raceModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.common.list.SetWidth(msg.Width)
		return r, nil
	}

	var cmd tea.Cmd
	r.common.list, cmd = r.common.list.Update(msg)
	return r, cmd
}

func (r raceModel) View() string {
	return "\n" + r.common.list.View()
}
