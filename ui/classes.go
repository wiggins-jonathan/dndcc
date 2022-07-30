package ui

import (
	"fmt"
	"os"

	"gitlab.com/wiggins.jonathan/dndcc/data"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type classModel struct {
	list     list.Model
	selected string
}

// Instantiates classModel with a list of races
func newClassModel() classModel {
	classes, err := data.ListClassNames()
	if err != nil || len(classes) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	items := make([]list.Item, len(classes))
	for i, class := range classes {
		items[i] = item(class)
	}

	defaultWidth := 24
	listHeight := 24

	l := list.New(items, itemDelegate{}, defaultWidth, listHeight)
	l.Title = "Choose a class:"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return classModel{list: l}
}

func (c classModel) Init() tea.Cmd {
	return nil
}

func (c classModel) Update(msg tea.Msg) (classModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.list.SetWidth(msg.Width)
		return c, nil
	case tea.KeyMsg:
		if c.list.FilterState() == list.Filtering { // don't match if filtering
			break
		}

		switch keypress := msg.String(); keypress {
		case "s":
			if c.list.ShowStatusBar() {
				c.list.SetShowStatusBar(false)
				return c, nil
			}
			c.list.SetShowStatusBar(true)
			return c, nil
		case "q":
			return c, tea.Quit
		}
	}

	var cmd tea.Cmd
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

func (c classModel) View() string {
	return "\n" + c.list.View()
}
