package ui

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type commonModel struct {
	list list.Model
}

func newCommonModel() *commonModel {
	listWidth := 24
	listHeight := 24
	l := list.New([]list.Item{}, itemDelegate{}, listWidth, listHeight)

	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(true)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	l.AdditionalFullHelpKeys = func() []key.Binding {
		return []key.Binding{
			key.NewBinding(key.WithHelp("s", "status bar")),
		}
	}

	return &commonModel{list: l}
}
