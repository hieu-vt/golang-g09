package model

import (
	"errors"
	"g09-to-do-list/common"
	"strings"
)

var (
	ErrTitleCannotBeEmpty = errors.New("title cannot be empty")
	ErrItemIsDeleted      = errors.New("item is deleted")
	ErrTitleCannotFound   = errors.New("item can not found")
)

const (
	TableName = "todo_items"
)

type TodoItem struct {
	common.SQLModel
	Title       string             `json:"title" gorm:"column:title;"`
	Description string             `json:"description" gorm:"column:description;"`
	Status      string             `json:"status" gorm:"column:status;"`
	UserId      int                `json:"-" gorm:"column:user_id;"`
	Owner       *common.SimpleUser `json:"owner" gorm:"foreignKey:UserId"`
	LikeCount   int                `json:"likeCount" gorm:"column:like_count"`
}

func (TodoItem) TableName() string {
	return TableName
}

func (item *TodoItem) Mask() {
	item.SQLModel.Mask(common.DbTypeItem)

	if v := item.Owner; v != nil {
		v.Mask()
	}
}

type TodoItemCreation struct {
	Id          int    `json:"id" gorm:"column:id;"`
	Title       string `json:"title" gorm:"column:title;"`
	Description string `json:"description" gorm:"column:description;"`
	UserId      int    `json:"userId" gorm:"column:user_id;"`
}

func (i *TodoItemCreation) Validate() error {
	i.Title = strings.TrimSpace(i.Title)

	if i.Title == "" {
		return ErrTitleCannotBeEmpty
	}

	return nil
}

func (TodoItemCreation) TableName() string { return TodoItem{}.TableName() }

type TodoItemUpdate struct {
	Title       *string `json:"title" gorm:"column:title;"`
	Description *string `json:"description" gorm:"column:description;"`
	Status      *string `json:"status" gorm:"column:status;"`
}

func (TodoItemUpdate) TableName() string { return TodoItem{}.TableName() }
