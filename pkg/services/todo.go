package services

type TodoService struct {
	todos map[int]*Todo
	id    int
}

type Todo struct {
	Id    int
	Title string
	Done  bool
}

func NewTodoService() *TodoService {
	return &TodoService{
		todos: map[int]*Todo{
			1: {Id: 1, Title: "Buy milk", Done: false},
			2: {Id: 2, Title: "Buy eggs", Done: false},
			3: {Id: 3, Title: "Buy bread", Done: true},
		},
	}
}

func (ts *TodoService) List() []*Todo {
	todos := make([]*Todo, 0, len(ts.todos))
	for _, todo := range ts.todos {
		todos = append(todos, todo)
	}
	return todos
}

func (ts *TodoService) Add(todo Todo) int {
	ts.id++
	todo.Id = ts.id
	ts.todos[ts.id] = &todo
	return ts.id
}

func (ts *TodoService) Remove(id int) {
	delete(ts.todos, id)
}

func (ts *TodoService) Get(id int) *Todo {
	return ts.todos[id]
}
