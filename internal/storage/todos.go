package storage

type Todo struct {
	Text string `json:"text"`
}

var todos []Todo

func AddTodo(todo Todo) {
	todos = append(todos, todo)
}

func GetTodos() []Todo {
	return todos
}
