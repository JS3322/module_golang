// package model

// import (
// 	"database/sql"
// 	"time"

// 	_ "github.com/mattn/go-sqlite3"
// )

// type sqliteHandler struct {
// 	db *sql.DB // instance
// }

// func newSqliteConnecter(filepath string) DBConnecter {
// 	db, err := sql.Open("sqlite3", filepath) // 파일db를 연다.
// 	if err != nil {
// 		panic(err)
// 	}
// 	statement, _ := db.Prepare( // 쿼리문 작성
// 		`CREATE TABLE IF NOT EXISTS todos (
// 			id 		  INTEGER PRIMARY KEY AUTOINCREMENT,
// 			sessionId STRING,
// 			name 	  TEXT,
// 			completed BOOLEAN,
// 			createdAt DATETIME
// 		);
// 		CREATE INDEX IF NOT EXISTS sessionIdIndexOnTodos ON todos (
// 			sessionId ASC
// 		);`)
// 	statement.Exec() // query setup

// 	return &sqliteHandler{db} // save instance
// }

// func (s *sqliteHandler) Close() { // db close function
// 	s.db.Close()
// }

// func (s *sqliteHandler) GetTodos(sessionId string) []*Todo {
// 	todos := []*Todo{}
// 	// todos에서 검색하면 모든 항목을 비교해서 가져와야하므로 비효율적 O(N) -> 위에 CREATE 할 때 만들어진 인덱스를 사용해야함
// 	// 지금은 디비가 작아서 그냥 todos에서 가져온다.
// 	rows, err := s.db.Query("SELECT id, name, completed, createdAt FROM todos WHERE sessionId=?", sessionId)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer rows.Close()

// 	for rows.Next() { // 다음 행으로 가면서 레코드를 읽음(모든 데이터)
// 		var todo Todo
// 		rows.Scan(&todo.ID, &todo.Name, &todo.Completed, &todo.CreatedAt) // 레코드를 Todo구조체에 넣어준다.
// 		todos = append(todos, &todo)
// 	}
// 	return todos
// }
// func (s *sqliteHandler) AddTodo(name string, sessionId string) *Todo {
// 	stmt, err := s.db.Prepare("INSERT INTO todos (sessionId, name, completed, createdAt) VALUES (?, ?, ?, datetime('now'))") // 쿼리문 작성
// 	if err != nil {
// 		panic(err)
// 	}
// 	rst, err := stmt.Exec(sessionId, name, false) // 쿼리문 실행(? 아규먼트 순서대로)
// 	if err != nil {
// 		panic(err)
// 	}
// 	id, _ := rst.LastInsertId() // 마지막으로 추가된 레코드의 id를 알려준다.
// 	var todo Todo
// 	todo.ID = int(id)
// 	todo.Name = name
// 	todo.Completed = false
// 	todo.CreatedAt = time.Now()

// 	return &todo
// }
// func (s *sqliteHandler) RemoveTodo(id int) bool {
// 	stmt, err := s.db.Prepare("DELETE FROM todos WHERE id=?") // 쿼리문 작성
// 	if err != nil {
// 		panic(err)
// 	}
// 	rst, err := stmt.Exec(id) // 쿼리문 실행(? 아규먼트 순서대로)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cnt, err := rst.RowsAffected() // 쿼리문으로 영향받은 레코드 갯수를 반환
// 	if err != nil {
// 		panic(err)
// 	}
// 	return cnt > 0 // 변경된 내용이 1개면 true, 0개면 false
// }
// func (s *sqliteHandler) CompleteTodo(id int, complete bool) bool {
// 	stmt, err := s.db.Prepare("UPDATE todos SET completed=? WHERE id=?") // 쿼리문 작성
// 	if err != nil {
// 		panic(err)
// 	}
// 	rst, err := stmt.Exec(complete, id) // 쿼리문 실행(? 아규먼트 순서대로)
// 	if err != nil {
// 		panic(err)
// 	}
// 	cnt, err := rst.RowsAffected() // 쿼리문으로 영향받은 레코드 갯수를 반환
// 	if err != nil {
// 		panic(err)
// 	}

// 	return cnt > 0 // 변경된 내용이 1개면 true, 0개면 false
// }