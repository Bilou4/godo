package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/Bilou4/godo/model"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gorm.io/gorm"
)

var (
	helpStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("241"))
)

type TaskForm struct {
	listId   int
	taskId   int
	title    textinput.Model
	dueDate  textinput.Model
	priority textinput.Model
	msg      string
}

func newForm(listId, taskId int, currentTitle, currentPriority string, currentDueDate time.Time) *TaskForm {
	form := &TaskForm{listId: listId}
	form.title = textinput.New()
	form.title.SetValue(currentTitle)
	form.dueDate = textinput.New()
	if currentDueDate.IsZero() {
		form.dueDate.Placeholder = "2006-05-20 15:12"
	} else {
		form.dueDate.SetValue(currentDueDate.Format("2006-05-20 15:12"))
	}
	form.priority = textinput.New()
	form.priority.SetValue(currentPriority)
	form.taskId = taskId
	form.title.Focus()
	return form
}

func (m TaskForm) NewTask() tea.Msg {
	dueDateStr := m.dueDate.Value()
	var dueDate time.Time
	if strings.ToLower(dueDateStr) == "none" || dueDateStr == "" {
		dueDate = time.Time{}
	} else {
		// we already checked the error in the dudeDateIsValid() function
		dueDate, _ = time.Parse("2006-01-02 15:04", m.dueDate.Value())
	}

	task := model.Task{ListId: uint(m.listId), Model: gorm.Model{ID: uint(m.taskId)}, Name: m.title.Value(), DueDate: dueDate, Priority: model.Priority(strings.ToUpper(m.priority.Value()))}
	return task
}

func (m TaskForm) Init() tea.Cmd {
	return textinput.Blink
}

func (m TaskForm) dueDateIsValid() bool {
	dueDateStr := m.dueDate.Value()
	if strings.ToLower(dueDateStr) == "none" || dueDateStr == "" {
		return true
	}
	_, err := time.Parse("2006-01-02 15:04", dueDateStr)
	return err == nil
}

func (m TaskForm) priorityIsValid() bool {
	err := model.Priority.IsValid(model.Priority(strings.ToUpper(m.priority.Value())))
	return err == nil
}

func (m TaskForm) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if key.Matches(msg, key.NewBinding(key.WithKeys("q"))) {
			return m, tea.Quit
		}
		switch msg.String() {
		case "enter":
			if m.title.Focused() {
				m.title.Blur()
				m.dueDate.Focus()
				return m, textinput.Blink
			} else if m.dueDate.Focused() {
				m.msg = ""
				if m.dueDateIsValid() {
					m.dueDate.Blur()
					m.priority.Focus()
					return m, textinput.Blink
				} else {
					m.msg = "dueDate is invalid"
					return m, textinput.Blink
				}
			} else {
				if m.priorityIsValid() {
					// switch to previous model, add task
					return mainModel, m.NewTask
				} else {
					m.msg = "priority is not valid"
					return m, textinput.Blink
				}
			}
		}
	}
	if m.title.Focused() {
		m.title, cmd = m.title.Update(msg)
		return m, cmd
	} else if m.dueDate.Focused() {
		m.dueDate, cmd = m.dueDate.Update(msg)
		return m, cmd
	} else if m.priority.Focused() {
		m.priority, cmd = m.priority.Update(msg)
		return m, cmd
	} else {
		return m, cmd
	}
}

func (m TaskForm) helpMenu() string {
	var msg string
	if m.title.Focused() || m.dueDate.Focused() {
		msg = "next"
	} else {
		msg = "submit"
	}
	return helpStyle.Render(fmt.Sprintf("enter: %s", msg))
}

func (m TaskForm) View() string {
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
	return lipgloss.JoinVertical(lipgloss.Left, m.title.View(), m.dueDate.View(), m.priority.View(), statusBar.String(), m.helpMenu())
}