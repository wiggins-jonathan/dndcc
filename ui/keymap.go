package ui

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Up           key.Binding
	Down         key.Binding
	Select       key.Binding
	Back         key.Binding
	Filter       key.Binding
	Quit         key.Binding
	ShowFullHelp key.Binding

	Reset        key.Binding
	ToggleStatus key.Binding
}

func DefaultKeyMap() *keyMap {
	km := new(keyMap)
	km.Up = key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	)
	km.Down = key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	)
	km.Select = key.NewBinding(
		key.WithKeys("enter", " ", "tab"),
		key.WithHelp("tab", "select"),
	)
	km.Back = key.NewBinding(
		key.WithKeys("shift+tab"),
		key.WithHelp("shift+tab", "back"),
	)
	km.Filter = key.NewBinding(
		key.WithKeys("/"),
		key.WithHelp("/", "filter"),
	)
	km.Quit = key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	)
	km.ShowFullHelp = key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "more"),
	)

	km.Reset = key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "reset"),
	)
	km.ToggleStatus = key.NewBinding(
		key.WithKeys("s"),
		key.WithHelp("s", "status bar"),
	)

	return km
}

func (k *keyMap) ShortHelp() []key.Binding {
	return []key.Binding{
		k.Up, k.Down, k.Select, k.Back, k.Filter, k.Quit,
	}
}

func (k *keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		[]key.Binding{
			k.Down,
		},
	}
}
