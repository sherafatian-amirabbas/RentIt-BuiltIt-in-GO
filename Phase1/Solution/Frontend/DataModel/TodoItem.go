
package DataModel

// TodoItem represents the todo item
type TodoItem struct {
	ID         int `json:"ID"`
	Title      string `json:"Title"`
	Desciption string `json:"Desciption"`
	Priority   int   `json:"Priority"` // 0 is highest
}