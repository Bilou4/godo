package tui

import (
	"github.com/charmbracelet/bubbles/key"
)

// keyMapTask defines a set of keybindings. To work for help it must satisfy
// key.Map. It could also very easily be a map[string]key.Binding.
type keyMapTask struct {
	Help   key.Binding
	Back   key.Binding
	Submit key.Binding
	Quit   key.Binding
}

// ShortHelp returns keybindings to be shown in the mini help view. It's part
// of the key.Map interface.
func (k keyMapTask) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit, k.Back, k.Submit}
}

// FullHelp returns keybindings for the expanded help view. It's part of the
// key.Map interface.
func (k keyMapTask) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Help, k.Quit},   // first column
		{k.Back, k.Submit}, // second column
	}
}

func getKeybindingsTask() keyMapTask {
	return keyMapTask{

		Help: key.NewBinding(
			key.WithKeys("?"),
			key.WithHelp("?", "Toggle help"),
		),
		Quit: key.NewBinding(
			key.WithKeys("ctrl+q", "ctrl+c"),
			key.WithHelp("ctrl+q", "Quit"),
		),
		Back: key.NewBinding(
			key.WithKeys("esc"),
			key.WithHelp("esc", "Cancel"),
		),
		Submit: key.NewBinding(
			key.WithKeys("enter"),
			key.WithHelp("enter", "Next/Submit"),
		),
	}
}
