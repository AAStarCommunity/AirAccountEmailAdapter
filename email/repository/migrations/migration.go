package migrations

import (
	"AirAccountEmailAdapter/conf"
	"gorm.io/gorm"
)

type Migration interface {
	Up(db *gorm.DB) error
	Down(db *gorm.DB) error
}

var migrations []Migration

func init() {
	// ascending order
	migrations = []Migration{
		&Migration20240310{},
	}
}

func AutoMigrate() {
	db := conf.GetDB()

	// TODO：skip migrate if exists '__migration' in db
	Migrate(db)
}

// Migrate 数据库变更同步
func Migrate(db *gorm.DB) {

	for i := 0; i < len(migrations); i++ {
		_ = migrations[i].Up(db)
	}
}

// Rollback 数据库变更回滚
func Rollback(db *gorm.DB) {

	for i := len(migrations) - 1; i >= 0; i-- {
		_ = migrations[i].Down(db)
	}
}
