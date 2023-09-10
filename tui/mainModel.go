package tui

import (
	"fmt"
	"strings"
	"time"

	"github.com/Bilou4/godo/model"
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/paginator"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	mainModel tea.Model
)

const (
	divisor = 3
)

type tuiModel struct {
	focus          int
	loaded         bool
	lists          []list.Model
	paginator      paginator.Model
	quitting       bool
	tr             model.TaskRepository
	lr             model.ListRepository
	msg            string
	updatingTask   bool
	updatingTaskId int
	updatingList   bool
	width, height  int
	keys           keyMap
	help           help.Model
}

func NewModel(tr model.TaskRepository, lr model.ListRepository) tuiModel {
	m := tuiModel{focus: 0, loaded: false, tr: tr, lr: lr}
	m.help = help.New()
	m.keys = getKeybindings()

	m.paginator = paginator.New()
	m.paginator.Type = paginator.Dots
	m.paginator.PerPage = divisor
	m.paginator.ActiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "235", Dark: "252"}).Render("•")
	m.paginator.InactiveDot = lipgloss.NewStyle().Foreground(lipgloss.AdaptiveColor{Light: "250", Dark: "238"}).Render("•")

	return m
}

func (m *tuiModel) initLists(width, height int) {
	lists, err := m.lr.GetAllLists()
	if err != nil {
		m.msg = "error retrieving all lists" // err.Error()
	}

	var modelLists []list.Model
	for _, l := range lists {
		tasks, _ := m.tr.GetTasksByListId(l.ID)
		newList := list.New([]list.Item{}, list.NewDefaultDelegate(), m.calculateListWidth(), m.calculateListHeight())
		newList.SetShowHelp(false)
		for idx := range tasks {
			newList.InsertItem(idx, list.Item(&tasks[idx]))
		}

		newList.Title = l.Name
		modelLists = append(modelLists, newList)
	}
	m.lists = modelLists
	m.paginator.SetTotalPages(len(m.lists))
}

func (m *tuiModel) Next() {
	if m.focus+1 > len(m.lists)-1 {
		m.focus = 0
		m.paginator.Page = 0
	} else {
		if (m.focus+1)%m.paginator.PerPage == 0 {
			m.paginator.NextPage()
		}
		m.focus++
	}
}

func (m *tuiModel) Prev() {
	if m.focus-1 < 0 {
		if len(m.lists) == 0 {
			m.focus = 0
			m.paginator.Page = 0
		} else {
			m.focus = len(m.lists) - 1
			m.paginator.Page = m.paginator.TotalPages - 1
		}
	} else {
		if (m.focus)%m.paginator.PerPage == 0 {
			m.paginator.PrevPage()
		}
		m.focus--
	}
}

func (m tuiModel) calculateListWidth() int {
	return (m.width / divisor) - divisor
}

func (m tuiModel) calculateListHeight() int {
	return m.height * 2 / divisor
}

