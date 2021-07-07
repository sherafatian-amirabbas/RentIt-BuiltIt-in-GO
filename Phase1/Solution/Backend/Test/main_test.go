package test

import (
	model "ESI-Homework1/Backend/DataModel"
	"testing"
)

var repo = model.GetRepo()

// ----------------------------------------------------------------------- unit test for Create

// TestCreateItem tests the creation of a new item
func TestCreateItem(t *testing.T) {

	item := &model.TodoItem{

		ID:         1,
		Title:      "Title",
		Desciption: "Description",
		Priority:   1,
		Status:     0,
	}

	repo.AddNewTodoItem(item)
	var newItem = repo.FindItemByID(1)

	if item != newItem {

		t.Error("tst100: item was not found when it was added")
		return
	}

	if item.Status != 0 {

		t.Error("tst110: item was created but its status is not initialized as 'in progress'")
		return
	}
}

// TestCreateItemWithNil tests the creation of a new item with nil as argument
func TestCreateItemWithNil(t *testing.T) {

	var err = repo.AddNewTodoItem(nil)
	if err == nil || err.Code != "Vldt100" {

		t.Error("tst120: nil passed as an argument to AddNewTodoItem and it didn't validate it correctly")
	}
}

// TestCreateItemWithDuplicateID tests the creation of a new item with duplicate id
func TestCreateItemWithDuplicateID(t *testing.T) {

	item := &model.TodoItem{

		ID:         1,
		Title:      "New Title",
		Desciption: "New Description",
		Priority:   2,
		Status:     0,
	}

	var err = repo.AddNewTodoItem(item)
	if err == nil || err.Code != "Vldt200" {

		t.Error("tst130: an item with duplicate id passed to AddNewTodoItem and it didn't validate it correctly")
	}
}

// TestCreateItemWithDuplicateTitle tests the creation of a new item with duplicate title
func TestCreateItemWithDuplicateTitle(t *testing.T) {

	item := &model.TodoItem{

		ID:         2,
		Title:      "Title",
		Desciption: "New Description",
		Priority:   2,
		Status:     0,
	}

	var err = repo.AddNewTodoItem(item)
	if err == nil || err.Code != "Vldt300" {

		t.Error("tst140: an item with duplicate title passed to AddNewTodoItem and it didn't validate it correctly")
	}
}

// ----------------------------------------------------------------------- unit test for Delete

// TestDeleteItem tests the deletion of the item
func TestDeleteItem(t *testing.T) {

	repo.DeleteTodoItemByID(1)
	var item = repo.FindItemByID(1)

	if item != nil {

		t.Error("tst150: item was not deleted")
	}
}

// TestDeleteItemWhileItsNotExisted tests the deletion of the item which is not existed
func TestDeleteItemWhileItsNotExisted(t *testing.T) {

	var err = repo.DeleteTodoItemByID(1)
	if err == nil || err.Code != "Dlt100" {

		t.Error("tst160: not existed id passed to be deleted, but wasn't caught correctly")
	}
}

// ----------------------------------------------------------------------- unit test for Update

// TestCreateItem tests the update of a new item
func TestUpdateItem(t *testing.T) {

	originalItem := &model.TodoItem{

		ID:         1,
		Title:      "Title",
		Desciption: "Description",
		Priority:   1,
		Status:     0,
	}

	repo.AddNewTodoItem(originalItem)

	newItem := &model.TodoItem{

		ID:         1,
		Title:      "New Title",
		Desciption: "New Description",
		Priority:   2,
		Status:     0,
	}
	repo.UpdateTodoItem(newItem)

	var updatedItem = repo.FindItemByID(1)

	if newItem.Title != updatedItem.Title ||
		newItem.Desciption != updatedItem.Desciption ||
		newItem.Priority != updatedItem.Priority ||
		newItem.Status != updatedItem.Status {

		t.Error("tst170: item was not updated correctly")
		return
	}
}

// TestUpdateItemWhileItsNotSupplied tests the update with nil argument
func TestUpdateItemWhileItsNotSupplied(t *testing.T) {

	var err = repo.UpdateTodoItem(nil)
	if err == nil || err.Code != "Updt100" {

		t.Error("tst180: nil argument to update was not caught correctly")
	}
}

// TestUpdateItemWhileItsNotExisted tests the update of an item which is not existed
func TestUpdateItemWhileItsNotExisted(t *testing.T) {

	newItem := &model.TodoItem{

		ID:         2,
		Title:      "New Title",
		Desciption: "New Description",
		Priority:   2,
		Status:     0,
	}
	var err = repo.UpdateTodoItem(newItem)
	if err == nil || err.Code != "Updt200" {

		t.Error("tst190: item with none existed id to update was not caught correctly")
	}

}

// TestUpdateItemWhileTitleIsDuplicated tests the update of an item which has a duplicate title
func TestUpdateItemWhileTitleIsDuplicated(t *testing.T) {

	Item1 := &model.TodoItem{

		ID:         2,
		Title:      "AAAA",
		Desciption: "Description AAAA",
		Priority:   1,
		Status:     0,
	}
	repo.AddNewTodoItem(Item1)

	Item2 := &model.TodoItem{

		ID:         3,
		Title:      "BBBB",
		Desciption: "Description BBBB",
		Priority:   1,
		Status:     0,
	}
	repo.AddNewTodoItem(Item2)

	Item3 := &model.TodoItem{

		ID:         3,
		Title:      "AAAA",
		Desciption: "Description BBBB",
		Priority:   1,
		Status:     0,
	}

	var err = repo.UpdateTodoItem(Item3)
	if err == nil || err.Code != "Vldt300" {

		t.Error("tst200: an item was updated with duplicate title")
	}
}

// --------------------------------------------------------------- unit test for making a TODO item Completed

// TestItemCompletion tests if an item has become completed correctly
func TestItemCompletion(t *testing.T) {

	newItem := &model.TodoItem{

		ID:         4,
		Title:      "Title 4",
		Desciption: "Description 4",
		Priority:   4,
		Status:     0,
	}
	repo.AddNewTodoItem(newItem)

	var err = repo.CompleteItemByID(4)
	var updatedItem = repo.FindItemByID(4)

	if err != nil || updatedItem.Status != 1 {

		t.Error("tst210: was not brought to completion state correctly")
	}

}

// TestItemCompletionWhileItsNotExisted tests if an item is going to be completed but it's not existed
func TestItemCompletionWhileItsNotExisted(t *testing.T) {

	newItem := &model.TodoItem{

		ID:         5,
		Title:      "Title 5",
		Desciption: "Description 5",
		Priority:   5,
		Status:     0,
	}
	repo.AddNewTodoItem(newItem)

	var err = repo.CompleteItemByID(50000)
	if err == nil || err.Code != "Cmplt100" {

		t.Error("tst220: an item with invalid id wasn't caught while it was becomming complete")
	}

}
