package model

import (
	application "ESI-Homework1/Backend/Application"
)

// Predicate defines a delegate
type Predicate func(*TodoItem) bool

// TodoRepository holds the TodoItems
type TodoRepository struct {
	Items []*TodoItem
}

var repo *TodoRepository = &TodoRepository{}

// GetRepo shares the current instance of the repository
func GetRepo() *TodoRepository {
	return repo
}

// GetTodoItems returns all the TodoItems
func (repo *TodoRepository) GetTodoItems() []*TodoItem {
	return repo.Items
}

// GetTodoItemsOrderedByPriorities returns all the TodoItems sorted by priority
func (repo *TodoRepository) GetTodoItemsOrderedByPriorities() []*TodoItem {

	items := repo.GetTodoItems()
	SortTodoItemListByPriority(items)
	return items
}

// GetInProgressTodoItems returns all the in progress TodoItems
func (repo *TodoRepository) GetInProgressTodoItems() []*TodoItem {

	return repo.FindItemsByPredicate(func(item *TodoItem) bool {

		return item.Status == 0
	})
}

// GetCompletedTodoItems returns all the completed TodoItems
func (repo *TodoRepository) GetCompletedTodoItems() []*TodoItem {

	return repo.FindItemsByPredicate(func(item *TodoItem) bool {

		return item.Status == 1
	})
}

// CompleteItemByID makes the status of an item completed
func (repo *TodoRepository) CompleteItemByID(id uint64) *application.CustomError {

	cErr := &application.CustomError{}

	var item = repo.FindItemByID(id)
	if item == nil {

		cErr.Code = "Cmplt100"
		cErr.Message = "item was not found to be completed"
		return cErr
	}

	item.Status = 1

	return nil
}

// AddNewTodoItem adds the current item and returns the error if exists
func (repo *TodoRepository) AddNewTodoItem(item *TodoItem) *application.CustomError {

	var cErr = repo.ValidateTodoItemForCreate(item)
	if cErr != nil {

		return cErr
	}

	// when items are created they must be in progress
	// this should be handled in UI, also here is overwritten
	item.Status = 0

	repo.Items = append(repo.Items, item)
	return nil
}

// ValidateTodoItemForCreate validates item to be inserted - if validation succeed, error will be nil
func (repo *TodoRepository) ValidateTodoItemForCreate(item *TodoItem) *application.CustomError {

	return repo.validateTodoItem(item, false)
}

// ValidateTodoItemForUpdate validates item to be updated - if validation succeed, error will be nil
func (repo *TodoRepository) ValidateTodoItemForUpdate(item *TodoItem) *application.CustomError {

	return repo.validateTodoItem(item, true)
}

func (repo *TodoRepository) validateTodoItem(item *TodoItem, isUpdateMode bool) *application.CustomError {

	cErr := &application.CustomError{}

	if item == nil {

		cErr.Code = "Vldt100"
		cErr.Message = "item is not defined"
		return cErr
	}

	// ID is checked just in the create mode
	if !isUpdateMode {

		var duplicateItem = repo.FindItemByID(item.ID)
		if duplicateItem != nil {

			cErr.Code = "Vldt200"
			cErr.Message = "the parameter ID must be unique"
			return cErr
		}
	}

	// Title must be unique
	var duplicateItem = repo.FindItemByPredicate(func(todoItem *TodoItem) bool {

		var cond = todoItem.Title == item.Title
		if isUpdateMode {

			cond = cond && todoItem.ID != item.ID
		}

		return cond
	})
	if duplicateItem != nil {

		cErr.Code = "Vldt300"
		cErr.Message = "the parameter Title must be unique"
		return cErr
	}

	return nil
}

// FindItemByID finds the item by ID
func (repo *TodoRepository) FindItemByID(id uint64) *TodoItem {

	var item = repo.FindItemByPredicate(func(item *TodoItem) bool {

		return item.ID == id
	})

	return item
}

// FindItemByTitle finds item by Title
func (repo *TodoRepository) FindItemByTitle(title string) *TodoItem {

	var item = repo.FindItemByPredicate(func(item *TodoItem) bool {

		return item.Title == title
	})

	return item
}

// FindItemByPredicate finds the item by condition
func (repo *TodoRepository) FindItemByPredicate(predicate Predicate) *TodoItem {

	var item *TodoItem = nil
	for _, value := range repo.Items {

		if predicate(value) {

			item = value
			break
		}
	}

	return item
}

// FindItemsByPredicate finds items by condition
func (repo *TodoRepository) FindItemsByPredicate(predicate Predicate) []*TodoItem {

	var items []*TodoItem
	for _, item := range repo.Items {

		if predicate(item) {
			items = append(items, item)
		}
	}

	return items
}

// DeleteTodoItemByID removes the item by ID and returns the error if exists
func (repo *TodoRepository) DeleteTodoItemByID(id uint64) *application.CustomError {

	cErr := &application.CustomError{}

	var item = repo.FindItemByID(id)
	if item == nil {

		cErr.Code = "Dlt100"
		cErr.Message = "item was not found to delete"
		return cErr
	}

	for index, value := range repo.Items {

		if value.ID == id {

			repo.Items = append(repo.Items[:index], repo.Items[index+1:]...)
			break
		}
	}

	return nil
}

// UpdateTodoItem adds the current item and returns the error if exists
func (repo *TodoRepository) UpdateTodoItem(item *TodoItem) *application.CustomError {

	cErr := &application.CustomError{}

	if item == nil {

		cErr.Code = "Updt100"
		cErr.Message = "item is not defined"
		return cErr
	}

	var originalItem = repo.FindItemByID(item.ID)
	if originalItem == nil {

		cErr.Code = "Updt200"
		cErr.Message = "item was not found to update"
		return cErr
	}

	cErr = repo.ValidateTodoItemForUpdate(item)
	if cErr != nil {

		return cErr
	}

	originalItem.Title = item.Title
	originalItem.Desciption = item.Desciption
	originalItem.Priority = item.Priority
	originalItem.Status = item.Status

	return nil
}
