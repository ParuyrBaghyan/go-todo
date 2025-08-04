package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	dsn := "root:Paruyr-2004-03-24@tcp(127.0.0.1:3306)/todo?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		panic("Could not connect to database" + err.Error())
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	err = DB.Ping()
	if err != nil {
		panic(err)
	}

	createTable()

	fmt.Println("Successfully connected to MySQL!")
}

func createTable() {
	createUsersTable := `
	CREATE TABLE IF NOT EXISTS users(
		id INTEGER PRIMARY KEY AUTO_INCREMENT,
		email VARCHAR(255) NOT NULL UNIQUE,
		password VARCHAR(255) NOT NULL
	); 
	`

	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic("Could not create users table." + err.Error())
	}

	createTodoTable := `
	CREATE TABLE IF NOT EXISTS todos (
	id INTEGER PRIMARY KEY AUTO_INCREMENT,
	title TEXT NOT NULL,
	description TEXT NOT NULL,
	dateTime DATETIME NOT NULL,
	user_id INTEGER,
	FOREIGN KEY(user_id) REFERENCES users(id)
	);
	`
	_, err = DB.Exec(createTodoTable)
	if err != nil {
		panic("Could not create todos table." + err.Error())
	}

}
