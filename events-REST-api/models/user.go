package models

import (
	"errors"

	"example.com/events-api/db"
	"example.com/events-api/utils"
)

type User struct {
	ID       int64  `json:"id"`
	Name    string `json:"name"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}


func (user User) Save() error {
	query := `INSERT INTO users(name, email, password)
	VALUES (?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	
	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Name, user.Email, hashedPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	user.ID = id
	return err
}

func (user *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, user.Email)
	
	var retrivedPassword string
	err := row.Scan(&user.ID, &retrivedPassword)
	if err != nil {
		return errors.New("invalid credentails")
	}

	passwordIsValid := utils.CheckPasswordHash(user.Password, retrivedPassword)

	if !passwordIsValid {
		return errors.New("invalid credentails")
	} 
	return nil

}