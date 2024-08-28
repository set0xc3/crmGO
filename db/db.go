package db

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type Client struct {
	Id       int    `json:"id"`
	Bookmark bool   `json:"bookmark"`
	Tags     string `json:"tags"`
	FullName string `json:"full_name"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
	Email    string `json:"email"`
	Date     string `json:"date"`
}

var DB *sql.DB

func Init() {
	var err error
	DB, err = sql.Open("sqlite3", "clients.db")
	if err != nil {
		log.Fatal(err)
	}

	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS clients (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			bookmark BOOL,
			tags TEXT,
			full_name TEXT,
			phone TEXT,
			address TEXT,
			email TEXT,
			date TEXT
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
}

func DeInit() {
	err := DB.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func CreateClient() (int64, error) {
	result, err := DB.Exec(
		`INSERT INTO clients (bookmark, tags, full_name, phone, address, email, date) 
		VALUES(?, ?, ?, ?, ?, ?, ?);`,
		false,
		"None",
		"Иван Иванов Иванович",
		"+7 (000) 000 0000",
		"None",
		"none@none.com",
		"2024/08/28",
	)
	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func DeleteClient(id int64) error {
	_, err := DB.Exec("DELETE FROM clients WHERE id = ?", id)
	return err
}

func ReadClientList() []Client {
	rows, err := DB.Query("SELECT * FROM clients")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	clients := make([]Client, 0)
	for rows.Next() {
		var client Client
		rows.Scan(&client.Id, &client.Bookmark, &client.Tags, &client.FullName, &client.Phone, &client.Address, &client.Email, &client.Date)
		if err != nil {
			log.Fatal(err)
		}
		clients = append(clients, client)
		fmt.Println(client)
	}

	return clients
}
