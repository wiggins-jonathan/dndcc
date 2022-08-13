package ui

import (
	"fmt"
	"io"

	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
)

func newListModel() list.Model {
	listWidth := 24
	listHeight := 24
	l := list.New([]list.Item{}, itemDelegate{}, listWidth, listHeight)

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	// Custom key definitions in help
	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithHelp("tab", "select")),
			key.NewBinding(key.WithHelp("shift+tab", "back")),
		}
	}

	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithHelp("enter/tab/space", "select")),
			key.NewBinding(key.WithHelp("shift+tab", "back")),
			key.NewBinding(key.WithHelp("esc", "reset")),
			key.NewBinding(key.WithHelp("s", "status bar")),
		}
	}

	return l
}

// This type & the FilterValue() method allow filtering lists
type item string

func (i item) FilterValue() string { return string(i) }

type itemDelegate struct{}

func (d itemDelegate) Height() int                               { return 1 }
func (d itemDelegate) Spacing() int                              { return 0 }
func (d itemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }

func (d itemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(item)
	if !ok {
		return
	}

	str := fmt.Sprintf("%s", i)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render(s)
		}
	}

	fmt.Fprintf(w, fn(str))
}

// Utility function to transform a []string to []list.Item
func stringToItem(s []string) []list.Item {
	items := make([]list.Item, len(s))
	for i, v := range s {
		items[i] = item(v)
	}

	return items
}
