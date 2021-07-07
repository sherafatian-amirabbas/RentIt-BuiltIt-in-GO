package model

import (
	"sort"
)

// TodoItem represents the todo item
type TodoItem struct {
	ID         uint64 `json:"ID"`
	Title      string `json:"Title"`
	Desciption string `json:"Desciption"`
	Priority   uint   `json:"Priority"` // 0 is highest
	Status     uint8  `json:"Status"`   // 0: in progress, 1: completed
}

type todoItemList []*TodoItem

func (e todoItemList) Len() int {
	return len(e)
}

func (e todoItemList) Less(i, j int) bool {
	return e[i].Priority < e[j].Priority
}

func (e todoItemList) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

// SortTodoItemListByPriority sorts items by priority
func SortTodoItemListByPriority(items []*TodoItem) {

	sort.Sort(todoItemList(items))
}
