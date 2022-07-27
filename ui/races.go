package ui

import (
	"fmt"
	"os"

	"gitlab.com/wiggins.jonathan/dndcc/data"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type raceModel struct {
	//data  []data.Races
	list   list.Model
	choice string
}

// Instatiates raceModel with a list of races
func newRaceModel() raceModel {
	races, err := data.ListRaceNames()
	if err != nil || len(races) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	var items []list.Item
	for _, race := range races {
		items = append(items, item(race))
	}

	defaultWidth := 24
	listHeight := 24

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Choose a race:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return raceModel{list: l}
}

// This type & the FilterValue() method allow filtering the list
type item string

func (i item) FilterValue() string { return string(i) }

func (r raceModel) Init() tea.Cmd {
	return nil
}

func (r raceModel) Update(msg tea.Msg) (raceModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.list.SetWidth(msg.Width)
		return r, nil
	case tea.KeyMsg:
		if r.list.FilterState() == list.Filtering { // don't match if filtering
			break
		}

		switch keypress := msg.String(); keypress {
		case "s":
			if r.list.ShowStatusBar() {
				r.list.SetShowStatusBar(false)
				return r, nil
			}
			r.list.SetShowStatusBar(true)
			return r, nil
		case "q":
			return r, tea.Quit
		}
	}

	var cmd tea.Cmd
	r.list, cmd = r.list.Update(msg)
	return r, cmd
}

func (r raceModel) View() string {
	return "\n" + r.list.View()
}