func (m tuiModel) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m tuiModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.help.Width = msg.Width
		if !m.loaded {
			columnStyle.Width(m.calculateListWidth())
			focusedStyle.Width(m.calculateListWidth())
			m.initLists(m.width, m.height)
			m.loaded = true
			// when no list exists
			if len(m.lists) == 0 {
				mainModel = m
				return newListForm(0, "").Update(nil)
			}
		} else {
			w := m.calculateListWidth()
			h := m.calculateListHeight()
			for i := range m.lists {
				m.lists[i].SetHeight(h)
				m.lists[i].SetWidth(w)
			}
			columnStyle.Width(w)
			focusedStyle.Width(w)
		}
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, key.NewBinding(key.WithKeys("q"))):
			m.quitting = true
			return m, tea.Quit
		case key.Matches(msg, m.keys.Help):
			m.help.ShowAll = !m.help.ShowAll
			return m, cmd
		case key.Matches(msg, m.keys.Right):
			m.msg = ""
			m.Next()
		case key.Matches(msg, m.keys.Left):
			m.Prev()
			m.msg = ""
		case key.Matches(msg, m.keys.RenameList):
			m.updatingList = true
			m.msg = ""
			// save the current model before switching to a new one
			mainModel = m
			// + 1 because in the db, listId starts at 1 and focus starts at 0
			return newListForm(m.focus+1, m.lists[m.focus].Title).Update(nil)
		case key.Matches(msg, m.keys.AddList):
			mainModel = m
			m.msg = ""
			return newListForm(0, "").Update(nil)
		case key.Matches(msg, m.keys.DeleteList):
			if len(m.lists[m.focus].Items()) != 0 {
				m.msg = "cannot delete a List containing Tasks"
				return m, cmd
			}
			lName := m.lists[m.focus].Title
			// remove the list from the db
			err := m.lr.DeleteListByName(lName)
			if err != nil {
				m.msg = fmt.Sprintf("Error deleting List: %s", err.Error())
			}
			// remove the list in the view
			m.lists = remove(m.lists, m.focus)
			// focus the previous list
			m.Prev()
			m.msg = fmt.Sprintf("List '%s' deleted", lName)
			m.paginator.SetTotalPages(len(m.lists))

			// if we deleted the last List, then creates a new one.
			if len(m.lists) == 0 {
				mainModel = m
				return newListForm(0, "").Update(nil)
			}
			return m, cmd
		case key.Matches(msg, m.keys.UpdateTask):
			m.msg = ""
			m.updatingTaskId = m.lists[m.focus].Index()
			task, ok := m.lists[m.focus].SelectedItem().(*model.Task)
			if !ok {
				m.msg = fmt.Sprintf("Error casting the Task (id=%d)", m.updatingTaskId)
				return m, cmd
			}
			m.updatingTask = true
			mainModel = m
			return newForm(m.focus+1, int(task.ID), task.Name, string(task.Priority), task.DueDate).Update(nil)
		case key.Matches(msg, m.keys.ToggleTask):
			selectedIdx := m.lists[m.focus].Index()
			task, ok := m.lists[m.focus].SelectedItem().(*model.Task)
			if !ok {
				m.msg = fmt.Sprintf("Error casting the Task (id=%d)", selectedIdx)
				return m, cmd
			}
			err := m.tr.UpdateTaskIsDone(task.ID, !task.Done)
			if err != nil {
				m.msg = fmt.Sprintf("Error toggling Task (id=%d): %s", selectedIdx, err.Error())
				return m, cmd
			}
			m.msg = fmt.Sprintf("Task '%s' toggled", task.Name)
			task.Done = !task.Done
			return m, m.lists[m.focus].SetItem(selectedIdx, task)
		case key.Matches(msg, m.keys.NewTask):
			m.msg = ""
			mainModel = m
			return newForm(m.focus+1, 0, "", "", time.Time{}).Update(nil)
		case key.Matches(msg, m.keys.DeleteTask):
			selectedIdx := m.lists[m.focus].Index()
			task, ok := m.lists[m.focus].SelectedItem().(*model.Task)
			if !ok {
				m.msg = fmt.Sprintf("Error casting the Task (id=%d)", m.updatingTaskId)
				return m, cmd
			}
			err := m.tr.DeleteTaskById(task.ID)
			if err != nil {
				m.msg = fmt.Sprintf("Error deleting Task: %s", err.Error())
				return m, cmd
			}
			m.lists[m.focus].RemoveItem(selectedIdx)
			if len(m.lists[m.focus].Items()) > 0 {
				// go back to the element previous to the one deleted
				if selectedIdx == 0 {
					m.lists[m.focus].Select(0)
				} else {
					m.lists[m.focus].Select(selectedIdx - 1)
				}
			}
			m.msg = fmt.Sprintf("Task '%s' removed", task.Name)
			return m, cmd
		}
	// new Task or update Task
	case model.Task:
		task := msg
		if m.updatingTask {
			// we are updating an existing task
			idx := m.updatingTaskId
			m.updatingTaskId = 0
			m.updatingTask = false
			err := m.tr.UpdateTaskName(task.ID, task.Name)
			if err != nil {
				m.msg = fmt.Sprintf("Error updating Task name: %s", err.Error())
				return m, cmd
			}
			err = m.tr.UpdateTaskDueDate(task.ID, task.DueDate)
			if err != nil {
				m.msg = fmt.Sprintf("Error updating Task due date: %s", err.Error())
				return m, cmd
			}
			err = m.tr.UpdateTaskPriority(task.ID, task.Priority)
			if err != nil {
				m.msg = fmt.Sprintf("Error updating Task priority: %s", err.Error())
				return m, cmd
			}
			m.msg = fmt.Sprintf("Task '%s' updated", task.Name)
			return m, m.lists[m.focus].SetItem(idx, &task)
		} else {
			// it is a new task
			t, err := m.tr.CreateTask(task.ListId, task.Name, task.DueDate, task.Priority)
			if err != nil {
				m.msg = fmt.Sprintf("Error while creating '%s' task %s", task.Name, err)
				return m, cmd
			}
			m.msg = fmt.Sprintf("Task '%s' created", task.Name)
			return m, m.lists[m.focus].InsertItem(len(m.lists[m.focus].Items()), t)
		}
	// new List or update List
	case model.List:
		nlist := msg
		if m.updatingList {
			m.updatingList = false
			err := m.lr.UpdateListName(nlist.ID, nlist.Name)
			if err != nil {
				m.msg = fmt.Sprintf("Error updating List: %s", err.Error())
				return m, cmd
			}
			m.lists[m.focus].Title = nlist.Name
			return m, cmd
		} else {
			newList := list.New([]list.Item{}, list.NewDefaultDelegate(), m.calculateListWidth(), m.calculateListHeight())
			newList.SetShowHelp(false)
			newList.Title = nlist.Name
			m.lists = append(m.lists, newList)
			_, err := m.lr.CreateList(newList.Title)
			if err != nil {
				m.msg = fmt.Sprintf("Error creating List: %s", err.Error())
			}
			m.paginator.SetTotalPages(len(m.lists))
			m.msg = fmt.Sprintf("List: '%s' added", nlist.Name)
			return m, cmd
		}
	}
	currList, cmd := m.lists[m.focus].Update(msg)
	m.lists[m.focus] = currList
	return m, cmd
}

