package models

import (
	"errors"
	"go-todo/db"
	"go-todo/utils"

)

type User struct {
	Id int64
	Email    string
	Password string
}

func (u User) SaveUser() error {
	query := "INSERT INTO users (email,password) VALUES(?,?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	return err

}

func (u User) ValidateUserCredentials() error {
	query := "SELECT id ,password FROM users WHERE email =?"
	row := db.DB.QueryRow(query,u.Email)

	var retrivedPassword string
	err := row.Scan(&u.Id, &retrivedPassword)
	if err != nil {
		return errors.New("Credentials invalid")
	}

	passwordIsValide := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if !passwordIsValide {
		return errors.New("Credentials invallid")
	}

	return  nil
}
