package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // Import go-sqlite3 library
)

func main() {
	os.Remove("db_sqlite.db") // I delete the file to avoid duplicated records.
	// SQLite is a file based database.

	log.Println("Creating db_sqlite.db...")
	file, err := os.Create("db_sqlite.db") // Create SQLite file
	if err != nil {
		log.Fatal(err.Error())
	}
	file.Close()
	log.Println("db_sqlite.db created")

	sqliteDatabase, _ := sql.Open("sqlite3", "./db_sqlite.db") // Open the created SQLite File
	defer sqliteDatabase.Close()                               // Defer Closing the database
	createTable(sqliteDatabase)                                // Create Database Tables

	// INSERT RECORDS
	insertStudent(sqliteDatabase, "0001", "JS1", "CPU")
	insertStudent(sqliteDatabase, "0002", "JS2", "GPU")
	insertStudent(sqliteDatabase, "0003", "JS3", "KEYBOARD")
	insertStudent(sqliteDatabase, "0004", "JS4", "KEYBOARD")
	insertStudent(sqliteDatabase, "0005", "JS5", "KEYBOARD")
	insertStudent(sqliteDatabase, "0006", "JS6", "GPU")
	insertStudent(sqliteDatabase, "0007", "JS7", "GPU")
	insertStudent(sqliteDatabase, "0008", "JS8", "PHD")
	insertStudent(sqliteDatabase, "0009", "JS9", "CPU")

	// DISPLAY INSERTED RECORDS
	displayStudents(sqliteDatabase)
}

func createTable(db *sql.DB) {
	createStudentTableSQL := `CREATE TABLE student (
		"idStudent" integer NOT NULL PRIMARY KEY AUTOINCREMENT,		
		"code" TEXT,
		"name" TEXT,
		"program" TEXT		
	  );` // SQL Statement for Create Table

	log.Println("Create student table...")
	statement, err := db.Prepare(createStudentTableSQL) // Prepare SQL Statement
	if err != nil {
		log.Fatal(err.Error())
	}
	statement.Exec() // Execute SQL Statements
	log.Println("student table created")
}

// We are passing db reference connection from main to our method with other parameters
func insertStudent(db *sql.DB, code string, name string, program string) {
	log.Println("Inserting student record ...")
	insertStudentSQL := `INSERT INTO student(code, name, program) VALUES (?, ?, ?)`
	statement, err := db.Prepare(insertStudentSQL) // Prepare statement.
	// This is good to avoid SQL injections
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(code, name, program)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func displayStudents(db *sql.DB) {
	row, err := db.Query("SELECT * FROM student ORDER BY name")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()
	for row.Next() { // Iterate and fetch the records from result cursor
		var id int
		var code string
		var name string
		var program string
		row.Scan(&id, &code, &name, &program)
		log.Println("Student: ", code, " ", name, " ", program)
	}
}
