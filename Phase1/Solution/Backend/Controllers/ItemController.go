package controller

import (
	model "ESI-Homework1/Backend/DataModel"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	// Type go get -u github.com/gorilla/mux to install
	"github.com/gorilla/mux"
)

// ItemController provides the controller functionalities
type ItemController struct {
}

// GetItems retreives all the items
func (IController ItemController) GetItems(w http.ResponseWriter, r *http.Request) {

	var repo = model.GetRepo()

	params := mux.Vars(r)
	key := params["key"]

	var items []*model.TodoItem
	if key == "sorted" {
		items = repo.GetTodoItemsOrderedByPriorities()
	} else if key == "inprogress" {
		items = repo.GetInProgressTodoItems()
	} else if key == "completed" {
		items = repo.GetCompletedTodoItems()
	} else {
		items = repo.GetTodoItems()
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// CompleteItem makes the status of the item as completed
func (IController ItemController) CompleteItem(w http.ResponseWriter, r *http.Request) {

	var repo = model.GetRepo()

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)

	cErr := repo.CompleteItemByID(id)
	if cErr == nil {

		fmt.Println("CompleteItem called: item completed")
	} else {

		fmt.Println("CompleteItem called - ", cErr.ToString())
	}
}

// CreateItem creates a new item
func (IController ItemController) CreateItem(w http.ResponseWriter, r *http.Request) {

	fmt.Println("CreateItem called")

	var repo = model.GetRepo()

	w.Header().Set("Content-Type", "application/json")

	var item *model.TodoItem
	_ = json.NewDecoder(r.Body).Decode(&item)

	cErr := repo.AddNewTodoItem(item)
	if cErr == nil {

		fmt.Println("CreateItem called: item created")
	} else {

		fmt.Println("CreateItem called - ", cErr.ToString())
	}
}

// DeleteItem removes the item
func (IController ItemController) DeleteItem(w http.ResponseWriter, r *http.Request) {

	var repo = model.GetRepo()

	params := mux.Vars(r)
	id, _ := strconv.ParseUint(params["id"], 10, 64)
	cErr := repo.DeleteTodoItemByID(id)
	if cErr == nil {

		fmt.Println("DeleteItem called: item deleted")
	} else {

		fmt.Println("DeleteItem called - ", cErr.ToString())
	}
}

// UpdateItem updates the item
func (IController ItemController) UpdateItem(w http.ResponseWriter, r *http.Request) {

	var repo = model.GetRepo()

	var item *model.TodoItem
	_ = json.NewDecoder(r.Body).Decode(&item)
	cErr := repo.UpdateTodoItem(item)
	if cErr == nil {

		fmt.Println("UpdateItem called: item updated")
	} else {

		fmt.Println("UpdateItem called - ", cErr.ToString())
	}
}
