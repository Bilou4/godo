package tui

import (
	"strings"

	"github.com/Bilou4/godo/model"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gorm.io/gorm"
)

type ListForm struct {
	listId int
	title  textinput.Model
	msg    string
	styles *TuiStyles
	keys   keyMapList
	help   help.Model
}

func newListForm(listId int, currentTitle string, styles *TuiStyles) *ListForm {
	form := &ListForm{listId: listId, styles: styles}
	form.title = textinput.New()
	form.help = help.New()
	form.keys = getKeybindingsList()
	form.title.SetValue(currentTitle)
	form.title.Focus()
	return form
}

func (m ListForm) NewList() tea.Msg {
	list := model.List{Model: gorm.Model{ID: uint(m.listId)}, Name: m.title.Value()}
	return list
}

func (m ListForm) Init() tea.Cmd {
	return textinput.Blink
}

func (m ListForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, cmd
		case key.Matches(msg, m.keys.Submit):
			if m.title.Focused() {
				return mainModel, m.NewList
			}
			return m, cmd
		case key.Matches(msg, m.keys.Back):
			return mainModel, nil
		}
	}
	if m.title.Focused() {
		m.title, cmd = m.title.Update(msg)
		return m, cmd
	} else {
		return m, cmd
	}
}

func (m ListForm) View() string {

	// Status bar
	statusBar := strings.Builder{}
	statusVal := m.styles.StatusText.Copy().
		Width(112 - lipgloss.Width(statusKey)).
		Render(m.msg)
	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		m.styles.StatusStyle.Render(statusKey),
		statusVal,
	)
	statusBar.WriteString(m.styles.StatusBarStyle.Width(112).Render(bar))
	return lipgloss.JoinVertical(lipgloss.Left, m.title.View(), statusBar.String(), m.help.View(m.keys))
}
