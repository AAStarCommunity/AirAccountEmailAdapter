package infra

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

const DB_FILE = "mails.db"

// SaveInbox will save a matched message into sqlite with Unread status whatever if it read on server
func SaveInbox() {
	//db, err := sql.Open("sqlite3", DB_FILE)
	//defer db.Close()
	//
	//if err != nil {
	//	log.Fatal(err)
	//}
	//insertSQL := "INSERT INTO mails (msg, log_at) VALUES (?, ?)"
	//_, err = db.Exec(insertSQL, info+*msg, time.Now())
	//if err != nil {
	//	log.Fatal(err)
	//}
}

// FetchUnreadMessage will fetch unread message by timestamp desc
func FetchUnreadMessage() {

}

func init() {
	db, err := sql.Open("sqlite3", DB_FILE)

	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	// Create a table
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS mails (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender varchar(256),
	    receiver varchar(256),
	    subject varchar(512),
	    unread bool,
	    fingerprint varchar(128),
		log_at DateTime
	);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		fmt.Println(err)
	}
}