func (m tuiModel) View() string {
	var cols []string
	if m.quitting {
		return "Nothing to do.\n"
	}
	if m.loaded {
		// TODO: status bar as a component? https://pkg.go.dev/github.com/charmbracelet/soft-serve/server/ui/components/statusbar
		// Status bar
		statusBar := strings.Builder{}
		statusVal := statusText.Copy().
			Width(m.width - lipgloss.Width(statusKey)).
			Render(m.msg)
		bar := lipgloss.JoinHorizontal(lipgloss.Top,
			statusStyle.Render(statusKey),
			statusVal,
		)
		statusBar.WriteString(statusBarStyle.Width(m.width).Render(bar))
		start, end := m.paginator.GetSliceBounds(len(m.lists))
		for idx, ml := range m.lists[start:end] {
			if m.focus == idx+start {
				cols = append(cols, focusedStyle.Render(ml.View()))
			} else {
				cols = append(cols, columnStyle.Render(ml.View()))
			}
		}
		return lipgloss.JoinVertical(lipgloss.Left, lipgloss.JoinHorizontal(lipgloss.Left, cols...), m.paginator.View(), statusBar.String(), m.help.View(m.keys))
	} else {
		return "Loading..."
	}
}

func remove(slice []list.Model, s int) []list.Model {
	return append(slice[:s], slice[s+1:]...)
}
