package ui

import (
	"fmt"
	"os"

	"gitlab.com/wiggins.jonathan/dndcc/data"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type classModel struct {
	common   *commonModel
	selected string
}

// Instantiates classModel with a list of races
func newClassModel(common *commonModel) classModel {
	classes, err := data.ListClassNames()
	if err != nil || len(classes) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	// Inject classes into common list
	items := make([]list.Item, len(classes))
	for i, class := range classes {
		items[i] = item(class)
	}
	common.list.SetItems(items)

	common.list.Title = "Choose a class:"

	return classModel{common: common}
}

func (c classModel) Init() tea.Cmd {
	return nil
}

func (c classModel) Update(msg tea.Msg) (classModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		c.common.list.SetWidth(msg.Width)
		return c, nil
	}

	var cmd tea.Cmd
	c.common.list, cmd = c.common.list.Update(msg)
	return c, cmd
}

func (c classModel) View() string {
	return "\n" + c.common.list.View()
}
