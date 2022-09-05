package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"gitlab.com/wiggins.jonathan/dndcc/data"

	tea "github.com/charmbracelet/bubbletea"
)

type backgroundModel struct {
	data     []data.Backgrounds
	list     list.Model
	selected string
	footer   *footer
}

// Instantiates backgroundModel with a list of backgrounds
func NewBackgroundModel(f *footer) backgroundModel {
	d, err := data.NewBackgrounds()
	if err != nil || len(d) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	// Inject backgrounds into list
	backgrounds := data.ListBackgroundNames(d)
	items := stringToItem(backgrounds)
	l := newListModel()
	l.SetItems(items)

	l.Title = "Choose a background:"

	return backgroundModel{list: l, data: d, footer: f}
}

func (b backgroundModel) Init() tea.Cmd {
	return nil
}

func (b backgroundModel) Update(msg tea.Msg) (backgroundModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		b.list.SetWidth(msg.Width)
		return b, nil
	case tea.KeyMsg:
		if b.list.FilterState() == list.Filtering {
			break // don't match keys if filtering
		}

		switch keypress := msg.String(); keypress {
		case "enter", " ", "tab":
			selected, ok := b.list.SelectedItem().(item)
			if ok {
				b.selected = string(selected)
				b.footer.BackgroundSelected = b.selected
			}
			return b, nil
		case "esc": // Reset selection
			b.list.ResetFilter()
			b.list.Select(0)
			b.footer.BackgroundSelected = "Background"
			return b, nil
		case "s":
			if b.list.ShowStatusBar() {
				b.list.SetShowStatusBar(false)
				return b, nil
			}
			b.list.SetShowStatusBar(true)
			return b, nil
		}
	}

	var cmd tea.Cmd
	b.list, cmd = b.list.Update(msg)
	return b, cmd
}

func (b backgroundModel) View() string {
	return "\n" + b.list.View()
}
