package db

import (
	"AirAccountEmailAdapter/email/db/migrations"
)

func init() {

	// db migrate automatic
	migrations.AutoMigrate()

	//db := GetDB()
	//
	//if db != nil {
	//
	//}
	//
	//// Create a table
	//createTableSQL := `
	//CREATE TABLE IF NOT EXISTS mails (
	//	id INTEGER PRIMARY KEY AUTOINCREMENT,
	//	sender varchar(256),
	//    receiver varchar(256),
	//    subject varchar(512),
	//    unread bool,
	//    fingerprint varchar(128),
	//	log_at DateTime
	//);`
	//_, err = db.Exec(createTableSQL)
	//if err != nil {
	//	fmt.Println(err)
	//}
}
