package conf

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"os"
)

func GetDB() *gorm.DB {
	if "1" == os.Getenv("UnitTestEnv") {
		db, _ := getInMemoryDbClient()
		db = db.Debug()
		return db
	}

	c := Get()

	var db *gorm.DB
	switch c.Mail.Db.Type {
	case PgSql:
		if pgDb, err := gorm.Open(postgres.Open(c.Mail.Db.Connection), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		}); err != nil {
			panic(err)
		} else {
			db = pgDb
		}
	case Sqlite:
		if sqliteDb, err := gorm.Open(sqlite.Open(c.Mail.Db.Connection), &gorm.Config{
			DisableForeignKeyConstraintWhenMigrating: true,
		}); err != nil {
			panic(err)
		} else {
			db = sqliteDb
		}
	default:
		log.Fatal("unknown database type")
	}

	return db
}

// getInMemoryDbClient used for unit tests ONLY
func getInMemoryDbClient() (*gorm.DB, error) {
	if client, err := gorm.Open(sqlite.Open("file::memory:?cache=private"), &gorm.Config{}); err != nil {
		return nil, err
	} else {
		return client, nil
	}
}
