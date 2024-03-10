package conf

import (
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func GetDB() *gorm.DB {
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
