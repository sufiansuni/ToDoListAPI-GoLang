package database

import (
	"ToDoListAPI/env"
	"ToDoListAPI/models"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

var (
	host     = env.Get("DB_HOST")
	port     = env.Get("DB_PORT")
	user     = env.Get("DB_USER")
	password = env.Get("DB_PASSWORD")
	dbname   = env.Get("DB_NAME")
)

var DB *sql.DB

// Connect to PSQL DB
func Connect() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	DB = db
}

// Creates "tasks" table in database
func CreateTaskTable(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS " +
		"tasks" +
		" (" +
		"ID SERIAL PRIMARY KEY, " +
		"Name TEXT, " +
		"Done BOOL " +
		")")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Table Exists/Created: tasks")
	}
}

// Insert a new task entry into database
func InsertTask(db *sql.DB, newTask models.Task) (id int, err error) {
	sqlStatement := "INSERT INTO tasks (Name, Done) VALUES ($1, $2) RETURNING id"
	err = db.QueryRow(
		sqlStatement,
		newTask.Name,
		newTask.Done).Scan(&id)

	if err != nil {
		panic(err)
	}

	fmt.Println("New record ID is:", id)

	return
}

// Update a task entry in database
func UpdateTask(db *sql.DB, myTask models.Task) (rowsAffected int64, err error) {
	sqlStatement := "UPDATE tasks SET name = $2, done = $3 WHERE id = $1"
	result, err := db.Exec(
		sqlStatement,
		myTask.ID,
		myTask.Name,
		myTask.Done)

	if err != nil {
		panic(err)
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}
