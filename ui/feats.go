package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"gitlab.com/wiggins.jonathan/dndcc/data"

	tea "github.com/charmbracelet/bubbletea"
)

type featModel struct {
	data     []data.Feats
	list     list.Model
	selected string
}

// Instantiates featModel with a list of feats
func NewFeatModel() featModel {
	d, err := data.NewFeats()
	if err != nil || len(d) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	// Inject feats into list
	feats := data.ListFeatNames(d)
	items := stringToItem(feats)
	l := newListModel()
	l.SetItems(items)

	l.Title = "Choose a feat:"

	return featModel{list: l, data: d}
}

func (f featModel) Init() tea.Cmd {
	return nil
}

func (f featModel) Update(msg tea.Msg) (featModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		f.list.SetWidth(msg.Width)
		return f, nil
	case tea.KeyMsg:
		if f.list.FilterState() == list.Filtering {
			break // don't match keys if filtering
		}
		switch keypress := msg.String(); keypress {
		case "esc": // Reset selection
			f.list.ResetFilter()
			f.list.Select(0)
			return f, nil
		case "s":
			if f.list.ShowStatusBar() {
				f.list.SetShowStatusBar(false)
				return f, nil
			}
			f.list.SetShowStatusBar(true)
			return f, nil
		}
	}

	var cmd tea.Cmd
	f.list, cmd = f.list.Update(msg)
	return f, cmd
}

func (f featModel) View() string {
	return "\n" + f.list.View()
}
