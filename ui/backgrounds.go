package ui

import (
	"fmt"
	"os"

	"gitlab.com/wiggins.jonathan/dndcc/data"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

type backgroundModel struct {
	common        *commonModel
	selected      string
	selectedIndex int
}

// Instantiates raceModel with a list of races
func newBackgroundModel(common *commonModel) backgroundModel {
	backgroundData, err := data.NewBackgrounds()
	if err != nil || len(backgroundData) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	// Inject races into common list
	backgrounds := data.ListBackgroundNames(backgroundData)
	items := make([]list.Item, len(backgrounds))
	for i, background := range backgrounds {
		items[i] = item(background)
	}
	common.list.SetItems(items)

	common.list.Title = "Choose a background:"

	return backgroundModel{common: common}
}

func (b backgroundModel) Init() tea.Cmd {
	return nil
}

func (b backgroundModel) Update(msg tea.Msg) (backgroundModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.common.list.SetWidth(msg.Width)
		return b, nil
	}

	var cmd tea.Cmd
	b.common.list, cmd = b.common.list.Update(msg)
	return b, cmd
}

func (b backgroundModel) View() string {
	return "\n" + b.common.list.View()
}
