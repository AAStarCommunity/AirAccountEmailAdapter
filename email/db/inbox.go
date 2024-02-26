package db

import (
	"database/sql"
	"fmt"
)

const DB_FILE = "mails.db"

// MailInbox represents the recv and verify emails which ready / already process to gateway
func MailInbox() []*Mail {
	panic("Not Implemented")
}

// Save represent save the mail verified
func Save(m *Mail) error {
	panic("Not Implemented")
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
