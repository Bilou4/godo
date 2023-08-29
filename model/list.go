package model

import (
	"fmt"

	"gorm.io/gorm"
)

type List struct {
	gorm.Model
	Name string `gorm:"index:,unique"`
}

func NewList(name string) *List {
	return &List{Name: name}
}

func (l List) String() string {
	return fmt.Sprintf("%d: '%s'\n", l.ID, l.Name)
}

type ListRepository struct {
	DB *gorm.DB
}

func (lr *ListRepository) CreateList(name string) (*List, error) {
	l := NewList(name)
	tx := lr.DB.Create(&l)
	if tx.Error != nil {
		return nil, fmt.Errorf("list creation error %w", tx.Error)
	}
	return l, nil
}

// get all lists
func (lr *ListRepository) GetAllLists() ([]*List, error) {
	lists := []*List{}
	results := lr.DB.Find(&lists)
	if results.Error != nil {
		return nil, fmt.Errorf("database error: %w", results.Error)
	}
	return lists, nil
}

// update list name
func (lr *ListRepository) UpdateListName(listId uint, newListName string) error {
	tx := lr.DB.Model(&List{}).Where("id = ?", listId).Update("name", newListName)
	if tx.Error != nil {
		return fmt.Errorf("list name update error %w", tx.Error)
	}
	return nil
}

// remove task(s) from list
func (lr *ListRepository) DeleteList(listId uint) error {
	tx := lr.DB.Unscoped().Delete(&List{}, listId)
	if tx.Error != nil {
		return fmt.Errorf("list deletion error %w", tx.Error)
	}
	return nil
}

func (lr *ListRepository) getListByName(listName string) *List {
	l := &List{}
	result := lr.DB.First(l, "name = ?", listName)
	if result.Error != nil {
		return nil
	}
	return l
}

func (lr *ListRepository) getListById(listId uint) *List {
	l := &List{}
	result := lr.DB.First(l, "id = ?", listId)
	if result.Error != nil {
		return nil
	}
	return l
}

func (lr *ListRepository) ListExists(listId uint) (bool, error) {
	var exists bool
	res := lr.DB.Model(&List{}).
		Select("count(*) > 0").
		Where("id = ?", listId).
		Find(&exists)
	return exists, res.Error
}
