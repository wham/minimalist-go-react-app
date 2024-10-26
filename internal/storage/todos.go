package storage

type Todo struct {
	Text string `json:"text"`
}

type Todos struct {
	list []Todo
}

func (t *Todos) Add(todo Todo) {
	t.list = append(t.list, todo)
}

func (t *Todos) List() []Todo {
	return t.list
}
