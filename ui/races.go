package ui

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"gitlab.com/wiggins.jonathan/dndcc/data"

	tea "github.com/charmbracelet/bubbletea"
)

type raceModel struct {
	data     []data.Races
	list     list.Model
	selected string
}

// Instantiates raceModel with a list of races
func NewRaceModel() raceModel {
	d, err := data.NewRaces()
	if err != nil || len(d) < 1 {
		fmt.Println("Can't read from data source:", err)
		os.Exit(1)
	}

	// Inject races into list
	races := data.ListRaceNames(d)
	items := stringToItem(races)
	l := newListModel()
	l.SetItems(items)

	l.Title = "Choose a race:"

	return raceModel{data: d, list: l}
}

func (r raceModel) Init() tea.Cmd {
	return nil
}

func (r raceModel) Update(msg tea.Msg) (raceModel, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		r.list.SetWidth(msg.Width)
		return r, nil
	case tea.KeyMsg:
		switch keypress := msg.String(); keypress {
		case "esc": // Reset selection
			r.list.ResetFilter()
			r.list.Select(0)
			return r, nil
		case "s":
			if r.list.ShowStatusBar() {
				r.list.SetShowStatusBar(false)
				return r, nil
			}
			r.list.SetShowStatusBar(true)
			return r, nil
		}
	}

	var cmd tea.Cmd
	r.list, cmd = r.list.Update(msg)
	return r, cmd
}

func (r raceModel) View() string {
	details := func(r raceModel) string {
		var size, speed, ability, proficiency, spellAbility, traits string
		item, ok := r.list.SelectedItem().(item)
		if ok {
			for _, races := range r.data {
				for _, race := range races.Race {
					if string(item) == race.Name {
						d := detailName.Render
						size = d("\nSize: ") + race.Size
						speed = d("\nSpeed: ") + race.Speed

						// ability is blank for custom lineage
						if race.Ability != "" {
							ability = d("\nAbility: ") + race.Ability
						}

						// proficiency & spellAbility has whitespace we need to
						// trim and can also be blank
						proficiency = strings.TrimSpace(race.Proficiency)
						if proficiency != "" {
							proficiency = d("\nProficiencies: ") + proficiency
						}
						spellAbility = strings.TrimSpace(race.SpellAbility)
						if spellAbility != "" {
							spellAbility = d("\nSpell Ability: ") + spellAbility
						}

						traits = "\n"
						for i, trait := range race.Trait {
							// second trait is missing newline character
							if i == 1 {
								traits += "\n"
							}

							traits += d(trait.Name+": ") + trait.Text
						}
					}
				}
			}
		}
		return "\n" + size + speed + ability + proficiency + spellAbility +
			traits
	}(r)

	details = detailsStyle.Render(details)
	return "\n" + lipgloss.JoinHorizontal(lipgloss.Top, r.list.View(), details)
}
