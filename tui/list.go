package tui

import (
	"strings"

	"github.com/Bilou4/godo/model"
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
}

func newListForm(listId int, currentTitle string) *ListForm {
	form := &ListForm{listId: listId}
	form.title = textinput.New()
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
		if key.Matches(msg, key.NewBinding(key.WithKeys("q"))) {
			return m, tea.Quit
		}
		switch msg.String() {
		case "enter":
			if m.title.Focused() {
				return mainModel, m.NewList
			}
		}
	}
	if m.title.Focused() {
		m.title, cmd = m.title.Update(msg)
		return m, cmd
	} else {
		return m, cmd
	}
}

func (m ListForm) helpMenu() string {
	return helpStyle.Render("enter: submit")
}

func (m ListForm) View() string {

	// Status bar
	statusBar := strings.Builder{}
	statusVal := statusText.Copy().
		Width(112 - lipgloss.Width(statusKey)).
		Render(m.msg)
	bar := lipgloss.JoinHorizontal(lipgloss.Top,
		statusStyle.Render(statusKey),
		statusVal,
	)
	statusBar.WriteString(statusBarStyle.Width(112).Render(bar))
	return lipgloss.JoinVertical(lipgloss.Left, m.title.View(), statusBar.String(), m.helpMenu())
}
