package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

// keyMap defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMap struct {
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
	Help       key.Binding
	RenameList key.Binding
	AddList    key.Binding
	DeleteList key.Binding
	UpdateTask key.Binding
	ToggleTask key.Binding
	NewTask    key.Binding
	DeleteTask key.Binding
	Quit       key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right},         // first column
		{k.AddList, k.RenameList, k.DeleteList}, // second column
		{k.NewTask, k.UpdateTask, k.ToggleTask, k.DeleteTask},
		{k.Help, k.Quit},
	}
}

func getKeybindings() keyMap {
	return keyMap{
		Up: key.NewBinding(
			key.WithKeys("up", "k"),
			key.WithHelp("↑/k", "Move up"),
		),
		Down: key.NewBinding(
			key.WithKeys("down", "j"),
			key.WithHelp("↓/j", "Move down"),
		),
		Left: key.NewBinding(
			key.WithKeys("left", "h"),
			key.WithHelp("←/h", "Move left"),
		),
		Right: key.NewBinding(
			key.WithKeys("right", "l"),
			key.WithHelp("→/l", "Move right"),
		),
		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "Toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("q", "esc", "ctrl+c"),
			key.WithHelp("q", "Quit"),
		),
		RenameList: key.NewBinding(
			key.WithKeys("r"),
			key.WithHelp("r", "Rename a List"),
		),
		AddList: key.NewBinding(
			key.WithKeys("a"),
			key.WithHelp("a", "Create a List"),
		),
		DeleteList: key.NewBinding(
			key.WithKeys("D"),
			key.WithHelp("D", "Delete an empty List"),
		),
		UpdateTask: key.NewBinding(
			key.WithKeys("u"),
			key.WithHelp("u", "Update a Task"),
		),
		ToggleTask: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Toggle a Task"),
		),
		NewTask: key.NewBinding(
			key.WithKeys("n"),
			key.WithHelp("n", "Create a Task"),
		),
		DeleteTask: key.NewBinding(
			key.WithKeys("d"),
			key.WithHelp("d", "Delete a Task"),
		),
	}
}
