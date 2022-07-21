package cmd

import (
    "os"
    "fmt"

    "gitlab.com/wiggins.jonathan/dndcc/data"

    "github.com/spf13/cobra"
    tea "github.com/charmbracelet/bubbletea"
)

var rootCmd = &cobra.Command{
    Use     : "dndcc",
    Short   : "A DnD 5e Character Creator",
    Long    : "dndcc is a 5th Edition Dungeons & Dragons Character Creator",
    RunE    : func(cmd *cobra.Command, args []string) error {
        p := tea.NewProgram( initialModel() )
        if err := p.Start(); err != nil {
            return err
        }
        return nil
    },
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        os.Exit(1)
    }
}

type model struct {
    choices []string
    cursor  int
    selected map[int]struct{}
}

func initialModel() model {
    races, err := data.GetRaces()
    if err != nil {
        fmt.Println("Can't read from data source: ", err)
        os.Exit(1)
    }

    return model {
        choices: races,
        selected: make(map[int]struct{}),
    }
}

func (m model) Init() tea.Cmd {
    return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    // key presses
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.String() {
        case "ctrl+c", "q":
            return m, tea.Quit
        case "up", "k":
            if m.cursor > 0 {
                m.cursor--
            }
        case "down", "j":
            if m.cursor < len(m.choices) - 1 {
                m.cursor++
            }
        case "home", "g":           // go to top of list
            if m.cursor > 0 {
                m.cursor = 0
            }
        case "end", "G":            // go to bottom of list
            if m.cursor < len(m.choices) - 1 {
                m.cursor = len(m.choices) - 1
            }
        case "pgdown", "ctrl+f":    // down 10
        case "pgup", "ctrl+b":      // up 10
        case "enter", " ":          // select
            _, ok := m.selected[m.cursor]
            if ok {
                delete(m.selected, m.cursor)
            } else {
                m.selected[m.cursor] = struct{}{}
            }
        }
    }
    return m, nil
}

func (m model) View() string {
    s := "Choose a race:\n"
    for i, choice := range m.choices {
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        s += fmt.Sprintf("%s %s\n", cursor, choice)
    }

    s += "\n Press q to quit.\n"
    return s
}
