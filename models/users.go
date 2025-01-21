package models

import (
	"errors"
	"event-planning/db"
	"event-planning/utils"
)

type User struct {
	ID int64 `json:"id"`
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (user User) SaveUsers() error {
	query := db.InsertIntoUsers()
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	password, err := utils.GenerateHashedPassword(user.Password)

	if err != nil {
		return err
	}

	result, err := stmt.Exec(user.Email, password)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	user.ID = id
	return err
}

func GetAllUsers() ([]User, error) {
	query := db.GetUsers()
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Email, &user.Password)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err = rows.Err(); err != nil {  
        return nil, err
    }  

	return users, nil
}

func (user *User) Login()  error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.DB.QueryRow(query, user.Email)

	var hashedPassword string
	err := row.Scan(&user.ID, &hashedPassword)

	if err != nil {
		return errors.New("Invalid credentials")
	}

	result := utils.ComparePassword(hashedPassword, user.Password)

	if !result {
		return errors.New("Invalid credentials")
	}

	return nil
}