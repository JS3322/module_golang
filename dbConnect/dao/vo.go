// package dbConnect

// import "time"

// type Todo struct {
// 	Id        int       `json:"id"`
// 	Name      string    `json:"name"`
// 	Completed bool      `json:"completed"`
// 	CreatedAt time.Time `json:"created_at"`
// }

// type DBConnecter interface {
// 	GetTodos(sessionId string) []*Todo
// 	AddTodo(sessionId string, name string) *Todo
// 	RemoveTodo(id int) bool
// 	CompleteTodo(id int, complete bool) bool
// 	Close()
// }

// func NewDBConnect(filePath string) DBConnecter {
// 	return newSqliteConnecter(filepath)
// }
