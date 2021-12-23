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

// Delete a task entry in database
func DeleteTask(db *sql.DB, taskID int) (rowsAffected int64, err error) {
	sqlStatement := "DELETE FROM tasks WHERE id = $1"
	result, err := db.Exec(
		sqlStatement,
		taskID,
	)

	if err != nil {
		panic(err)
	}

	rowsAffected, err = result.RowsAffected()
	if err != nil {
		panic(err)
	}

	return
}

// Selects ONE task entry in database
func SelectTask(db *sql.DB, taskID int) (myTask models.Task, err error) {
	sqlStatement := "SELECT id, name, done FROM tasks WHERE id=$1"
	err = db.QueryRow(sqlStatement, taskID).Scan(
		&myTask.ID,
		&myTask.Name,
		&myTask.Done,
	)

	if err != nil && err != sql.ErrNoRows {
		panic(err)
	}

	return
}

// Selects ALL task entries in database
func SelectAllTasks(db *sql.DB) (myTasks []models.Task, err error) {
	rows, err := db.Query("SELECT id, name, done FROM tasks")
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	for rows.Next() {
		var newTask models.Task
		err = rows.Scan(&newTask.ID, &newTask.Name, &newTask.Done)
		if err != nil {
			panic(err)
		}
		myTasks = append(myTasks, newTask)
	}

	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return
}
