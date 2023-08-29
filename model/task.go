package model

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Priority string

const (
	HIGH   Priority = "HIGH"
	MEDIUM Priority = "MEDIUM"
	LOW    Priority = "LOW"
)

func (p Priority) IsValid() error {
	switch p {
	case HIGH, MEDIUM, LOW:
		return nil
	default:
		return fmt.Errorf("'%s' is not a valid priority", p)
	}
}

type Task struct {
	gorm.Model
	ListId   uint   `gorm:"foreignKey:List;index:listIdTaskName,unique"`
	Name     string `gorm:"index:listIdTaskName,unique"`
	DueDate  time.Time
	Done     bool
	Priority Priority
}

// for the bubbletea interface to be respected
func (t Task) FilterValue() string {
	return t.Name
}
func (t Task) Title() string {
	return t.Name
}

func (t Task) Description() string {
	var dueDateStr string = "None"
	if !t.DueDate.IsZero() {
		dueDateStr = t.DueDate.Format("2006-01-02 15:04")
	}
	var done string = "🗹"
	if !t.Done {
		done = "□"
	}
	return fmt.Sprintf("%s | %s | %s", dueDateStr, done, t.Priority)
}

func NewTask(name string, dueDate time.Time, priority Priority, listId uint) *Task {
	return &Task{
		Name:     name,
		DueDate:  dueDate,
		Done:     false,
		Priority: priority,
		ListId:   listId,
	}
}

func (t Task) String() string {
	return fmt.Sprintf("Id: %d Name: '%s'\n\tDue date: %s\n\tDone: %t\n\tPriority: %s", t.ID, t.Name, t.DueDate, t.Done, t.Priority)
}

type TaskRepository struct {
	DB *gorm.DB
}

func (tr *TaskRepository) CreateTask(listId uint, name string, dueDate time.Time, priority Priority) (*Task, error) {
	t := NewTask(
		name,
		dueDate,
		priority,
		listId,
	)

	tx := tr.DB.Create(t)
	if tx.Error != nil {
		return nil, fmt.Errorf("task creation error %w", tx.Error)
	}
	return t, nil
}

func (tr *TaskRepository) getTaskByName(taskName string) *Task {
	t := &Task{}
	result := tr.DB.First(t, "name = ?", taskName)
	if result.Error != nil {
		return nil
	}
	return t
}

func (tr *TaskRepository) GetTaskById(taskId uint) (*Task, error) {
	t := &Task{}
	result := tr.DB.First(t, "id = ?", taskId)
	if result.Error != nil {
		return nil, result.Error
	}
	return t, nil
}

// update task's Name
func (tr *TaskRepository) UpdateTaskName(taskId uint, newTaskName string) error {
	tx := tr.DB.Model(&Task{}).Where("id = ?", taskId).Update("name", newTaskName)
	if tx.Error != nil {
		return fmt.Errorf("task name update error %w", tx.Error)
	}
	return nil
}

// update task's DueDate
func (tr *TaskRepository) UpdateTaskDueDate(taskId uint, dueDate time.Time) error {
	tx := tr.DB.Model(&Task{}).Where("id = ?", taskId).Update("due_date", dueDate)
	if tx.Error != nil {
		return fmt.Errorf("task due date update error %w", tx.Error)
	}
	return nil
}

// update task's Done
func (tr *TaskRepository) UpdateTaskIsDone(taskId uint, isDone bool) error {
	tx := tr.DB.Model(&Task{}).Where("id = ?", taskId).Update("done", isDone)
	if tx.Error != nil {
		return fmt.Errorf("task is done update error %w", tx.Error)
	}
	return nil
}

// update task's Priority
func (tr *TaskRepository) UpdateTaskPriority(taskId uint, priority Priority) error {
	tx := tr.DB.Model(&Task{}).Where("id = ?", taskId).Update("priority", priority)
	if tx.Error != nil {
		return fmt.Errorf("task priority update error %w", tx.Error)
	}
	return nil
}

func (tr *TaskRepository) DeleteTaskById(taskId uint) error {
	tx := tr.DB.Unscoped().Delete(&Task{}, taskId)
	if tx.Error != nil {
		return fmt.Errorf("task deletion error %w", tx.Error)
	}
	return nil
}

func (tr *TaskRepository) DeleteTasks(listId uint) error {
	tx := tr.DB.Unscoped().Where("list_id = ?", listId).Delete(&Task{})
	if tx.Error != nil {
		return fmt.Errorf("tasks deletion error %w", tx.Error)
	}
	return nil
}

func (tr *TaskRepository) DeleteDoneTasks(listId uint) error {
	tx := tr.DB.Unscoped().Where("list_id = ?", listId).Where("done = 1").Delete(&Task{})
	if tx.Error != nil {
		return fmt.Errorf("tasks deletion error %w", tx.Error)
	}
	return nil
}

func (tr *TaskRepository) GetTasksByListId(listId uint) ([]Task, error) {
	var tasks []Task
	res := tr.DB.Where("list_id = ?", listId).Find(&tasks)
	if res.Error != nil {
		return nil, fmt.Errorf("error retrieving all tasks from project id: %d. %w", listId, res.Error)
	}
	return tasks, nil
}
func (tr *TaskRepository) GetDoneTasksByListId(listId uint) ([]Task, error) {
	var tasks []Task
	res := tr.DB.Where("list_id = ?", listId).Where("done = 1").Find(&tasks)
	if res.Error != nil {
		return nil, fmt.Errorf("error retrieving all tasks from project id: %d. %w", listId, res.Error)
	}
	return tasks, nil
}
func (tr *TaskRepository) TaskExists(taskId uint) (bool, error) {
	var exists bool
	res := tr.DB.Model(&Task{}).
		Select("count(*) > 0").
		Where("id = ?", taskId).
		Find(&exists)
	return exists, res.Error
}
