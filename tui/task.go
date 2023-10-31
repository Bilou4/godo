package tui

import (
	"strings"
	"time"

	"github.com/Bilou4/godo/configuration"
	"github.com/Bilou4/godo/model"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"gorm.io/gorm"
)

type TaskForm struct {
	listId   int
	taskId   int
	title    textinput.Model
	dueDate  textinput.Model
	priority textinput.Model
	msg      string
	styles   *TuiStyles
	keys     keyMapTask
	help     help.Model
}

func newForm(listId, taskId int, currentTitle, currentPriority string, currentDueDate time.Time, styles *TuiStyles) *TaskForm {
	form := &TaskForm{listId: listId, styles: styles}
	form.title = textinput.New()
	form.title.SetValue(currentTitle)
	form.dueDate = textinput.New()
	if currentDueDate.IsZero() {
		form.dueDate.Placeholder = configuration.DueDateFormat
	} else {
		form.dueDate.SetValue(currentDueDate.Format(configuration.DueDateFormat))
	}
	form.priority = textinput.New()
	form.priority.SetValue(currentPriority)
	form.taskId = taskId
	form.title.Focus()
	form.help = help.New()
	form.keys = getKeybindingsTask()
	return form
}

func (m TaskForm) NewTask() tea.Msg {
	dueDateStr := m.dueDate.Value()
	var dueDate time.Time
	if strings.ToLower(dueDateStr) == "none" || dueDateStr == "" {
		dueDate = time.Time{}
	} else {
		// we already checked the error in the dudeDateIsValid() function
		dueDate, _ = time.Parse(configuration.DueDateFormat, m.dueDate.Value())
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
	_, err := time.Parse(configuration.DueDateFormat, dueDateStr)
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
		switch {
		case key.Matches(msg, m.keys.Quit):
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, cmd
		case key.Matches(msg, m.keys.Submit):
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
		case key.Matches(msg, m.keys.Back):
			return mainModel, nil
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

func (m TaskForm) View() string {
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
	return lipgloss.JoinVertical(lipgloss.Left, m.title.View(), m.dueDate.View(), m.priority.View(), statusBar.String(), m.help.View(m.keys))
}
