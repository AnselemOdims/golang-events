package db

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	var err error

	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	
	createTable()
}

func createTable() {
	createUsers := `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		email TEXT NOT NULL,
		PASSWORD TEXT NOT NULL
	);
	`
	_, err := DB.Exec(createUsers)

	if err != nil {
		fmt.Println("Error:", err)
		panic("Could not create users table")
	}

	createEvents := `
	CREATE TABLE IF NOT EXISTS events (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	venue TEXT NOT NULL,
	createdBy TEXT NOT NULL,
	FOREIGN KEY(createdBy) REFERENCES users(id)
	);
	`

	_, err = DB.Exec(createEvents)

	if (err != nil) {
		fmt.Println("Error:", err)
		panic("Could not create event table")
	}
}

func InsertIntoEvents() string {
	query := `
	INSERT INTO events(title, description, venue, createdBy)
	VALUES (?,?,?,?)
	`
	return query;
}

func InsertIntoUsers() string {
	query := `
	INSERT INTO users(email, password)
	VALUES (?,?)
	`
	return query;
}

func GetAllEvents() string {
	query := `SELECT * FROM events`
	return query
}

func GetEventByID() string {
	query := `SELECT * FROM events WHERE id = ?`
	return query
}

func UpdateEvent() string  {
	query := `
	UPDATE events
	SET title = ?, description = ?, venue = ?, createdBy = ?
	WHERE id = ?
	`
	return query
}

func DeleteEvent() string {
	query := `DELETE FROM events WHERE id = ?`
	return query
}

func GetUsers() string {
	query := `SELECT * FROM users`
	return query
}