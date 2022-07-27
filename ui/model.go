package ui

import (
	"fmt"
	"os"

	"gitlab.com/wiggins.jonathan/dndcc/data"

	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2).Foreground(lipgloss.Color("#0099ff"))
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().MarginLeft(4).Background(lipgloss.Color("#ff3399")).Bold(true)
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type model struct {
	list   list.Model
	choice string
}

func NewModel() model {
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

	return model{list: l}
}

type item string

func (i item) FilterValue() string { return string(i) }

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.list.SetWidth(msg.Width)
		return m, nil

	case tea.KeyMsg: // Key presses
		if m.list.FilterState() == list.Filtering { // don't match if filtering
			break
		}

		switch keypress := msg.String(); keypress {
		case "s": // toggle statusbar
			if m.list.ShowStatusBar() {
				m.list.SetShowStatusBar(false)
				return m, nil
			}
			m.list.SetShowStatusBar(true)
			return m, nil

		case "enter", " ":
			i, ok := m.list.SelectedItem().(item)
			if ok {
				m.choice = string(i)
			}
			return m, tea.Quit
		}
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.choice != "" {
		return quitTextStyle.Render(fmt.Sprintf("%s? Sounds good to me.", m.choice))
	}
	return "\n" + m.list.View()
}
