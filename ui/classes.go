package ui

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"gitlab.com/wiggins.jonathan/dndcc/data"

	tea "github.com/charmbracelet/bubbletea"
)

type classModel struct {
	data     []data.Classes
	list     list.Model
	selected string
}

// Instantiates classModel with a list of classes
func NewClassModel() classModel {
	d, err := data.NewClasses()
	if err != nil || len(d) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	// Inject classes into list
	classes := data.ListClassNames(d)
	items := stringToItem(classes)
	l := newListModel()
	l.SetItems(items)

	l.Title = "Choose a class:"

	return classModel{data: d, list: l}
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
		if c.list.FilterState() == list.Filtering {
			break // don't match keys if filtering
		}

		switch keypress := msg.String(); keypress {
		case "esc": // Reset selection
			c.list.ResetFilter()
			c.list.Select(0)
			return c, nil
		case "s":
			if c.list.ShowStatusBar() {
				c.list.SetShowStatusBar(false)
				return c, nil
			}
			c.list.SetShowStatusBar(true)
			return c, nil
		}
	}

	var cmd tea.Cmd
	c.list, cmd = c.list.Update(msg)
	return c, cmd
}

func (c classModel) View() string {
	details := func(c classModel) string {
		item, ok := c.list.SelectedItem().(item)
		if ok {
			for _, class := range c.data {
				if string(item) == class.Class.Name {
					d := detailName.Render
					hd := d("\nHD: ") + class.Class.Hd
					proficiency := d("\nProficiencies: ") + class.Class.Proficiency
					numSkills := d("\nNum Skills: ") + class.Class.NumSkills
					armor := d("\nArmor Proficiencies: ") + class.Class.Armor
					weapons := d("\nWeapon Proficiencies: ") + class.Class.Weapons
					tools := d("\nTools Proficiencies: ") + class.Class.Tools
					wealth := d("\nStarting Wealth: ") + class.Class.Wealth

					return "\n" + hd + proficiency + numSkills + armor +
						weapons + tools + wealth
				}
			}
		}
		return ""
	}(c)

	details = detailsStyle.Render(details)
	return "\n" + lipgloss.JoinHorizontal(lipgloss.Top, c.list.View(), details)
}
