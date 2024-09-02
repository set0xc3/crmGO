package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
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

func CreateClient(client Client) (int64, error) {
	result, err := DB.Exec(
		`INSERT INTO clients (bookmark, tags, full_name, phone, address, email, date) 
		VALUES(?, ?, ?, ?, ?, ?, ?);`,
		client.Bookmark,
		client.Tags,
		client.FullName,
		client.Phone,
		client.Address,
		client.Email,
		client.Date,
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

func GetClientList() []Client {
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
	}

	return clients
}

func GetClientById(id int64) (bool, Client) {
	var client Client

	rows, err := DB.Query("SELECT * FROM clients WHERE id = ?", id)
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}

	if rows.Next() {
		rows.Scan(&client.Id, &client.Bookmark, &client.Tags, &client.FullName, &client.Phone, &client.Address, &client.Email, &client.Date)
		if err != nil {
			log.Fatal(err)
		}
		return true, client
	}

	return false, client
}

func UpdateClientById(id int64, client Client) bool {
	_, err := DB.Exec(
		"UPDATE clients SET bookmark = ?, tags = ?, full_name = ?, phone = ?, address = ?, email = ?, date = ? WHERE id = ?",
		client.Bookmark,
		client.Tags,
		client.FullName,
		client.Phone,
		client.Address,
		client.Email,
		client.Date,
		id,
	)
	if err != nil {
		log.Fatal(err)
		return false
	}

	return true
}
